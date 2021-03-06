package hksoundtouch

import (
	"fmt"
	"net"
	"time"

	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/log"
	"github.com/brutella/hc/service"
	ba "github.com/llun/hkbridge/accessories"
	"github.com/llun/soundtouch-golang"
)

const BASE_TYPE = "00074"

type SoundTouch struct {
	*accessory.Accessory

	Speaker *service.Speaker
	Volume  *characteristic.Volume
	AUX     *Button
	Preset1 *Button
	Preset2 *Button
	Preset3 *Button
	Preset4 *Button
	Preset5 *Button
	Preset6 *Button
	IP      *IP

	SoundTouch *soundtouch.Speaker

	power      bool
	nowPlaying soundtouch.NowPlaying

	worker    *ba.Worker
	pollingCh <-chan time.Time
}

func Lookup(iface *net.Interface, worker *ba.Worker) []*SoundTouch {
	var services []*SoundTouch
	speakerCh := soundtouch.Lookup(iface)
	for speaker := range speakerCh {
		info, err := speaker.Info()
		if err != nil {
			log.Debug.Fatal(err)
		}
		services = append(services, NewSoundTouch(speaker, info.DeviceID, info.Type, worker))
	}
	return services
}

func NewSoundTouch(speaker *soundtouch.Speaker, serial, model string, worker *ba.Worker) *SoundTouch {
	info := accessory.Info{
		Name:         "SoundTouch",
		Manufacturer: "Bose",
		SerialNumber: serial,
		Model:        model,
	}

	acc := SoundTouch{
		worker: worker,
	}
	acc.Accessory = accessory.New(info, accessory.TypeOther)
	acc.SoundTouch = speaker

	acc.Speaker = acc.createSpeakerService()
	acc.AddService(acc.Speaker.Service)
	go acc.PollingState()

	return &acc
}

func (s *SoundTouch) PollingState() {
	s.worker.AddAction(NewGetNowPlaying(s, s.SoundTouch))
	s.worker.AddAction(NewGetVolume(s, s.SoundTouch))
	for range s.pollingCh {
		s.worker.AddAction(NewGetNowPlaying(s, s.SoundTouch))
		s.worker.AddAction(NewGetVolume(s, s.SoundTouch))
	}
}

func (s *SoundTouch) createSpeakerService() *service.Speaker {
	speaker := service.NewSpeaker()
	s.Volume = s.createVolume(s.SoundTouch)
	speaker.AddCharacteristic(s.Volume.Characteristic)

	s.IP = s.createIP(s.SoundTouch)
	speaker.AddCharacteristic(s.IP.Characteristic)

	s.Preset1 = s.createButton(fmt.Sprintf("%v%v", 1, BASE_TYPE), "Preset1", soundtouch.PRESET_1, s.SoundTouch)
	speaker.AddCharacteristic(s.Preset1.Characteristic)
	s.Preset2 = s.createButton(fmt.Sprintf("%v%v", 2, BASE_TYPE), "Preset2", soundtouch.PRESET_2, s.SoundTouch)
	speaker.AddCharacteristic(s.Preset2.Characteristic)
	s.Preset3 = s.createButton(fmt.Sprintf("%v%v", 3, BASE_TYPE), "Preset3", soundtouch.PRESET_3, s.SoundTouch)
	speaker.AddCharacteristic(s.Preset3.Characteristic)
	s.Preset4 = s.createButton(fmt.Sprintf("%v%v", 4, BASE_TYPE), "Preset4", soundtouch.PRESET_4, s.SoundTouch)
	speaker.AddCharacteristic(s.Preset4.Characteristic)
	s.Preset5 = s.createButton(fmt.Sprintf("%v%v", 5, BASE_TYPE), "Preset5", soundtouch.PRESET_5, s.SoundTouch)
	speaker.AddCharacteristic(s.Preset5.Characteristic)
	s.Preset6 = s.createButton(fmt.Sprintf("%v%v", 6, BASE_TYPE), "Preset6", soundtouch.PRESET_6, s.SoundTouch)
	speaker.AddCharacteristic(s.Preset6.Characteristic)

	s.AUX = s.createButton(fmt.Sprintf("%v%v", 20, BASE_TYPE), "AUX", soundtouch.AUX, s.SoundTouch)
	speaker.AddCharacteristic(s.AUX.Characteristic)

	s.setupMute(speaker.Mute, s.SoundTouch)
	return speaker
}

func AllAccessories(config ba.AccessoryConfig, iface *net.Interface, worker *ba.Worker) []*accessory.Accessory {
	speakers := Lookup(iface, worker)
	soundtouchAccessories := make([]*accessory.Accessory, len(speakers))
	for idx, speaker := range speakers {
		soundtouchAccessories[idx] = speaker.Accessory
	}
	log.Info.Println("Soundtouchs, %v", soundtouchAccessories)
	return soundtouchAccessories
}
