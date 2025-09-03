package service

import (
	"errors"

	"github.com/dekguh/learn-go-api/internal/api/model"
	"github.com/dekguh/learn-go-api/internal/api/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	GetUserByEmail(email string) (*model.User, error)
	GetUserById(ID uint) (*model.User, error)
	RegisterUser(email, name, password string) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func (service *userService) GetUserByEmail(email string) (*model.User, error) {
	if email == "" {
		return nil, errors.New("email is required")
	}

	result, err := service.repo.FindByEmail(email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, errors.New("failed to get user by email")
	}

	return &model.User{
		ID:        result.ID,
		Email:     result.Email,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
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

	return &model.User{
		ID:        result.ID,
		Email:     result.Email,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (service *userService) RegisterUser(email, name, password string) (*model.User, error) {
	if exists, _ := service.repo.FindByEmail(email); exists != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:    email,
		Name:     name,
		Password: string(hashedPassword),
	}

	if err := service.repo.Create(user); err != nil {
		return nil, err
	}

	return &model.User{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}
