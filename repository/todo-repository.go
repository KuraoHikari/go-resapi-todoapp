package repository

import (
	"github.com/KuraoHikari/golang-todo-api/entity"
	"gorm.io/gorm"
)
type TodoRepository interface {
	GetAllUserTodo(userID string) ([]entity.Todo, error)
	InsertTodo(todo entity.Todo)(entity.Todo, error)
	UpdateTodo(todo entity.Todo)(entity.Todo, error)
	DeleteTodo(todoID string) error
	FindOneTodoByID(ID string) (entity.Todo, error)
	FindAllTodo(userID string) ([]entity.Todo, error)
}

type todoRepository struct {
	connection *gorm.DB
}
func NewTodoRepo(connection *gorm.DB) TodoRepository {
	return &todoRepository{
		connection: connection,
	}
}

func (c *todoRepository)GetAllUserTodo(userID string) ([]entity.Todo, error){
	todos := []entity.Todo{}
	c.connection.Preload("User").Where("user_id = ?", userID).Find(&todos)
	return todos, nil
}
func (c *todoRepository)InsertTodo(todo entity.Todo)(entity.Todo, error){
	c.connection.Save(&todo)
	c.connection.Preload("User").Find(&todo)
	return todo, nil
}
func (c *todoRepository)UpdateTodo(todo entity.Todo)(entity.Todo, error){
	c.connection.Save(&todo)
	c.connection.Preload("User").Find(&todo)
	return todo, nil
}
func (c *todoRepository)DeleteTodo(todoID string) error{
	var todo entity.Todo
	res := c.connection.Preload("User").Where("id = ?", todoID).Take(&todo)
	if res.Error != nil {
		return res.Error
	}
	c.connection.Delete(&todo)
	return nil
}
func (c *todoRepository)FindOneTodoByID(ID string) (entity.Todo, error){
	var todo entity.Todo
	res := c.connection.Preload("User").Where("id = ?", ID).Take(&todo)
	if res.Error != nil {
		return todo, res.Error
	}
	return todo, nil
}
func (c *todoRepository)FindAllTodo(userID string) ([]entity.Todo, error){
	todos := []entity.Todo{}
	c.connection.Where("user_id = ?", userID).Find(&todos)
	return todos, nil
}