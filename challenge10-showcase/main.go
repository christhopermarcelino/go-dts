package main

import (
	"challenge10/config"
	"challenge10/router"

	_ "challenge10/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	config.InitDatabase()
}

// @title           Book Swagger API
// @version         1.0
// @description     This is documentation of book API
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /
func main() {
	PORT := ":8000"

	r := gin.Default()

	router.AddBookRouter(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(PORT)
}
