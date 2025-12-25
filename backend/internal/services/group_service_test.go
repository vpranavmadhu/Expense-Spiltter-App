package services

import (
	"esapp/internal/dto"
	"esapp/models"
	"testing"
)

type mockGroupRepo struct {
	createErr error
	addErr    error

	requesterIsMember bool
	userAlreadyMember bool
	isMemberErr       error

	groups    []models.Group
	groupsErr error

	members    []models.User
	membersErr error
}

func (m *mockGroupRepo) CreateGroup(group *models.Group) error {
	group.ID = 1
	return m.createErr
}

func (m *mockGroupRepo) AddMember(member *models.GroupMember) error {
	return m.addErr
}

func (m *mockGroupRepo) IsMember(groupID, userID uint) (bool, error) {
	if userID == 1 {
		return m.requesterIsMember, m.isMemberErr
	}
	return m.userAlreadyMember, m.isMemberErr
}

func (m *mockGroupRepo) GetGroupsByUserID(userID uint) ([]models.Group, error) {
	return m.groups, m.groupsErr
}

func (m *mockGroupRepo) GetMembersByGroupID(groupID uint) ([]models.User, error) {
	return m.members, m.membersErr
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
		requesterIsMember: true,
		userAlreadyMember: false,
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

func TestListMembersSuccess(t *testing.T) {
	groupRepo := &mockGroupRepo{
		requesterIsMember: true,
		members: []models.User{
			{ID: 1, Username: "person", Email: "p@test.com"},
			{ID: 2, Username: "friend", Email: "f@test.com"},
		},
	}

	userRepo := &mockUserRepo{}

	service := NewGroupService(groupRepo, userRepo)

	members, err := service.ListMembers(1, 10)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(members) != 2 {
		t.Fatalf("expected 2 members, got %d", len(members))
	}
}
