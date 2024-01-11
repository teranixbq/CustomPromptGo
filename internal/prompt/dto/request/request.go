package request

type RequestPrompt struct {
	Instructions string `json:"instructions"`
	Category     string `json:"category"`
}