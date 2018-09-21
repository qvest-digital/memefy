package ws

type ClientRegistration struct {
	CurrentMemes []string `json:"currentMemes"`
}

type NewMemes struct {
	NewMemes []string `json:"newMemes"`
}

type Trigger struct {
	Meme string `json:"meme"`
}
