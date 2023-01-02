package main

import (
	"net/http"

	"github.com/KuraoHikari/golang-todo-api/config"
	"github.com/KuraoHikari/golang-todo-api/helper"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)
var db *gorm.DB = config.SetupDatabaseConnection()

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.New()
	server.Use(gin.LoggerWithFormatter(helper.LoggerConsole))
	server.Use(gin.Recovery())
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		  "message": "pong",
		})
	})
	server.Run()
}