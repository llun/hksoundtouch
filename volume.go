package soundtouch

import (
  "log"

  "github.com/brutella/hc/characteristic"
  "github.com/llun/soundtouch-golang"
)

func (s *SoundTouch) createVolume(speaker *soundtouch.Speaker) *characteristic.Volume {
  volume := characteristic.NewVolume()
  volume.SetValue(0)
  volume.OnValueRemoteUpdate(func(level float64) {
    speaker.SetVolume(int(level))
  })

  go func(volume *characteristic.Volume) {
    currentVolume, err := speaker.Volume()
    if err != nil {
      log.Println("Cannot read speaker volume because of, %v", err)
      return
    }
    volume.UpdateValue(float64(currentVolume.ActualVolume))
  }(volume)

  return volume
}
