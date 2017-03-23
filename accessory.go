package soundtouch

import (
  "fmt"

  "github.com/brutella/hc/accessory"
  "github.com/brutella/hc/characteristic"
  "github.com/brutella/hc/service"
  "github.com/llun/soundtouch-golang"
)

const BASE_TYPE = "00074"

var speakerCh chan *soundtouch.Speaker = nil

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
}

func Lookup() {

}

func NewSoundTouch(serial string, model string) *SoundTouch {
  if speakerCh == nil {
    speakerCh = make(chan *soundtouch.Speaker, 1)
    soundtouch.Lookup(speakerCh)
  }

  info := accessory.Info{
    Name:         "SoundTouch",
    Manufacturer: "Bose",
    SerialNumber: serial,
    Model:        model,
  }

  acc := SoundTouch{}
  acc.Accessory = accessory.New(info, accessory.TypeOther)
  acc.SoundTouch = <-speakerCh

  acc.Speaker = acc.createSpeakerService()
  acc.AddService(acc.Speaker.Service)
  return &acc
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
