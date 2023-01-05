package main

import (
	"github.com/KuraoHikari/golang-todo-api/config"
	"github.com/KuraoHikari/golang-todo-api/controller"
	"github.com/KuraoHikari/golang-todo-api/helper"
	"github.com/KuraoHikari/golang-todo-api/middleware"
	"github.com/KuraoHikari/golang-todo-api/repository"
	"github.com/KuraoHikari/golang-todo-api/service"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)
var (
	db *gorm.DB = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	todoRepository repository.TodoRepository = repository.NewTodoRepo(db)
	jwtService helper.JWTService = helper.NewJWTService()
	userService service.UserService = service.NewUserService(userRepository)
	todoService service.TodoService = service.NewTodoService(todoRepository)
	userController controller.UserController = controller.NewUserController(jwtService,userService)
	todoController controller.TodoController = controller.NewTodoController(todoService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.New()
	server.Use(gin.LoggerWithFormatter(helper.LoggerConsole))
	server.Use(gin.Recovery())
	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("/login", userController.Login)
		authRoutes.POST("/register", userController.Register)
	}
	userRoutes := server.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}
	todoRoutes := server.Group("api/todo", middleware.AuthorizeJWT(jwtService))
	{
		todoRoutes.GET("/", todoController.GetAllTodo)
		todoRoutes.POST("/", todoController.CreateTodo)
		todoRoutes.GET("/:id", todoController.FindOneTodoByID)
		todoRoutes.PUT("/:id", todoController.UpdateTodo)
		todoRoutes.DELETE("/:id", todoController.DeleteTodo)
	}
	server.Run()
}