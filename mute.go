package hksoundtouch

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/log"
	"github.com/llun/soundtouch-golang"
)

func (s *SoundTouch) setupMute(mute *characteristic.Mute, speaker *soundtouch.Speaker) {
	mute.SetValue(false)
	mute.OnValueRemoteUpdate(func(newValue bool) {
		if s.power == newValue {
			speaker.PressKey(soundtouch.POWER)
		}
	})

	go func(mute *characteristic.Mute) {
		nowPlaying, err := speaker.NowPlaying()
		if err != nil {
			log.Debug.Printf("Cannot read speaker now-playing because of, %v", err)
			return
		}

		log.Debug.Printf("%v", nowPlaying)

		if nowPlaying.Source == soundtouch.STANDBY {
			mute.UpdateValue(true)
			s.power = false
			s.nowPlaying = nowPlaying
		} else {
			mute.UpdateValue(false)
			s.power = true
			s.nowPlaying = nowPlaying
		}
	}(mute)
}
