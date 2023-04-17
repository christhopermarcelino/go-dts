package main

import (
	"challenge8/config"
	"challenge8/model"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func main() {
	db = config.ConnectDB()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/book", GetBooks)
	r.GET("/book/:id", GetBook)
	r.POST("/book", CreateBook)
	r.PUT("/book/:id", UpdateBook)
	r.DELETE("/book/:id", DeleteBook)

	r.Run(":8080")
}

func GetBooks(c *gin.Context) {
	books := []model.Book{}

	sqlStatement := "SELECT * FROM books"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		book := model.Book{}

		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Description)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		books = append(books, book)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func GetBook(c *gin.Context) {
	id := c.Param("id")

	sqlStatement := "SELECT * FROM books WHERE id = $1"
	row := db.QueryRow(sqlStatement, id)

	book := model.Book{}
	err := row.Scan(&book.Id, &book.Title, &book.Author, &book.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Book not found",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func CreateBook(c *gin.Context) {
	inputBook := model.Book{}

	if err := c.ShouldBindJSON(&inputBook); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	sqlStatement := `
	INSERT INTO books (title, author, description)
	VALUES ($1, $2, $3)
	Returning *
	`

	var book model.Book
	err := db.QueryRow(sqlStatement, inputBook.Title, inputBook.Author, inputBook.Description).
		Scan(&book.Id, &book.Title, &book.Author, &book.Description)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": book,
	})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	inputBook := model.Book{}

	if err := c.ShouldBindJSON(&inputBook); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	sqlStatement := `
	UPDATE books
	SET title = $2, author = $3, description = $4
	WHERE id = $1;
	`

	res, err := db.Exec(sqlStatement, id, inputBook.Title, inputBook.Author, inputBook.Description)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	count, err := res.RowsAffected()
	if err != nil || count == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book updated successfully",
	})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	sqlStatement := "DELETE FROM books WHERE id = $1;"

	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	count, err := res.RowsAffected()
	if err != nil || count == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})
}
