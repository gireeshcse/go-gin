package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userStructure struct {
	username string
	password string
	name     string
}

var userDB = make(map[string]userStructure)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	r.GET("/hello/:username", func(c *gin.Context) {
		username := c.Params.ByName("username")
		value, ok := userDB[username]
		var message string
		if ok {
			message = "Welcome back " + value.name
		} else {
			message = "Hello Guest " + value.name
		}
		c.JSON(http.StatusOK, gin.H{"name": username, "message": message})
	})

	r.POST("/hello/register", func(c *gin.Context) {

		// Parse JSON
		var json struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
			Name     string `json:"name" binding:"required"`
		}
		if c.Bind(&json) == nil {
			_, ok := userDB[json.Username]
			if ok {
				c.JSON(http.StatusOK, gin.H{"status": "Already Registered"})
			} else {
				userDB[json.Username] = userStructure{username: json.Username, password: json.Password, name: json.Name}
				c.JSON(http.StatusOK, gin.H{"status": "OK"})
			}

		} else {
			c.JSON(http.StatusOK, gin.H{"status": "Not OK"})
		}
	})

	return r
}
