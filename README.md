### Go-Gin Project


#### Initial Setup

##### Project Setup

```
mkdir go-gin
cd go-gin
go mod init github.com/gireeshcse/go-gin
```

creates go.mod with the following kind of content

```
module github.com/gireeshcse/go-gin

go 1.15

```

##### Libraries/Packages used

```
go get -u github.com/gin-gonic/gin
go get github.com/dgrijalva/jwt-go
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -v golang.org/x/crypto/argon2

```
##### Build and install the program

```
go install github.com/gireeshcse/go-gin
go install .
go install
```
Creates an executable binary and then install that binary in **$HOME/go/bin/go-gin**


#### importing packages from our own module

```
mkdir web
vi web/routes.go
```

web/routes.go

```
package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	return r
}
```

```
cd web
go build
```

#### Creating Database and User 

```
CREATE DATABASE go_app;

GRANT ALL PRIVILEGES ON go_app.* TO 'go_app_user'@'localhost' IDENTIFIED BY 'app123' WITH GRANT OPTION;
```

### Important

[For Logging Request Header Debugging](https://github.com/tpkeeper/gin-dump)


#### Credits 

[pragmaticreviews](https://pragmaticreviews.com)