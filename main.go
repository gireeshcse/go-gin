package main

import (
	"fmt"

	"github.com/gireeshcse/go-gin/web"
)

func main() {
	fmt.Println("Hello, world.")
	r := web.SetupRouter()
	r.Run(":8080")
}
