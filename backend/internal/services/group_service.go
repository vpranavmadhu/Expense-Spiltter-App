package services

import (
	"errors"
	"esapp/internal/dto"
	"esapp/internal/repository"
	"esapp/models"
)

type GroupService interface {
	CreateGroup(userID uint, req dto.CreateGroupRequest) error
}

type groupService struct {
	groupRepo repository.GroupRepository
}

func NewGroupService(groupRepo repository.GroupRepository) GroupService {
	return &groupService{groupRepo: groupRepo}
}

func (s *groupService) CreateGroup(userID uint, req dto.CreateGroupRequest) error {
	if req.Name == "" {
		return errors.New("group name required")
	}

	group := &models.Group{
		Name:      req.Name,
		CreatedBy: userID,
	}

	if err := s.groupRepo.CreateGroup(group); err != nil {
		return err
	}

	member := &models.GroupMember{
		GroupID: group.ID,
		UserID:  userID,
	}

	return s.groupRepo.AddMember(member)
}
