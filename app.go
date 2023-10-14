package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Nonne46/number106/internal/config"
	"github.com/Nonne46/number106/internal/mixer"
	"github.com/Nonne46/number106/internal/tts"
)

// App struct
type App struct {
	ctx context.Context
	tts *tts.TTS
}

// NewApp creates a new App application struct
func NewApp() *App {
	cfg := config.NewPaerser("./app.yaml")
	if err := cfg.Load(); err != nil {
		log.Panicf("Failed loading config: %s", err.Error())
	}

	url := cfg.Instance().Num.URL
	token := cfg.Instance().Num.Token
	cache := cfg.Instance().Num.Cache

	tts := tts.NewTTS(tts.TTSConfig{URL: url, Token: token, CachePath: cache})

	// Fetch some shit
	tts.GetSpeakers()
	tts.GetEffects()

	// Create cache
	tts.CreateCacheDir()

	return &App{tts: &tts}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// GetSpeakers ...
func (a *App) GetSpeakers() (*tts.Speakers, error) {
	return a.tts.GetSpeakers()
}

// GetPitches ...
func (a *App) GetPitches() []string {
	return a.tts.GetPitches()
}

// GetEffects ...
func (a *App) GetEffects() ([]string, error) {
	return a.tts.GetEffects()
}

// GetDevices ...
func (a *App) GetDevices() ([]string, error) {
	devices, err := mixer.GetDevices(false)
	if err != nil {
		return devices, err
	}

	return devices, nil
}

// PlayAudio ...
func (a *App) PlayAudio(DeviceName string, FileName string) error {
	if len(DeviceName) == 0 {
		return fmt.Errorf("Device is not selected")
	}

	if mixer.IsPlaying() {
		return fmt.Errorf("Audio is playing")
	}

	err := mixer.PlayAudio(DeviceName, FileName)
	if err != nil {
		return err
	}

	return nil
}

// IsPlaying ...
func (a *App) IsPlaying() bool {
	return mixer.IsPlaying()
}

// GetTTS ...
func (a *App) GetTTS(speaker, pitch, text, effect string) (string, error) {
	audioPath, err := a.tts.GetTTS(speaker, pitch, text, effect)
	if err != nil {
		return "", err
	}

	return audioPath, nil
}
