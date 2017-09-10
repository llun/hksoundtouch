package hksoundtouch

import (
	"github.com/llun/soundtouch-golang"
)

type GetNowPlaying struct {
	store   *SoundTouch
	speaker *soundtouch.Speaker
}

func NewGetNowPlaying(store *SoundTouch, speaker *soundtouch.Speaker) *GetNowPlaying {
	return &GetNowPlaying{store, speaker}
}

func (a *GetNowPlaying) Run() {
	value, err := a.speaker.NowPlaying()
	if err != nil {
		return
	}

	a.store.power = value.Source != soundtouch.STANDBY
	a.store.Speaker.Mute.UpdateValue(!a.store.power)
	if value.Source != soundtouch.STANDBY {
		a.store.nowPlaying = value
	}
}

func (a *GetNowPlaying) Name() string {
	return "GetNowPlaying"
}

func (a *GetNowPlaying) RemoveDuplicateCommand() bool {
	return true
}
