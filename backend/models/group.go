package models

import "time"

type Group struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	CreatedBy uint   `gorm:"not null"`
	CreatedAt time.Time
}
