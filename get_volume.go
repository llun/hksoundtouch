package hksoundtouch

import (
	"github.com/llun/soundtouch-golang"
)

type GetVolume struct {
	store   *SoundTouch
	speaker *soundtouch.Speaker
}

func NewGetVolume(store *SoundTouch, speaker *soundtouch.Speaker) *GetVolume {
	return &GetVolume{store, speaker}
}

func (a *GetVolume) Run() {
	value, err := a.speaker.Volume()
	if err != nil {
		return
	}
	a.store.Volume.UpdateValue(float64(value.ActualVolume))
}

func (a *GetVolume) Name() string {
	return "GetVolume"
}

func (a *GetVolume) RemoveDuplicateCommand() bool {
	return true
}
