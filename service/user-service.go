package service

import (
	"errors"

	"github.com/gireeshcse/go-gin/entity"
)

// UserService : to create a interface for user service
type UserService interface {
	Save(entity.User) entity.User
	FindAll() []entity.User
	Find(username string) (entity.User, error)
}

type userService struct {
	users []entity.User
}

// New : to create user service
func New() UserService {
	return &userService{
		users: []entity.User{},
	}
}

func (service *userService) Save(user entity.User) entity.User {
	service.users = append(service.users, user)
	return user
}
func (service *userService) FindAll() []entity.User {
	return service.users
}

func (service *userService) Find(username string) (entity.User, error) {
	users := service.users
	var user1 entity.User
	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}
	return user1, errors.New("No user found with the username: " + username)
}
