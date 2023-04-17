package router

import (
	"github.com/gin-gonic/gin"

	"book-api/controller"
)

func AddBookRouter(r *gin.Engine) {
	bookRouter := r.Group("/book")
	{
		bookRouter.GET("/", controller.GetAllBook)
		bookRouter.GET("/:id", controller.GetBookById)
		bookRouter.POST("/", controller.AddBook)
		bookRouter.PUT("/:id", controller.UpdateBook)
		bookRouter.DELETE("/:id", controller.DeleteBook)
	}
}
