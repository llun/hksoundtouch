package soundtouch

import (
  "github.com/brutella/hc/characteristic"
  "github.com/llun/soundtouch-golang"
)

const TypeIP = "100075"

type IP struct {
  *characteristic.String
}

func NewIP(value string) *IP {
  char := characteristic.NewString(TypeIP)
  char.Format = characteristic.FormatString
  char.Perms = []string{characteristic.PermRead}
  char.Description = "IP"
  char.SetValue(value)
  return &IP{char}
}

func (s *SoundTouch) createIP(speaker *soundtouch.Speaker) *IP {
  return NewIP(speaker.IP.String())
}
