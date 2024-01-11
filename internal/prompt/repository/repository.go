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

type RepositoryInterface interface {
	Insert(data request.RequestPrompt) error
	SelectAll(question string) ([]response.ResponsePrompt, error)
	SelectByID(id string) (response.ResponsePrompt, error)
	Update(id string, data request.RequestPrompt) error
	Delete(id string) error
}

func NewPromptRepository(db *gorm.DB) RepositoryInterface {
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

func (prompt *promptRepository) SelectAll(question string) ([]response.ResponsePrompt, error) {
	dataPrompt := []model.Prompt{}

	if question != "" {
		tx := prompt.db.Where("instructions LIKE ?", "%"+question+"%").Find(&dataPrompt)
		if tx.Error != nil {
			return nil, tx.Error
		}

		if tx.RowsAffected == 0 {
			return nil, errors.New("data not found")
		}

	} else {
		tx := prompt.db.Find(&dataPrompt)
		if tx.Error != nil {
			return nil, tx.Error
		}
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

func (prompt *promptRepository) Update(id string, data request.RequestPrompt) error {
	input := request.RequestPromptToModel(data)

	tx := prompt.db.Where("id = ?", id).Updates(&input)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (prompt *promptRepository) Delete(id string) error {
	dataPrompt := model.Prompt{}

	tx := prompt.db.Where("id = ?", id).Delete(&dataPrompt)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
