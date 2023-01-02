package service

import (
	"errors"
	"log"

	"github.com/KuraoHikari/golang-todo-api/dto"
	"github.com/KuraoHikari/golang-todo-api/entity"
	"github.com/KuraoHikari/golang-todo-api/helper"
	"github.com/KuraoHikari/golang-todo-api/repository"
	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(registerRequest dto.RegisterRequest) (*helper.UserResponse, error)
	UpdateUser(updateUserRequest dto.UpdateUserRequest) (*helper.UserResponse,error)
	FindUserByEmail(email string) (*helper.UserResponse, error)
	FindUserByID(userID string) (*helper.UserResponse, error)
	VerifyCredential(email string, password string) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}


func (c *userService)CreateUser(registerRequest dto.RegisterRequest) (*helper.UserResponse, error){
	user, err := c.userRepository.FindByEmail(registerRequest.Email)
	if err == nil {
		return nil, errors.New("user already exists")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	err = smapping.FillStruct(&user, smapping.MapFields(&registerRequest))
	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}
	user, _ = c.userRepository.InsertUser(user)
	res := helper.NewUserResponse(user)
	return &res, nil
}

func (c *userService)UpdateUser(updateUserRequest dto.UpdateUserRequest) (*helper.UserResponse,error){
	user := entity.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&updateUserRequest))
	if err != nil {
		return nil, err
	}
	user, err = c.userRepository.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	res := helper.NewUserResponse(user)
	return &res, nil
}

func (c *userService)FindUserByEmail(email string) (*helper.UserResponse, error){
	user, err := c.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	userResponse := helper.NewUserResponse(user)
	return &userResponse, nil
}

func (c *userService)FindUserByID(userID string) (*helper.UserResponse, error){
	user, err := c.userRepository.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	userResponse := helper.UserResponse{}
	err = smapping.FillStruct(&userResponse, smapping.MapFields(&user))
	if err != nil {
		return nil, err
	}
	return &userResponse, nil
}

func (c *userService)VerifyCredential(email string, password string) error{
	user, err := c.userRepository.FindByEmail(email)
	if err != nil {
		println("hehe")
		println(err.Error())
		return err
	}
	isValidPassword := helper.CompareHashAndPassword(user.Password, []byte(password))
	if !isValidPassword {
		return errors.New("failed to login. check your credential")
	}
	return nil
}

