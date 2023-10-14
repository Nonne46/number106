package tts

type Speakers struct {
	Voices []Voice `json:"voices"`
}

type Voice struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Speakers    []string `json:"speakers"`
}

type Effects struct {
	Effects []string `json:"effects"`
}
