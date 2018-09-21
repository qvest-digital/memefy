package ws

type ClientRegistration struct {
	currentMemes []string `json:"currentMemes"`
}

type NewMemes struct {
	newMemes []string `json:"newMemes"`
}