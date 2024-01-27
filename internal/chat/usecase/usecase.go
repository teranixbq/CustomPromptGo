package usecase

import (
	"context"
	"encoding/json"
	"promptgo/app/config"
	"promptgo/internal/prompt/dto/response"
	"promptgo/internal/prompt/repository"

	"github.com/sashabaranov/go-openai"
)

type ChatUsecase struct {
	promptRepository repository.RepositoryInterface
}

type UsecaseInterface interface {
	GetPrompt() ([]response.AllPrompt, error)
	Completion(question string) (string, error)
}

func NewChatUsecase(promptRepository repository.RepositoryInterface) UsecaseInterface {
	return &ChatUsecase{
		promptRepository: promptRepository,
	}
}

func (chat *ChatUsecase) GetPrompt() ([]response.AllPrompt, error) {
	dataRecybot, err := chat.promptRepository.SelectAll()
	if err != nil {
		return nil, err
	}

	result := response.ListModelToAllPrompt(dataRecybot)
	return result, nil
}

func (chat *ChatUsecase) Completion(question string) (string, error) {
	dataRecybot, err := chat.GetPrompt()
	if err != nil {
		return "", err
	}

	jsonData, err := json.Marshal(dataRecybot)
	if err != nil {
		return "", err
	}

	data := string(jsonData)
	cfg := config.InitConfig()
	ctx := context.Background()
	client := openai.NewClient(cfg.OPENAIKEY)
	model := openai.GPT3Dot5Turbo
	messages := []openai.ChatCompletionMessage{
		{
			Role:    "system",
			Content: data,
		},
		{
			Role:    "user",
			Content: question,
		},
	}

	response, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)
	if err != nil {
		return "", err
	}

	answer := response.Choices[0].Message.Content
	return answer, nil

}
