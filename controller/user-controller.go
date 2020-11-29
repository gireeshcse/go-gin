package controller

import (
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

	c.service.Save(user)
	// return user
	return user, nil
}
