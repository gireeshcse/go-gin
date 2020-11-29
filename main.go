package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gireeshcse/go-gin/model"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/gireeshcse/go-gin/config"
	"github.com/gireeshcse/go-gin/controller"
	"github.com/gireeshcse/go-gin/middlewares"
	"github.com/gireeshcse/go-gin/service"
)

var (
	userService    service.UserService       = service.New()
	userController controller.UserController = controller.New(userService)
	jwtService     service.JWTService        = service.NewJWTService(config.ISSUER, config.SECRET, config.EXPIRATIONTIME)
)

func setupLogOutput() {
	f, _ := os.Create("server.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	// r := web.SetupRouter()
	setupLogOutput()
	server := gin.New()
	model.InitiateApp()

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger())

	apiRoutes := server.Group("/api")
	{
		apiAuthRoutes := apiRoutes.Group("/admin")
		apiAuthRoutes.Use(middlewares.BasicAuth())
		{
			apiAuthRoutes.GET("/users", func(ctx *gin.Context) {
				ctx.JSON(200, userController.FindAll())
			})
		}

		apiRoutes.POST("/users", func(ctx *gin.Context) {
			user, err := userController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(200, user)
			}
		})

		apiRoutes.POST("/login", func(ctx *gin.Context) {
			user, err := userController.Login(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				var roles = []string{"Basic", "Authenticated"}
				token := jwtService.GenerateToken(user.Username, roles)
				ctx.JSON(200, gin.H{"user": user, "token": token})
			}
		})

		apiJWTAuthRoutes := apiRoutes.Group("/auth")
		apiJWTAuthRoutes.Use(middlewares.JWTAuth())
		{
			apiJWTAuthRoutes.GET("/user", func(ctx *gin.Context) {
				token, _ := ctx.Get("token")
				ctx.JSON(200, gin.H{"token": token})
			})
		}

		apiDBAuthRoutes := apiRoutes.Group("/db")
		apiDBAuthRoutes.Use(middlewares.MySQLDB())
		{
			apiDBAuthRoutes.GET("/test", func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{"db": "Connection to db is ok"})
			})
			// curl http://localhost:8080/api/db/users
			apiDBAuthRoutes.GET("/users", func(ctx *gin.Context) {
				db, _ := ctx.Get("db")
				database := db.(*gorm.DB)
				rows := map[string]interface{}{}
				database.Table("users").Find(&rows)
				ctx.JSON(200, gin.H{"db": rows})
			})
			// curl http://localhost:8080/api/db/migrate
			apiDBAuthRoutes.GET("/migrate", func(ctx *gin.Context) {
				db, _ := ctx.Get("db")
				database := db.(*gorm.DB)
				database.AutoMigrate(&model.User{})
				database.AutoMigrate(&model.Company{})

				ctx.JSON(200, gin.H{"db": "migration success"})
			})

			apiDBAuthRoutes.GET("/create", func(ctx *gin.Context) {
				db, _ := ctx.Get("db")
				database := db.(*gorm.DB)
				company := model.Company{Name: "Google", City: "California", State: "California", Country: "USA"}
				database.Create(&company)
				user := model.User{Name: "Ram", Username: "ram", Email: "ram@test.com", Password: "ram123", CompanyID: uint64(company.ID)}
				database.Create(&user)
				ctx.JSON(200, gin.H{"db": user})
			})
		}

	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/users", userController.ShowAll)
	}

	server.Run(":8080")
}
