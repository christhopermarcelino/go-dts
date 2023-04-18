package main

import (
	"challenge9/config"
	"challenge9/router"

	"github.com/gin-gonic/gin"
)

func init() {
	config.InitDatabase()
}

func main() {
	PORT := ":8000"

	r := gin.Default()

	router.AddBookRouter(r)

	r.Run(PORT)
}
