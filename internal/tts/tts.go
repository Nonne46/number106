package tts

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/gosimple/slug"
)

const (
	URL       = "http://134.249.87.182:8800/tts?"
	cachePath = "./client_cache"
)

func GetTTS(speaker, pitch, text string) (string, error) {
	if len(speaker) == 0 {
		return "", fmt.Errorf("Empty speaker")
	}

	if len(pitch) == 0 {
		return "", fmt.Errorf("Empty pitch")
	}

	if len(text) == 0 {
		return "", fmt.Errorf("Empty text")
	}

	// Create cache
	createCacheDir()

	// Generate filename
	slug.MaxLength = 100
	filename := fmt.Sprintf("%s_%s.wav", slug.Make(text), pitch)
	audioPath := fmt.Sprintf("%s/%s/%s", cachePath, speaker, filename)

	// Check file is in cache
	_, err := os.Stat(audioPath)
	if err == nil {
		return audioPath, nil
	}

	// Download file
	params := url.Values{}
	params.Add("speaker", speaker)
	params.Add("pitch", pitch)
	params.Add("text", text)

	resp, err := http.Get(URL + params.Encode())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
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

func createCacheDir() {
	_ = os.Mkdir(cachePath, os.ModePerm)
	for _, speaker := range GetSpeakers() {
		speakerPath := fmt.Sprintf("%s/%s", cachePath, speaker)
		_ = os.Mkdir(speakerPath, os.ModePerm)
	}
}

// GetSpeakers ...
func GetSpeakers() []string {
	return []string{
		"aidar",
		"baya",
		"kseniya",
		"xenia",
		"eugene",
		"mykyta",
		"ru_random",
		"ua_random",
	}
}

// GetPitches ...
func GetPitches() []string {
	return []string{
		"x-low",
		"low",
		"medium",
		"high",
		"x-high",
		"robot",
	}
}
