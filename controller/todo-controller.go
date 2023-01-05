package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/KuraoHikari/golang-todo-api/dto"
	"github.com/KuraoHikari/golang-todo-api/helper"
	"github.com/KuraoHikari/golang-todo-api/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type TodoController interface {
	GetAllTodo(ctx *gin.Context)
	CreateTodo(ctx *gin.Context)
	UpdateTodo(ctx *gin.Context)
	DeleteTodo(ctx *gin.Context)
	FindOneTodoByID(ctx *gin.Context)
}

type todoController struct {
	todoService service.TodoService
	jwtService  helper.JWTService
}

func NewTodoController(todoService service.TodoService, jwtService helper.JWTService) TodoController {
	return &todoController{
		todoService: todoService,
		jwtService:  jwtService,
	}
}

func (c *todoController) GetAllTodo(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	todos, err := c.todoService.GetAllTodos(userID)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse(true, "OK!", todos)
	ctx.JSON(http.StatusOK, response)
}

func (c *todoController) CreateTodo(ctx *gin.Context) {
	var createTodoReq dto.CreateTodoRequest
	err := ctx.ShouldBind(&createTodoReq)

	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	res, err := c.todoService.CreateTodo(createTodoReq, userID)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusCreated, response)

}

func (c *todoController) FindOneTodoByID(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.todoService.FindOneTodoByID(id)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *todoController) DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	err := c.todoService.DeleteTodo(id, userID)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildResponse(true, "OK!", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}

func (c *todoController) UpdateTodo(ctx *gin.Context) {
	updateTodoRequest := dto.UpdateTodoRequest{}
	err := ctx.ShouldBind(&updateTodoRequest)

	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	id, _ := strconv.ParseInt(ctx.Param("id"), 0, 64)
	updateTodoRequest.ID = id
	todo, err := c.todoService.UpdateTodo(updateTodoRequest, userID)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.BuildResponse(true, "OK!", todo)
	ctx.JSON(http.StatusOK, response)

}