package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gireeshcse/go-gin/controller"
	"github.com/gireeshcse/go-gin/middlewares"
	"github.com/gireeshcse/go-gin/service"
)

var (
	userService    service.UserService       = service.New()
	userController controller.UserController = controller.New(userService)
)

func setupLogOutput() {
	f, _ := os.Create("server.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	// r := web.SetupRouter()
	setupLogOutput()
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger())

	server.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, userController.FindAll())
	})

	server.POST("/users", func(ctx *gin.Context) {
		ctx.JSON(200, userController.Save(ctx))
	})

	server.Run(":8080")
}
