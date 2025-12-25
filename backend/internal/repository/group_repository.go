package repository

import (
	"esapp/models"

	"gorm.io/gorm"
)

type GroupRepository interface {
	CreateGroup(group *models.Group) error
	AddMember(member *models.GroupMember) error
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
