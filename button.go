package soundtouch

import (
  "time"

  "github.com/brutella/hc/characteristic"
  "github.com/llun/soundtouch-golang"
)

const TypeButton = "74"

type Button struct {
  *characteristic.Bool
}

func NewButton(uuid, name string) *Button {
  char := characteristic.NewBool(uuid)
  char.Format = characteristic.FormatBool
  char.Perms = []string{
    characteristic.PermRead,
    characteristic.PermWrite,
    characteristic.PermEvents,
  }
  char.Description = name
  char.SetValue(false)
  return &Button{char}
}

func (s *SoundTouch) createButton(uuid, name string, value soundtouch.Key, speaker *soundtouch.Speaker) *Button {
  button := NewButton(uuid, name)
  button.OnValueRemoteUpdate(func(press bool) {
    speaker.PressKey(value)
    time.AfterFunc(1*time.Second, func() {
      button.SetValue(false)
    })
  })
  return button
}
