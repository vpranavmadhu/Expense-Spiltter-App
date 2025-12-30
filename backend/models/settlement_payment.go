package models

import "time"

type SettlementPayment struct {
	ID         uint    `gorm:"primaryKey"`
	GroupID    uint    `gorm:"not null;index"`
	Group      Group   `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE"`
	FromUserID uint    `gorm:"not null"`
	FromUser   User    `gorm:"foreignKey:FromUserID"`
	ToUserID   uint    `gorm:"not null"`
	ToUser     User    `gorm:"foreignKey:ToUserID"`
	Amount     float64 `gorm:"not null"`
	CreatedAt  time.Time
}
