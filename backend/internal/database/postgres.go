package database

import (
	"esapp/configs"
	"esapp/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() (*gorm.DB, error) {

	config := configs.LoadConfig()

	dsn := config.DBUrl

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitDb(db *gorm.DB) error {
	db.AutoMigrate(&models.User{}, &models.Group{}, &models.GroupMember{})
	return nil
}
