package database

import (
	"fmt"
	"promptgo/app/config"
	"promptgo/internal/prompt/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDBPostgres(cfg *config.Config) *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=Asia/Jakarta",
		cfg.DBHOST, cfg.DBUSER, cfg.DBPASS, cfg.DBNAME, cfg.DBPORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.Exec("SELECT 1").Error
	if err != nil {
		panic(err)
	}

	AutoMigrate(db)
	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.Prompt{})
}
