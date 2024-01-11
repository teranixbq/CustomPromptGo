package response

import "time"

type ResponsePrompt struct {
	ID           string `json:"id"`
	Instructions string `json:"instructions"`
	Category     string `json:"category"`
	CreatedAt    time.Time `json:"created_at"`
}