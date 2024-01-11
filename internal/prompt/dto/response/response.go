package response

type ResponsePrompt struct {
	ID           string `json:"id"`
	Instructions string `json:"instructions"`
	Category     string `json:"category"`
	CreatedAt    string `json:"created_at"`
}