package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gireeshcse/go-gin/entity"
	"github.com/gireeshcse/go-gin/service"
	"github.com/gireeshcse/go-gin/validators"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// UserController : Interface for controller which states methods in this controller
type UserController interface {
	FindAll() []entity.User
	Save(context *gin.Context) (entity.User, error)
	ShowAll(context *gin.Context)
	Login(context *gin.Context) (entity.User, error)
}

type controller struct {
	service service.UserService
}

// New : Create new UserController
func New(service service.UserService) UserController {
	validate = validator.New()
	validate.RegisterValidation("containsUser", validators.ValidateUsername)
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.User {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) (entity.User, error) {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return user, err
	}

	// to validate custom validation
	err = validate.Struct(user)
	if err != nil {
		return user, err
	}

	// Hashing Password before saving
	hash, err := entity.GeneratePassword(entity.DefaultPasswordConfig, user.Password)
	if err != nil {
		// handle error
		panic(err)
	}
	user.Password = hash

	c.service.Save(user)
	user.Password = ""
	// return user
	return user, nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	users := c.service.FindAll()
	data := gin.H{
		"title": "Users Page",
		"users": users,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}

func (c *controller) Login(ctx *gin.Context) (entity.User, error) {
	var user entity.User
	var login entity.Login

	err := ctx.ShouldBindJSON(&login)

	if err != nil {
		return user, err
	}

	user, err = c.service.Find(login.Username)

	if err == nil {
		valid, _ := entity.ComparePassword(login.Password, user.Password)
		if valid == true {
			return user, nil
		}
	}

	return user, errors.New("Invalid username or password")

}
