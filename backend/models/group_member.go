package models

type GroupMember struct {
	ID      uint  `gorm:"primaryKey"`
	GroupID uint  `gorm:"not null;index"`
	Group   Group `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE"`
	UserID  uint  `gorm:"not null;index"`
	User    User  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
