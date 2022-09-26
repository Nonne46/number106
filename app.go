package main

import (
	"context"
	"fmt"

	"github.com/nullneso/number106/internal/mixer"
	"github.com/nullneso/number106/internal/tts"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// GetSpeakers ...
func (a *App) GetSpeakers() []string {
	return tts.GetSpeakers()
}

// GetPitches ...
func (a *App) GetPitches() []string {
	return tts.GetPitches()
}

// GetDevices ...
func (a *App) GetDevices() ([]string, error) {
	devices, err := mixer.GetDevices(false)
	if err != nil {
		return devices, err
	}

	return devices, nil

	// return []string{"a", "b"}
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
func (a *App) GetTTS(speaker, pitch, text string) (string, error) {
	audioPath, err := tts.GetTTS(speaker, pitch, text)
	if err != nil {
		return "", err
	}

	return audioPath, nil
}
