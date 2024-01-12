package request

import "promptgo/internal/prompt/model"

func RequestPromptToModel(data RequestPrompt) model.Prompt {
	return model.Prompt{
		Instructions: data.Instructions,
		Category:     data.Category,
	}
}

// func ListRequestRecybotToCoreRecybot(data []RecybotManageRequest) []entity.RecybotCore {
// 	list := []entity.RecybotCore{}
// 	for _, v := range data {
// 		result := ManageRequestRecybotToCoreRecybot(v)
// 		list = append(list, result)
// 	}
// 	return list
// }

// func RequestRecybotToCoreRecybot(data RecybotRequest) entity.RecybotCore {
// 	return entity.RecybotCore{
// 		Question: data.Question,
// 	}
// }
