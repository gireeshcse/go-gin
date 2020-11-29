package middlewares

import "github.com/gin-gonic/gin"

// BasicAuth - takes static usernames and passwords for the authentication
func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"admin": "admin123",
		"user":  "user123",
	})
}
