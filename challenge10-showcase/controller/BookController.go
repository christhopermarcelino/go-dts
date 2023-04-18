package controller

import (
	"challenge10/config"
	"challenge10/model"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllBooks godoc
// @Summary Get All Books
// @Description Get All Books
// @Tags Book
// @Accept json
// @Produce json
// @Success 200 {array} model.Book
// @Router /book [get]
func GetAllBook(c *gin.Context) {
	db := config.GetDB()

	var books []model.Book

	err := db.Find(&books).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, books)
}

// GetBookById godoc
// @Summary Get Single Book
// @Description Get Single Book by Id
// @Tags Book
// @Accept json
// @Produce json
// @Param id path int true "ID of the book"
// @Success 200 {object} model.Book
// @Router /book/{id} [get]
func GetBookById(c *gin.Context) {
	db := config.GetDB()

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	var book model.Book

	err = db.First(&book, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Book not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get book",
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

// CreateBook godoc
// @Summary Create Book
// @Description Create Book
// @Tags Book
// @Accept json
// @Produce json
// @Param model.Book body model.Book true "Book data"
// @Success 201 {object} model.Book
// @Router /book [post]
func CreateBook(c *gin.Context) {
	db := config.GetDB()

	var book model.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := db.Create(&book).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create book",
		})
		return
	}

	c.JSON(http.StatusCreated, book)
}

// UpdateBook godoc
// @Summary Update Book
// @Description Update Book by Id
// @Tags Book
// @Accept json
// @Produce json
// @Param model.Book body model.Book true "Book data"
// @Param id path int true "ID of the book"
// @Success 200 {object} model.Book
// @Router /book:/id [put]
func UpdateBook(c *gin.Context) {
	db := config.GetDB()

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	var inputBook *model.Book
	if err = c.ShouldBindJSON(&inputBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var book *model.Book
	err = db.First(&book, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Book not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update book",
		})
		return
	}

	book.NameBook = inputBook.NameBook
	book.Author = inputBook.Author

	err = db.Save(&book).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update book",
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

// DeleteBook godoc
// @Summary Delete Book
// @Description Delete Book by Id
// @Tags Book
// @Accept json
// @Produce json
// @Param id path int true "ID of the book"
// @Success 200 {string} string "Book deleted successfully"
// @Router /book:/id [delete]
func DeleteBook(c *gin.Context) {
	db := config.GetDB()

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	var book model.Book
	err = db.First(&book, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Book not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete book",
		})
		return
	}

	err = db.Delete(&book).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete book",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})
}
