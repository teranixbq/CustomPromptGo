package repository

import (
	"errors"
	"promptgo/internal/prompt/dto/request"
	"promptgo/internal/prompt/dto/response"
	"promptgo/internal/prompt/model"

	"gorm.io/gorm"
)

type promptRepository struct {
	db *gorm.DB
}

type RepositoryPrompt interface {
	Insert(data request.RequestPrompt) error
	SelectAll() ([]response.ResponsePrompt, error)
	SelectByID(id string) (response.ResponsePrompt, error)
	Update(id, instructions, category string) error
	Delete(id string) error
}

func NewPromptRepository(db *gorm.DB) RepositoryPrompt {
	return &promptRepository{
		db: db,
	}
}

func (prompt *promptRepository) Insert(data request.RequestPrompt) error {
	input := request.RequestPromptToModel(data)

	tx := prompt.db.Create(&input)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (prompt *promptRepository) SelectAll() ([]response.ResponsePrompt, error) {
	dataPrompt := []model.Prompt{}

	tx := prompt.db.Find(&dataPrompt)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}

	response := response.ListModelToResponsePrompt(dataPrompt)
	return response, nil
}

func (prompt *promptRepository) SelectByID(id string) (response.ResponsePrompt, error) {
	dataPrompt := model.Prompt{}

	tx := prompt.db.Where("id = ?", id).First(&dataPrompt)
	if tx.Error != nil {
		return response.ResponsePrompt{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return response.ResponsePrompt{}, errors.New("data not found")
	}

	response := response.ModelToResponsePrompt(dataPrompt)
	return response, nil
}

func (prompt *promptRepository) Update(id string, instructions string, category string) error {
	panic("unimplemented")
}

func (*promptRepository) Delete(id string) error {
	panic("unimplemented")
}

