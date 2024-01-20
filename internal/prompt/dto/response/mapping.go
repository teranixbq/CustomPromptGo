package response

import "promptgo/internal/prompt/model"

func ModelToResponsePrompt(data model.Prompt) ResponsePrompt {
	return ResponsePrompt{
		ID:           data.ID,
		Category:     data.Category,
		Instructions: data.Instructions,
		CreatedAt:    data.CreatedAt,
	}
}

func ListModelToResponsePrompt(data []model.Prompt) []ResponsePrompt {
	list := []ResponsePrompt{}
	for _, v := range data {
		result := ModelToResponsePrompt(v)
		list = append(list, result)
	}
	return list
}


func ModelToAllPrompt(data model.Prompt) AllPrompt {
	return AllPrompt{
		Category:     data.Category,
		Instructions: data.Instructions,
	}
}

func ListModelToAllPrompt(data []model.Prompt) []AllPrompt {
	list := []AllPrompt{}
	for _, v := range data {
		result := ModelToAllPrompt(v)
		list = append(list, result)
	}
	return list
}
