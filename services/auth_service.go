package services

import (
	"PracticalProject/dto"
	"PracticalProject/models"
	"PracticalProject/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(input dto.RegisterInput) error
	Login(input dto.LoginInput) (*models.User, error)
}

type authService struct {
	repo repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	return &authService{repo}
}

func (s *authService) Register(input dto.RegisterInput) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Phone:    input.Phone,
		Address:  input.Address,
	}
	return s.repo.CreateUser(&user)
}

func (s *authService) Login(input dto.LoginInput) (*models.User, error) {
	user, err := s.repo.GetByIdEmail(input.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, err
	}
	return user, err
}
