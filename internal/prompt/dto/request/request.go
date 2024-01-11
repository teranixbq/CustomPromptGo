package request

type RequestPrompt struct {
	Instructions string `json:"instructions"`
	Category     string `json:"category"`
}

type RequestQuestion struct {
	Question string `json:"question"`
}