package mixer

import (
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

// GetDevices ...
func GetDevices(isCapture bool) ([]string, error) {
	err := sdl.Init(sdl.INIT_AUDIO)
	if err != nil {
		return nil, err
	}
	// defer sdl.Quit()

	numAudios := sdl.GetNumAudioDrivers()

	devices := make([]string, 0, numAudios)
	for i := 0; i < numAudios; i++ {
		deviceName := sdl.GetAudioDeviceName(i, isCapture)
		if len(deviceName) == 0 {
			continue
		}

		devices = append(devices, deviceName)
	}

	return devices, nil
}

// PlayAudio ...
func PlayAudio(DeviceName string, FileName string) error {
	if err := mix.OpenAudioDevice(mix.DEFAULT_FREQUENCY, mix.DEFAULT_FORMAT, mix.DEFAULT_CHANNELS, mix.DEFAULT_CHUNKSIZE, DeviceName, sdl.AUDIO_ALLOW_ANY_CHANGE); err != nil {
		return err
	}
	defer mix.CloseAudio()

	music, err := mix.LoadMUS(FileName)
	if err != nil {
		return err
	}
	defer music.Free()

	if err := music.Play(0); err != nil {
		return err
	}

	// Wait until it finishes playing
	for mix.PlayingMusic() {
		sdl.Delay(16)
	}

	return nil
}

// IsPlaying ...
func IsPlaying() bool {
	return mix.PlayingMusic()
}

// func main() {
// 	if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	defer sdl.Quit()

// 	if err := mix.Init(mix.INIT_MP3); err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	defer mix.Quit()

// 	numAudios := sdl.GetNumAudioDevices(false)

// 	devices := make([]string, 0, numAudios)
// 	for i := 0; i < numAudios; i++ {
// 		DeviceName := sdl.GetAudioDeviceName(i, false)
// 		devices = append(devices, DeviceName)
// 	}

// 	if err := mix.OpenAudioDevice(mix.DEFAULT_FREQUENCY, mix.DEFAULT_FORMAT, mix.DEFAULT_CHANNELS, mix.DEFAULT_CHUNKSIZE, devices[1], sdl.AUDIO_ALLOW_ANY_CHANGE); err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	defer mix.CloseAudio()

// 	music, err := mix.LoadMUS("test.wav")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer music.Free()

// 	if err := music.Play(0); err != nil {
// 		log.Fatal(err)
// 	}

// 	// Wait until it finishes playing
// 	for mix.PlayingMusic() {
// 		sdl.Delay(16)
// 	}
// }
