package models

import "time"

type Group struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	CreatedBy uint   `gorm:"not null"`
	Creator   User   `gorm:"foreignKey:CreatedBy;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
}
