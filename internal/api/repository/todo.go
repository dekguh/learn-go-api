package repository

import (
	"errors"
	"log"

	"github.com/dekguh/learn-go-api/internal/api/model"
	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(todo *model.Todo) error
	FindAll() ([]model.Todo, error)
	DeleteById(id uint) error
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

func (repo *todoRepository) DeleteById(id uint) error {
	result := repo.db.Delete(&model.Todo{}, id)
	if result.Error != nil {
		log.Println("failed to delete todo by id: ", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("todo not found")
	}

	return nil
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db}
}
