package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/KuraoHikari/golang-todo-api/dto"
	"github.com/KuraoHikari/golang-todo-api/helper"
	"github.com/KuraoHikari/golang-todo-api/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	Profile(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type userController struct {
	jwtService  helper.JWTService
	userService service.UserService
}

func NewUserController(
	jwtService helper.JWTService,
	userService service.UserService,
) UserController {
	return &userController{	
		jwtService:  jwtService,
		userService: userService,
	}
}

func (c *userController) Login(ctx *gin.Context) {
	var loginRequest dto.LoginRequest
	err := ctx.ShouldBind(&loginRequest)

	if err != nil {
		response := helper.BuildErrorResponse("failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = c.userService.VerifyCredential(loginRequest.Email, loginRequest.Password)
	if err != nil {
		response := helper.BuildErrorResponse("failed to login", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	user, _ := c.userService.FindUserByEmail(loginRequest.Email)

	token := c.jwtService.GenerateToken(strconv.FormatInt(user.ID, 10))
	user.Token = token
	response := helper.BuildResponse(true, "OK!", user)
	ctx.JSON(http.StatusOK, response)

}

func (c *userController) Register(ctx *gin.Context) {
	var registerRequest dto.RegisterRequest

	err := ctx.ShouldBind(&registerRequest)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, err := c.userService.CreateUser(registerRequest)
	if err != nil {
		response := helper.BuildErrorResponse(err.Error(), err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	token := c.jwtService.GenerateToken(strconv.FormatInt(user.ID, 10))
	user.Token = token
	response := helper.BuildResponse(true, "OK!", user)
	ctx.JSON(http.StatusCreated, response)

}

func (c *userController) getUserIDByHeader(ctx *gin.Context) string {
	header := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(header, ctx)

	if token == nil {
		response := helper.BuildErrorResponse("Error", "Failed to validate token", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return ""
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}

func (c *userController) Update(ctx *gin.Context) {
	var updateUserRequest dto.UpdateUserRequest

	err := ctx.ShouldBind(&updateUserRequest)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id := c.getUserIDByHeader(ctx)

	if id == "" {
		response := helper.BuildErrorResponse("Error", "Failed to validate token", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	_id, _ := strconv.ParseInt(id, 0, 64)
	updateUserRequest.ID = _id
	res, err := c.userService.UpdateUser(updateUserRequest)

	if err != nil {
		response := helper.BuildErrorResponse("Error", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	response := helper.BuildResponse(true, "OK", res)
	ctx.JSON(http.StatusOK, response)

}

func (c *userController) Profile(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(header, ctx)

	if token == nil {
		response := helper.BuildErrorResponse("Error", "Failed to validate token", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user, err := c.userService.FindUserByID(id)

	if err != nil {
		response := helper.BuildErrorResponse("Error", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	res := helper.BuildResponse(true, "OK", user)
	ctx.JSON(http.StatusOK, res)
}
