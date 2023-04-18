package router

import (
	"challenge10/controller"

	"github.com/gin-gonic/gin"
)

func AddBookRouter(r *gin.Engine) {
	bookRouter := r.Group("/api/book")
	{
		// Get All Books
		bookRouter.GET("/", controller.GetAllBook)
		// Get Book
		bookRouter.GET("/:id", controller.GetBookById)
		// Create Book
		bookRouter.POST("/", controller.CreateBook)
		// Update Book
		bookRouter.PUT("/:id", controller.UpdateBook)
		// Delete Book
		bookRouter.DELETE("/:id", controller.DeleteBook)
	}
}
