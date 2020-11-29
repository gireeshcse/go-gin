package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gireeshcse/go-gin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLDB to connect to the database
func MySQLDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
		// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
		dsn := config.DBUSER + ":" + config.DBPASSWORD + "@tcp(" + config.DBHOST + ":" + config.DBPORT + ")/" + config.DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unable to connect to database"})
		} else {
			database, err := db.DB()

			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unable to connect to database"})
			} else {

				c.Set("db", db)
				c.Next()
				database.Close()
			}
		}

	}
}
