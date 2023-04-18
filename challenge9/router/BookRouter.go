package router

import (
	"challenge9/controller"

	"github.com/gin-gonic/gin"
)

func AddBookRouter(r *gin.Engine) {
	bookRouter := r.Group("/book")
	{
		bookRouter.GET("/", controller.GetAllBook)
		bookRouter.GET("/:id", controller.GetBookById)
		bookRouter.POST("/", controller.CreateBook)
		bookRouter.PUT("/:id", controller.UpdateBook)
		bookRouter.DELETE("/:id", controller.DeleteBook)
	}
}
