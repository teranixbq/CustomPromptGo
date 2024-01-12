package response

import "time"

type ResponsePrompt struct {
	ID           string `json:"id"`
	Category     string `json:"category"`
	Instructions string `json:"instructions"`
	CreatedAt    time.Time `json:"created_at"`
}