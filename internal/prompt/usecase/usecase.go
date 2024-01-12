package usecase

import (
	"promptgo/internal/prompt/dto/request"
	"promptgo/internal/prompt/dto/response"
	"promptgo/internal/prompt/repository"
	"promptgo/util/validation"
	"strings"
)

type promptUsecase struct {
	promptRepository repository.RepositoryInterface
}

type UsecaseInterface interface {
	Insert(data request.RequestPrompt) error
	SelectAll() ([]response.ResponsePrompt, error)
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

	cleanSpace := validation.CleanSpace(data.Instructions)
	toLower := strings.ToLower(cleanSpace)

	errEmpty := validation.CheckDataEmpty(data.Category, data.Instructions)
	if errEmpty != nil {
		return errEmpty
	}

	errLength := validation.MaxLength(data.Instructions, 100)
	if errLength != nil {
		return errLength
	}

	data.Instructions = toLower
	err := prompt.promptRepository.Insert(data)
	if err != nil {
		return err
	}

	return nil
}

func (prompt *promptUsecase) SelectAll() ([]response.ResponsePrompt, error) {

	dataPrompt, err := prompt.promptRepository.SelectAll()
	if err != nil {
		return nil, err
	}

	return dataPrompt, nil
}

func (prompt *promptUsecase) SelectByID(id string) (response.ResponsePrompt, error) {

	dataPrompt, err := prompt.promptRepository.SelectByID(id)
	if err != nil {
		return response.ResponsePrompt{}, err
	}

	return dataPrompt, nil
}

func (prompt *promptUsecase) Update(id string, data request.RequestPrompt) error {

	cleanSpace := validation.CleanSpace(data.Instructions)
	toLower := strings.ToLower(cleanSpace)

	errEmpty := validation.CheckDataEmpty(data.Category, data.Instructions)
	if errEmpty != nil {
		return errEmpty
	}

	errLength := validation.MaxLength(data.Instructions, 100)
	if errLength != nil {
		return errLength
	}

	data.Instructions = toLower
	err := prompt.promptRepository.Update(id, data)
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
