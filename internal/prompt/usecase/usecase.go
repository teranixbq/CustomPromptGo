package usecase

import (
	"promptgo/internal/prompt/dto/response"
	"promptgo/internal/prompt/repository"
)

type promptUsecase struct {
	promptRepository repository.RepositoryPrompt
}

// Delete implements UsecasePrompt.
func (*promptUsecase) Delete(id string) error {
	panic("unimplemented")
}

// Insert implements UsecasePrompt.
func (*promptUsecase) Insert(instructions string, category string) error {
	panic("unimplemented")
}

// SelectAll implements UsecasePrompt.
func (*promptUsecase) SelectAll() ([]response.ResponsePrompt, error) {
	panic("unimplemented")
}

// SelectByID implements UsecasePrompt.
func (*promptUsecase) SelectByID(id string) (response.ResponsePrompt, error) {
	panic("unimplemented")
}

// Update implements UsecasePrompt.
func (*promptUsecase) Update(id string, instructions string, category string) error {
	panic("unimplemented")
}

type UsecasePrompt interface {
	Insert(instructions, category string) error
	SelectAll() ([]response.ResponsePrompt, error)
	SelectByID(id string) (response.ResponsePrompt, error)
	Update(id, instructions, category string) error
	Delete(id string) error
}

func NewPromptRepository(promptRepository repository.RepositoryPrompt) UsecasePrompt {
	return &promptUsecase{
		promptRepository: promptRepository,
	}
}
