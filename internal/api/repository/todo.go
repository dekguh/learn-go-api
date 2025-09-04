package repository

import (
	"github.com/dekguh/learn-go-api/internal/api/model"
	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(todo *model.Todo) error
	FindAll() ([]model.Todo, error)
}

type todoRepository struct {
	db *gorm.DB
}

func (repo *todoRepository) Create(todo *model.Todo) error {
	if err := repo.db.Create(todo).Error; err != nil {
		return err
	}

	return nil
}

func (repo *todoRepository) FindAll() ([]model.Todo, error) {
	var todos []model.Todo
	if err := repo.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db}
}
