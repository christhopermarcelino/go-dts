package main

import (
	"github.com/gin-gonic/gin"

	"book-api/router"
)

func main() {
	PORT := ":8080"

	r := gin.Default()

	router.AddBookRouter(r)

	r.Run(PORT)
}
