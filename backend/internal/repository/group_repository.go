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
	GetMembersByGroupID(groupID uint) ([]models.User, error)
	FindByID(groupID uint) (*models.Group, error)
	DeleteGroup(groupID uint) error
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
	err := r.db.
		Preload("Creator").
		Preload("Members").
		Joins("JOIN group_members gm ON gm.group_id = groups.id").
		Where("gm.user_id = ?", userID).
		Find(&groups).Error

	return groups, err
}

func (r *groupRepository) GetMembersByGroupID(groupID uint) ([]models.User, error) {
	var users []models.User

	err := r.db.
		Joins("JOIN group_members gm ON gm.user_id = users.id").
		Where("gm.group_id = ?", groupID).
		Find(&users).Error

	return users, err
}

func (r *groupRepository) FindByID(groupID uint) (*models.Group, error) {
	var group models.Group
	if err := r.db.First(&group, groupID).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *groupRepository) GetGroupByID(id uint) (*models.Group, error) {
	var group models.Group
	if err := r.db.First(&group, id).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *groupRepository) DeleteGroup(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE FROM group_members WHERE group_id = ?", id).Error; err != nil {
			return err
		}
		if err := tx.Where("group_id = ?", id).Delete(&models.Expense{}).Error; err != nil {
			return err
		}
		if err := tx.Where("group_id = ?", id).Delete(&models.SettlementPayment{}).Error; err != nil {
			return err
		}
		if err := tx.Delete(&models.Group{}, id).Error; err != nil {
			return err
		}
		return nil
	})
}
