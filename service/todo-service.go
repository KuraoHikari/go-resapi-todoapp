package service

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/KuraoHikari/golang-todo-api/dto"
	"github.com/KuraoHikari/golang-todo-api/entity"
	"github.com/KuraoHikari/golang-todo-api/helper"
	"github.com/KuraoHikari/golang-todo-api/repository"
	"github.com/mashingan/smapping"
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



func (c *todoService)GetAllTodos(userID string) (*[]helper.TodoResponse, error){
	todos, err := c.todoRepository.GetAllUserTodo(userID)
	if err != nil {
		return nil, err
	}

	allTodos := helper.NewTodoArrayResponse(todos)
	return &allTodos, nil
}

func (c *todoService)CreateTodo(todoRequest dto.CreateTodoRequest, userID string) (*helper.TodoResponse, error){
	todo := entity.Todo{}
	err := smapping.FillStruct(&todo, smapping.MapFields(&todoRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	id, _ := strconv.ParseInt(userID, 0, 64)
	todo.UserID = id
	p, err := c.todoRepository.InsertTodo(todo)
	if err != nil {
		return nil, err
	}

	res := helper.NewTodoResponse(p)
	return &res, nil
}

func (c *todoService)UpdateTodo(updateTodoRequest dto.UpdateTodoRequest, userID string) (*helper.TodoResponse, error){
	todo, err := c.todoRepository.FindOneTodoByID(fmt.Sprintf("%d", updateTodoRequest.ID))
	if err != nil {
		return nil, err
	}

	uid, _ := strconv.ParseInt(userID, 0, 64)
	if todo.UserID != uid {
		return nil, errors.New("task ini bukan milik anda")
	}

	todo = entity.Todo{}
	err = smapping.FillStruct(&todo, smapping.MapFields(&updateTodoRequest))

	if err != nil {
		return nil, err
	}

	todo.UserID = uid
	todo, err = c.todoRepository.UpdateTodo(todo)

	if err != nil {
		return nil, err
	}

	res := helper.NewTodoResponse(todo)
	return &res, nil
}

func (c *todoService)FindOneTodoByID(todoID string) (*helper.TodoResponse, error){
	todo, err := c.todoRepository.FindOneTodoByID(todoID)

	if err != nil {
		return nil, err
	}

	res := helper.NewTodoResponse(todo)
	return &res, nil
}

func (c *todoService)DeleteTodo(todoID string, userID string) error {
	todo, err := c.todoRepository.FindOneTodoByID(todoID)
	if err != nil {
		return err
	}

	if fmt.Sprintf("%d", todo.UserID) != userID {
		return errors.New("task ini bukan milik anda")
	}

	c.todoRepository.DeleteTodo(todoID)
	return nil
}
