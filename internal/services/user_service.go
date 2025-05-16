package services

import (
	"point-of-sales-app/helper"
	"point-of-sales-app/internal/dto"
	"point-of-sales-app/internal/repositories"
)

type UserService interface {
	GetUsers() ([]dto.UserResponse, error)
	GetUser(id int) (*dto.UserResponse, error)
	UpdateUser(id int, user *dto.UserRequest) error
	DeleteUser(id int) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (u *userService) GetUsers() ([]dto.UserResponse, error) {
	users, err := u.repo.GetUsers()
	if err != nil {
		return nil, err
	}

	var usersResponse = make([]dto.UserResponse, len(users))

	for i, user := range users {
		usersResponse[i] = dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
	}

	return usersResponse, nil
}

func (u *userService) GetUser(id int) (*dto.UserResponse, error) {
	user, err := u.repo.GetUser(id)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (u *userService) UpdateUser(id int, user *dto.UserRequest) error {

	updateData := map[string]interface{}{}
	if user.Name != "" {
		updateData["name"] = user.Name
	}
	if user.Email != "" {
		updateData["email"] = user.Email
	}
	if user.Password != "" {
		hashedPassword, err := helper.HashPassword(user.Password)
		if err != nil {
			return err
		}
		updateData["password"] = hashedPassword
	}

	if len(updateData) == 0 {
		return nil
	}

	return u.repo.UpdateUser(id, updateData)
}

func (u *userService) DeleteUser(id int) error {
	return u.repo.DeleteUser(id)
}
