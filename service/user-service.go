package service

import (
	"github.com/KuraoHikari/golang-todo-api/dto"
)

type UserService interface {
	CreateUser(registerRequest dto.RegisterRequest) error
	UpdateUser(updateUserRequest dto.UpdateUserRequest) error
}

