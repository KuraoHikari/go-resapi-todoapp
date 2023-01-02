package repository

import (
	"github.com/KuraoHikari/golang-todo-api/entity"
	"github.com/KuraoHikari/golang-todo-api/helper"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user entity.User)(entity.User, error)
	UpdateUser(user entity.User)(entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindByUserID(userID string) (entity.User, error)
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository(connection *gorm.DB) UserRepository{
	return &userRepository{
		connection: connection,
	}
}

func(c *userRepository)	InsertUser(user entity.User)(entity.User, error){
	user.Password = helper.HashAndSalt([]byte(user.Password))
	c.connection.Save(&user)
	return user, nil
}
func(c *userRepository)	UpdateUser(user entity.User)(entity.User, error){
	if user.Password != "" {
		user.Password = helper.HashAndSalt([]byte(user.Password))
	} else {
		var tempUser entity.User
		c.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	c.connection.Save(&user)
	return user, nil
}
func(c *userRepository)	FindByEmail(email string) (entity.User, error){
	var user entity.User
	res := c.connection.Where("email = ?", email).Take(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}
func(c *userRepository)	FindByUserID(userID string) (entity.User, error){
	var user entity.User
	res := c.connection.Where("id = ?", userID).Take(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}
