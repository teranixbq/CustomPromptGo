package usecase

import (
	"promptgo/internal/prompt/dto/request"
	"promptgo/internal/prompt/dto/response"
	"promptgo/internal/prompt/repository"
)

type promptUsecase struct {
	promptRepository repository.RepositoryInterface
}

type UsecaseInterface interface {
	Insert(data request.RequestPrompt) error
	SelectAll(question string) ([]response.ResponsePrompt, error)
	SelectByID(id string) (response.ResponsePrompt, error)
	Update(id string, data request.RequestPrompt) error
	Delete(id string) error
}

func NewPromptUsecase(promptRepository repository.RepositoryInterface) UsecaseInterface {
	return &promptUsecase{
		promptRepository: promptRepository,
	}
}

func (prompt *promptUsecase) Insert(data request.RequestPrompt) error {
	
	err := prompt.promptRepository.Insert(data)
	if err != nil {
		return err
	}

	return nil
}

func (prompt *promptUsecase) SelectAll(question string) ([]response.ResponsePrompt, error) {

	dataPrompt, err := prompt.promptRepository.SelectAll(question)
	if err != nil {
		return nil,err
	}

	return dataPrompt, nil
}

func (prompt *promptUsecase) SelectByID(id string) (response.ResponsePrompt, error) {

	dataPrompt, err := prompt.promptRepository.SelectByID(id)
	if err != nil {
		return response.ResponsePrompt{},err
	}

	return dataPrompt, nil
}

func (prompt *promptUsecase) Update(id string, data request.RequestPrompt) error {

	err := prompt.promptRepository.Update(id,data)
	if err != nil {
		return err
	}

	return nil
}

func (prompt *promptUsecase) Delete(id string) error {

	err := prompt.promptRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
