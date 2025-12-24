package models

import "time"

type User struct {
	ID uint `gorm:"primaryKey"`

	Username string  `gorm:"size:50;uniqueIndex;not null"`
	Email    string  `gorm:"size:100;uniqueIndex;not null"`
	Phone    *string `gorm:"size:10;uniqueIndex"`

	Password string `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}
