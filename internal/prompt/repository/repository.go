package repository

import (
	"errors"
	"promptgo/internal/prompt/dto/request"
	"promptgo/internal/prompt/dto/response"
	"promptgo/internal/prompt/model"
	"promptgo/util/constanta"

	"gorm.io/gorm"
)

type promptRepository struct {
	db *gorm.DB
}

type RepositoryInterface interface {
	Insert(data request.RequestPrompt) error
	SelectAll() ([]model.Prompt, error)
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

	err := prompt.FindInstructions(data.Instructions)
	if err != nil {
		return err
	}

	tx := prompt.db.Create(&input)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (prompt *promptRepository) SelectAll() ([]model.Prompt, error) {
	dataPrompt := []model.Prompt{}

	tx := prompt.db.Find(&dataPrompt).Order("created_at DESC")
	if tx.Error != nil {
		return nil, tx.Error
	}


	return dataPrompt, nil
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

	err := prompt.FindInstructions(data.Instructions)
	if err != nil {
		return err
	}

	tx := prompt.db.Where("id = ?", id).Updates(&input)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_NOT_FOUND)
	}

	return nil
}

func (prompt *promptRepository) Delete(id string) error {
	dataPrompt := model.Prompt{}

	tx := prompt.db.Where("id = ?", id).Delete(&dataPrompt)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_NOT_FOUND)
	}

	return nil
}

func (prompt *promptRepository) FindInstructions(instructions string) error {
	dataPrompt := model.Prompt{}

	err := prompt.db.Where("instructions = ? ", instructions).Find(&dataPrompt)
	if err.Error != nil {
		return err.Error
	}

	if err.RowsAffected != 0 {
		return errors.New(constanta.ERROR_DATA_EXIST)
	}

	return nil
}

// for chat

func (prompt *promptRepository) SelectAllPrompt() ([]response.AllPrompt, error) {
	dataPrompt := []model.Prompt{}

	tx := prompt.db.Find(&dataPrompt)
	if tx.Error != nil {
		return nil, tx.Error
	}

	response := response.ListModelToAllPrompt(dataPrompt)
	return response, nil
}
