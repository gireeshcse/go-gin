package main

import (
	"io"
	"net/http"
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

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	apiRoutes := server.Group("/api")
	{

		apiRoutes.GET("/users", func(ctx *gin.Context) {
			ctx.JSON(200, userController.FindAll())
		})

		apiRoutes.POST("/users", func(ctx *gin.Context) {
			user, err := userController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(200, user)
			}
		})

	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/users", userController.ShowAll)
	}

	server.Run(":8080")
}
