package models

import "time"

type Expense struct {
	ID       uint    `gorm:"primaryKey"`
	GroupID  uint    `gorm:"not null;index"`
	Group    Group   `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE"`
	PaidByID uint    `gorm:"not null"`
	PaidBy   User    `gorm:"foreignKey:PaidByID;constraint:OnDelete:CASCADE"`
	Amount   float64 `gorm:"not null"`
	Title    string  `gorm:"size:100;not null"`

	CreatedAt time.Time
}

type ExpenseSplit struct {
	ID        uint    `gorm:"primaryKey"`
	ExpenseID uint    `gorm:"not null;index"`
	Expense   Expense `gorm:"foreignKey:ExpenseID;constraint:OnDelete:CASCADE"`
	UserID    uint    `gorm:"not null"`
	User      User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Amount    float64 `gorm:"not null"`

	IsSettled bool `gorm:"default:false"`
}
