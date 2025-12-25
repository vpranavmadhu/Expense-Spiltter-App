package repository

import (
	"esapp/models"

	"gorm.io/gorm"
)

type GroupRepository interface {
	CreateGroup(group *models.Group) error
	AddMember(member *models.GroupMember) error
	IsMember(groupID, userID uint) (bool, error)
	GetGroupsByUserID(userID uint) ([]models.Group, error)
}

type groupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) GroupRepository {
	return &groupRepository{db: db}
}

func (r *groupRepository) CreateGroup(group *models.Group) error {
	return r.db.Create(group).Error
}

func (r *groupRepository) AddMember(member *models.GroupMember) error {
	return r.db.Create(member).Error
}

func (r *groupRepository) IsMember(groupID, userID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.GroupMember{}).Where("group_id = ? AND user_id = ?", groupID, userID).Count(&count).Error
	return count > 0, err
}

func (r *groupRepository) GetGroupsByUserID(userID uint) ([]models.Group, error) {
	var groups []models.Group

	err := r.db.Joins("JOIN group_members gm ON gm.group_id = groups.id").
		Where("gm.user_id = ?", userID).
		Find(&groups).Error

	return groups, err
}
