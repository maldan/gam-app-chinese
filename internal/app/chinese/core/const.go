package core

type Translate struct {
	Noun        []string `json:"noun"`
	Verb        []string `json:"verb"`
	Adjective   []string `json:"adjective"`
	Adverb      []string `json:"adverb"`
	Particle    []string `json:"particle"`
	Preposition []string `json:"preposition"`
}

type Word struct {
	Name          string    `json:"name"`
	Category      []string  `json:"category"`
	Transcription string    `json:"transcription"`
	Translate     Translate `json:"translate"`
	Power         int       `json:"power"`
}

var DataDir = ""
