package tts

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/Nonne46/number106/internal/utls"
	"github.com/gosimple/slug"
	"github.com/imroc/req/v3"
)

// const (
// 	URL       = "https://pubtts.ss14.su/api/v1/tts/"
// 	cachePath = "./client_cache"
// )

type TTS struct {
	*req.Client
	CachePath string
	Speakers  *Speakers
	Effects   []string
}

type TTSConfig struct {
	URL       string
	CachePath string
	Token     string
}

var DefaultConfig TTSConfig = TTSConfig{
	URL:       "http://127.0.0.1:2386/",
	CachePath: "./client_cache",
	Token:     "",
}

func NewTTS(config TTSConfig) TTS {
	tts := TTS{
		CachePath: config.CachePath,
		Speakers:  nil,
		Effects:   nil,
		Client: req.C().
			SetBaseURL(config.URL).
			SetCommonRetryCount(3).
			SetCommonRetryFixedInterval(1 * time.Second).
			SetCommonBearerAuthToken(config.Token),
	}

	return tts
}

func (t *TTS) GetTTS(speaker, pitch, text, effect string) (string, error) {
	if len(speaker) == 0 {
		return "", fmt.Errorf("Empty speaker")
	}

	if len(pitch) == 0 {
		return "", fmt.Errorf("Empty pitch")
	}

	if len(text) == 0 {
		return "", fmt.Errorf("Empty text")
	}

	crcHash := utls.CRC32Hash(speaker + pitch + text + effect)

	// Generate filename
	slug.MaxLength = 50
	filename := fmt.Sprintf("%s_%s.wav", slug.Make(text), crcHash)
	audioPath := fmt.Sprintf("%s/%s/%s", t.CachePath, speaker, filename)

	// Check file is in cache
	_, err := os.Stat(audioPath)
	if err == nil {
		return audioPath, nil
	}

	resp, err := t.R().
		AddQueryParam("speaker", speaker).
		AddQueryParam("pitch", pitch).
		AddQueryParam("text", text).
		AddQueryParam("effect", effect).
		Get("/")
	defer resp.Body.Close()

	if !resp.IsSuccessState() {
		err = fmt.Errorf("bad status code %q, body:%s", resp.StatusCode, resp.String())
		return "", err
	}

	// Save file
	out, err := os.Create(audioPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return audioPath, nil
}

// GetSpeakers ...
func (t *TTS) GetSpeakers() (*Speakers, error) {
	if t.Speakers != nil {
		return t.Speakers, nil
	}

	var speakers Speakers
	err := t.Get("/speakers").
		Do().
		Into(&speakers)
	if err != nil {
		return nil, err
	}

	t.Speakers = &speakers

	return t.Speakers, nil
}

// GetPitches ...
func (t *TTS) GetPitches() []string {
	return []string{
		"x-low",
		"low",
		"medium",
		"high",
		"x-high",
		"robot",
	}
}

// GetEffects ...
func (t *TTS) GetEffects() ([]string, error) {
	if t.Effects != nil {
		return t.Effects, nil
	}

	var effects Effects
	err := t.
		Get("/effects").
		Do().
		Into(&effects)
	if err != nil {
		return nil, err
	}

	t.Effects = effects.Effects

	return t.Effects, nil
}

func (t *TTS) CreateCacheDir() error {
	_ = os.Mkdir(t.CachePath, os.ModePerm)

	speakers, err := t.GetSpeakers()
	if err != nil {
		return err
	}

	for _, speaker := range speakers.Voices {
		speakerPath := fmt.Sprintf("%s/%s", t.CachePath, speaker.Speakers[0])
		_ = os.Mkdir(speakerPath, os.ModePerm)
	}

	return nil
}
