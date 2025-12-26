package database

import (
	"esapp/configs"
	"esapp/models"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectPostgres() (*gorm.DB, error) {

	config := configs.LoadConfig()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	dsn := config.DBUrl

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitDb(db *gorm.DB) error {
	db.AutoMigrate(&models.User{}, &models.Group{}, &models.GroupMember{}, &models.Expense{}, &models.ExpenseSplit{})
	return nil
}
