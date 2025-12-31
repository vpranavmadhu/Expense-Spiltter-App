package services

import (
	"errors"
	"esapp/internal/dto"
	"esapp/internal/repository"
	"esapp/models"
)

type GroupService interface {
	CreateGroup(userID uint, req dto.CreateGroupRequest) error
	AddMemberByEmail(requesterID uint, groupID uint, email string) error
	ListGroups(userID uint) ([]models.Group, error)
	ListMembers(requesterID, groupID uint) ([]models.User, error)
	GetGroupByID(groupID, userID uint) (*models.Group, error)
	DeleteGroup(userID uint, groupID uint) error
}

type groupService struct {
	groupRepo repository.GroupRepository
	userRepo  repository.UserRepository
}

func NewGroupService(groupRepo repository.GroupRepository, userRepo repository.UserRepository) GroupService {
	return &groupService{groupRepo: groupRepo, userRepo: userRepo}
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

func (s *groupService) GetGroupByID(groupID, userID uint) (*models.Group, error) {
	isMember, err := s.groupRepo.IsMember(groupID, userID)
	if err != nil || !isMember {
		return nil, errors.New("not authorized")
	}

	return s.groupRepo.FindByID(groupID)
}

func (s *groupService) AddMemberByEmail(requesterID uint, groupID uint, email string) error {

	isMember, err := s.groupRepo.IsMember(groupID, requesterID)
	if err != nil {
		return err
	}
	if !isMember {
		return errors.New("not authorized")
	}

	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	exists, err := s.groupRepo.IsMember(groupID, user.ID)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("user already in the group")
	}

	memeber := &models.GroupMember{
		GroupID: groupID,
		UserID:  user.ID,
	}

	return s.groupRepo.AddMember(memeber)
}

func (s *groupService) ListGroups(userID uint) ([]models.Group, error) {
	return s.groupRepo.GetGroupsByUserID(userID)
}

func (s *groupService) ListMembers(requesterID, groupID uint) ([]models.User, error) {

	isMember, err := s.groupRepo.IsMember(groupID, requesterID)
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, errors.New("not authorized")
	}

	return s.groupRepo.GetMembersByGroupID(groupID)
}

func (s *groupService) DeleteGroup(userID uint, groupID uint) error {
	group, err := s.groupRepo.FindByID(groupID)
	if err != nil {
		return errors.New("group not found")
	}
	if group.CreatedBy != userID {
		return errors.New("unauthorized")
	}
	return s.groupRepo.DeleteGroup(groupID)
}
