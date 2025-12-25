package services

import (
	"esapp/internal/dto"
	"esapp/models"
	"testing"
)

type mockGroupRepo struct {
	createErr error
	addErr    error
}

func (m *mockGroupRepo) CreateGroup(group *models.Group) error {
	group.ID = 1
	return m.createErr
}

func (m *mockGroupRepo) AddMember(member *models.GroupMember) error {
	return m.addErr
}

func TestCreatGroupSuccess(t *testing.T) {
	repo := &mockGroupRepo{}
	service := NewGroupService(repo)

	req := dto.CreateGroupRequest{
		Name: "Goa Trip",
	}

	err := service.CreateGroup(10, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestCreateGroupEmptyName(t *testing.T) {
	repo := &mockGroupRepo{}
	service := NewGroupService(repo)

	req := dto.CreateGroupRequest{}

	err := service.CreateGroup(10, req)
	if err == nil {
		t.Fatalf("expected error for empty name")
	}
}
