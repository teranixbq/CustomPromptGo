package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Prompt struct {
    ID           string
    Instructions string 
    Category     string 
    CreatedAt    time.Time      `gorm:"type:TIMESTAMP"`
    UpdatedAt    time.Time      `gorm:"type:TIMESTAMP"`
    DeletedAt    gorm.DeletedAt `gorm:"index"`
}


func (P *Prompt) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	P.ID = newUuid.String()

	return nil
}