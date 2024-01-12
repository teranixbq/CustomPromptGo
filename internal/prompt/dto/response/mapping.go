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

// func RequestRecybotToCoreRecybot(data RecybotRequest) entity.RecybotCore {
// 	return entity.RecybotCore{
// 		Question: data.Question,
// 	}
// }
