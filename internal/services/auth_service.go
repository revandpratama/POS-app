package services

import (
	"fmt"
	"point-of-sales-app/helper"
	"point-of-sales-app/internal/dto"
	"point-of-sales-app/internal/entities"
	"point-of-sales-app/internal/repositories"
)

type AuthService interface {
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
	Register(req *dto.RegisterRequest) error
}

type authService struct {
	repo repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {

	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if !helper.CheckPasswordHash(req.Password, user.Password) {
		return nil, fmt.Errorf("invalid credentials: password")
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{Token: token}, nil
}

func (s *authService) Register(req *dto.RegisterRequest) error {

	if s.repo.EmailExists(req.Email) {
		return fmt.Errorf("email already exists")
	}

	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := &entities.User{
		Name: req.Name, 
		Email: req.Email, 
		Password: hashedPassword,
	}

	if err := s.repo.CreateUser(user); err != nil {
		return err
	}
	return nil
}
