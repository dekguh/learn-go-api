package service

import (
	"github.com/dekguh/learn-go-api/internal/api/model"
	"github.com/dekguh/learn-go-api/internal/api/repository"
)

type TodoService interface {
	CreateTodo(todo *model.Todo) (*model.Todo, error)
	FindAllTodos() ([]model.Todo, error)
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

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{repo: repo}
}
