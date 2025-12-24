package services

import (
	"esapp/internal/dto"
	"esapp/models"
	"testing"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func strPtr(s string) *string {
	return &s
}

type mockUserRepo struct {
	findUser  *models.User
	findErr   error
	createErr error

	findByIDUser *models.User
	findByIDErr  error
}

func (m *mockUserRepo) FindByEmail(email string) (*models.User, error) {
	return m.findUser, m.findErr
}

func (m *mockUserRepo) Create(user *models.User) error {
	return m.createErr
}

func (m *mockUserRepo) FindByID(id uint) (*models.User, error) {
	return m.findByIDUser, m.findByIDErr
}

func TestRegisterSuccess(t *testing.T) {
	repo := &mockUserRepo{
		findUser:  nil,
		findErr:   gorm.ErrRecordNotFound, // email not found
		createErr: nil,                    // create succeeds
	}

	service := NewAuthService(repo)

	req := dto.RegisterRequest{
		Username: "pranav",
		Email:    "test@test.com",
		Phone:    strPtr("9999999999"),
		Password: "password123",
	}

	err := service.Register(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestRegisterDuplicateEmail(t *testing.T) {
	repo := &mockUserRepo{
		findUser: &models.User{Email: "test@test.com"},
		findErr:  nil, // email exists
	}

	service := NewAuthService(repo)

	req := dto.RegisterRequest{
		Username: "pranav",
		Email:    "test@test.com",
		Phone:    strPtr("9999999999"),
		Password: "password123",
	}

	err := service.Register(req)
	if err == nil {
		t.Fatal("expected error for duplicate email")
	}
}

func TestLoginSuccess(t *testing.T) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte("password123"), 10)

	repo := &mockUserRepo{
		findUser: &models.User{
			ID:       1,
			Email:    "test@test.com",
			Password: string(hashed),
		},
		findErr: nil,
	}

	service := NewAuthService(repo)

	req := dto.LoginRequest{
		Email:    "test@test.com",
		Password: "password123",
	}

	token, err := service.Login(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if token == "" {
		t.Fatal("expected token, got empty string")
	}
}

func TestLoginUserNotFound(t *testing.T) {
	repo := &mockUserRepo{
		findUser: nil,
		findErr:  gorm.ErrRecordNotFound,
	}

	service := NewAuthService(repo)

	req := dto.LoginRequest{
		Email:    "missing@test.com",
		Password: "password123",
	}

	_, err := service.Login(req)
	if err == nil {
		t.Fatal("expected error for user not found")
	}
}

func TestLoginWrongPassword(t *testing.T) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte("password123"), 10)

	repo := &mockUserRepo{
		findUser: &models.User{
			Email:    "test@test.com",
			Password: string(hashed),
		},
		findErr: nil,
	}

	service := NewAuthService(repo)

	req := dto.LoginRequest{
		Email:    "test@test.com",
		Password: "wrongpassword",
	}

	_, err := service.Login(req)
	if err == nil {
		t.Fatal("expected invalid credentials error")
	}
}

func TestGenerateAndValidateToken(t *testing.T) {
	token, err := GenerateToken(10)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("expected valid token, got %v", err)
	}

	if claims.UserID != 10 {
		t.Fatalf("expected userID 10, got %d", claims.UserID)
	}
}

func TestValidateToken_InvalidToken(t *testing.T) {
	_, err := ValidateToken("invalid.token.value")
	if err == nil {
		t.Fatal("expected error for invalid token")
	}
}
