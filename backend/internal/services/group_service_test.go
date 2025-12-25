package services

import (
	"esapp/internal/dto"
	"esapp/models"
	"testing"
)

type mockGroupRepo struct {
	createErr error
	addErr    error

	isMemberResult bool
	isMemberErr    error

	groups    []models.Group
	groupsErr error
}

func (m *mockGroupRepo) CreateGroup(group *models.Group) error {
	group.ID = 1
	return m.createErr
}

func (m *mockGroupRepo) AddMember(member *models.GroupMember) error {
	return m.addErr
}

func (m *mockGroupRepo) IsMember(groupID, userID uint) (bool, error) {
	return m.isMemberResult, m.isMemberErr
}

func (m *mockGroupRepo) GetGroupsByUserID(userID uint) ([]models.Group, error) {
	return m.groups, nil
}

func TestCreatGroupSuccess(t *testing.T) {
	groupRepo := &mockGroupRepo{}
	userRepo := &mockUserRepo{}
	service := NewGroupService(groupRepo, userRepo)

	req := dto.CreateGroupRequest{
		Name: "Goa Trip",
	}

	err := service.CreateGroup(10, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestCreateGroupEmptyName(t *testing.T) {
	groupRepo := &mockGroupRepo{}
	userRepo := &mockUserRepo{}
	service := NewGroupService(groupRepo, userRepo)

	req := dto.CreateGroupRequest{}

	err := service.CreateGroup(10, req)
	if err == nil {
		t.Fatalf("expected error for empty name")
	}
}

func TestAddMemberByEmail_Success(t *testing.T) {
	groupRepo := &mockGroupRepo{
		isMemberResult: true,
	}

	userRepo := &mockUserRepo{
		findUser: &models.User{
			ID:    2,
			Email: "friend@test.com",
		},
		findErr: nil,
	}

	service := NewGroupService(groupRepo, userRepo)

	err := service.AddMemberByEmail(1, 10, "friend@test.com")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestListGroupsSuccess(t *testing.T) {
	groupRepo := &mockGroupRepo{
		groups: []models.Group{
			{ID: 1, Name: "Goa Trip"},
			{ID: 2, Name: "Office Lunch"},
		},
	}

	userRepo := &mockUserRepo{}

	service := NewGroupService(groupRepo, userRepo)

	groups, err := service.ListGroups(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(groups) != 2 {
		t.Fatalf("expected 2 groups, got %d", len(groups))
	}
}
