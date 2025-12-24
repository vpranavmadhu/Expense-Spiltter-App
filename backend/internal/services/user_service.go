package services

import (
	"esapp/internal/repository"
	"esapp/models"
)

type UserService interface {
	GetMe(userID uint) (*models.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetMe(userID uint) (*models.User, error) {
	return s.userRepo.FindByID(userID)
}
