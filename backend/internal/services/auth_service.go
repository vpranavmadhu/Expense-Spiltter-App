package services

import (
	"errors"

	"esapp/internal/dto"
	"esapp/internal/repository"
	"esapp/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(req dto.RegisterRequest) error
	Login(req dto.LoginRequest) (string, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Register(req dto.RegisterRequest) error {
	// Check if email already exists
	if _, err := s.userRepo.FindByEmail(req.Email); err == nil {
		return errors.New("email already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return err
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: string(hash),
	}

	return s.userRepo.Create(&user)
}

func (s *authService) Login(req dto.LoginRequest) (string, error) {
	// Check email if present
	user, err := s.userRepo.FindByEmail(req.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("Invalid Credentials, Email not found")
		}
		return "", err
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errors.New("Invalid Credentials, Incorrect password")
	}

	//Generate Token
	token, err := GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
