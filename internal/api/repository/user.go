package repository

import (
	"github.com/dekguh/learn-go-api/internal/api/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*model.User, error)
	FindById(id uint) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (repo *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) FindById(ID uint) (*model.User, error) {
	var user model.User
	if err := repo.db.Where("id = ?", ID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
