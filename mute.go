package soundtouch

import (
  "log"

  "github.com/brutella/hc/characteristic"
  "github.com/llun/soundtouch-golang"
)

func (s *SoundTouch) setupMute(mute *characteristic.Mute, speaker *soundtouch.Speaker) {
  mute.SetValue(false)
  mute.OnValueRemoteUpdate(func(newValue bool) {
    speaker.PressKey(soundtouch.POWER)
  })

  go func(mute *characteristic.Mute) {
    nowPlaying, err := speaker.NowPlaying()
    if err != nil {
      log.Println("Cannot read speaker now-playing because of, %v", err)
      return
    }

    log.Printf("%v", nowPlaying)

    if nowPlaying.Source == soundtouch.STANDBY {
      mute.UpdateValue(true)
    } else {
      mute.UpdateValue(false)
    }
  }(mute)
}
