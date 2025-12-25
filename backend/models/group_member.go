package models

type GroupMember struct {
	ID      uint `gorm:"primaryKey"`
	GroupID uint `gorm:"not null;index"`
	UserID  uint `gorm:"not null;index"`
}
