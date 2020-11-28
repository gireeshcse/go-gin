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

```
##### Build and install the program

```
go install github.com/gireeshcse/go-gin
go install .
go install
```
Creates an executable binary and then install that binary in **$HOME/go/bin/go-gin**
