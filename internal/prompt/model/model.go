package model

type Prompt struct {
	ID           string
	Instructions string
	Category     string `gorm:"unique"`
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string 
}


