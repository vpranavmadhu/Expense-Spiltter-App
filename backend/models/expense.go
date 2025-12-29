package models

import "time"

type Expense struct {
	ID       uint    `gorm:"primaryKey"`
	GroupID  uint    `gorm:"not null;index"`
	PaidByID uint    `gorm:"not null"`
	Amount   float64 `gorm:"not null"`
	Title    string  `gorm:"size:100;not null"`

	CreatedAt time.Time
}

type ExpenseSplit struct {
	ID        uint    `gorm:"primaryKey"`
	ExpenseID uint    `gorm:"not null;index"`
	UserID    uint    `gorm:"not null"`
	Amount    float64 `gorm:"not null"`

	IsSettled bool `gorm:"default:false"`
}
