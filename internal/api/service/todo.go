package service

import (
	"errors"

	"github.com/dekguh/learn-go-api/internal/api/model"
	"github.com/dekguh/learn-go-api/internal/api/repository"
)

type TodoService interface {
	CreateTodo(todo *model.Todo) (*model.Todo, error)
	FindAllTodos() ([]model.Todo, error)
	DeleteTodoById(id uint) error
}

type todoService struct {
	repo repository.TodoRepository
}

func (service *todoService) CreateTodo(todo *model.Todo) (*model.Todo, error) {
	if err := service.repo.Create(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (service *todoService) FindAllTodos() ([]model.Todo, error) {
	return service.repo.FindAll()
}

func (service *todoService) DeleteTodoById(id uint) error {
	if id == 0 {
		return errors.New("id is required")
	}

	if err := service.repo.DeleteById(id); err != nil {
		return err
	}

	return nil
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{repo: repo}
}
