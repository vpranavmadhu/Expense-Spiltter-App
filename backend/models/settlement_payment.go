package models

import "time"

type SettlementPayment struct {
	ID         uint    `gorm:"primaryKey"`
	GroupID    uint    `gorm:"not null;index"`
	FromUserID uint    `gorm:"not null"` // payer
	ToUserID   uint    `gorm:"not null"` // receiver
	Amount     float64 `gorm:"not null"`
	CreatedAt  time.Time
}
