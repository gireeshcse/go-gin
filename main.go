package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gireeshcse/go-gin/controller"
	"github.com/gireeshcse/go-gin/service"
	"github.com/gireeshcse/go-gin/web"
)

var (
	userService    service.UserService       = service.New()
	userController controller.UserController = controller.New(userService)
)

func main() {
	fmt.Println("Hello, world.")
	r := web.SetupRouter()

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, userController.FindAll())
	})

	r.POST("/users", func(ctx *gin.Context) {
		ctx.JSON(200, userController.Save(ctx))
	})

	r.Run(":8080")
}
