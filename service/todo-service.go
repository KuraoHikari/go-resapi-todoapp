package service

import (
	"github.com/KuraoHikari/golang-todo-api/dto"
	"github.com/KuraoHikari/golang-todo-api/helper"
	"github.com/KuraoHikari/golang-todo-api/repository"
)

type TodoService interface {
	GetAllTodos(userID string) (*[]helper.TodoResponse, error)
	CreateTodo(todoRequest dto.CreateTodoRequest, userID string) (*helper.TodoResponse, error)
	UpdateTodo(updateTodoRequest dto.UpdateTodoRequest, userID string) (*helper.TodoResponse, error)
	FindOneTodoByID(todoID string) (*helper.TodoResponse, error)
	DeleteTodo(todoID string, userID string) error
}

type todoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(todoRepository repository.TodoRepository) TodoService {
	return &todoService{
		todoRepository: todoRepository,
	}
}



func (c *todoService)GetAllTodos(userID string) (*[]helper.TodoResponse, error)
func (c *todoService)CreateTodo(todoRequest dto.CreateTodoRequest, userID string) (*helper.TodoResponse, error)
func (c *todoService)UpdateTodo(updateTodoRequest dto.UpdateTodoRequest, userID string) (*helper.TodoResponse, error)
func (c *todoService)FindOneTodoByID(todoID string) (*helper.TodoResponse, error)
func (c *todoService)DeleteTodo(todoID string, userID string) error
