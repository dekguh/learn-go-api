package service

import (
	"errors"
	"strings"

	"github.com/dekguh/learn-go-api/internal/api/model"
	"github.com/dekguh/learn-go-api/internal/api/repository"
)

type UserService interface {
	GetUserByEmail(email string) (*model.User, error)
	GetUserById(ID uint) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func (service *userService) GetUserByEmail(email string) (*model.User, error) {
	if email == "" {
		return nil, errors.New("email is required")
	}

	result, err := service.repo.FindByEmail(email)
	if strings.Contains(err.Error(), "record not found") {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, errors.New("failed to get user by email")
	}

	return result, nil
}

func (service *userService) GetUserById(ID uint) (*model.User, error) {
	if ID == 0 {
		return nil, errors.New("id is required")
	}

	result, err := service.repo.FindById(ID)
	if err != nil {
		return nil, errors.New("failed to get user by id")
	}

	if result == nil {
		return nil, errors.New("user not found")
	}

	return result, nil
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}
