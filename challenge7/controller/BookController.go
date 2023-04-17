package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id          string `json:"id"`
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Description string `json:"description" binding:"required"`
}

var BookDatas = []Book{}
var indexCounter = 1

func GetAllBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": BookDatas,
	})
}

func GetBookById(c *gin.Context) {
	var (
		book   Book
		exists bool
		id     string
	)
	id = c.Param("id")

	for _, b := range BookDatas {
		if b.Id == id {
			book = b
			exists = true
			break
		}
	}

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Book not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func AddBook(c *gin.Context) {
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid data! Name, author, and description must be filled.",
		})
		return
	}

	book.Id = fmt.Sprintf("b%d", indexCounter)
	BookDatas = append(BookDatas, book)
	indexCounter++
func GetBooks() {
	employees := []model.Book{}

	sqlStatement := "SELECT * FROM employees"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		employee := model.Book{}

		err = rows.Scan(&employee.ID, &employee.FullName, &employee.Email, &employee.Age, &employee.Division)
		if err != nil {
			panic(err)
		}

		employees = append(employees, employee)
	}

	fmt.Printf("Employee data: %v\n", employees)
}

func GetBookById() {
	
}

func CreateBook() {
	employee := model.Book{}

	sqlStatement := `
	INSERT INTO employees (age, division, email, full_name)
	VALUES ($1, $2, $3, $4)
	Returning *
	`

	err = db.QueryRow(sqlStatement, 64, "Sales", "rick@aws.com", "Rick").
		Scan(&employee.ID, &employee.FullName, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee: %+v\n", employee)
}

func UpdateBook() {
	sqlStatement := `
	UPDATE employees
	SET full_name = $2, email = $3, age = $4, division = $5
	WHERE id = $1;
	`

	res, err := db.Exec(sqlStatement, 7, "Rick Sanchez", "sanch@edu.com", 10, "CEO")
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Updated %d rows\n", count)
}

func DeleteBook() {
	sqlStatement := "DELETE FROM employees WHERE id = $1;"

	res, err := db.Exec(sqlStatement, 3)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Deleted %d rows\n", count)
}


	c.JSON(http.StatusCreated, gin.H{
		"message": "Book added successfully",
		"data":    book,
	})
}

func UpdateBook(c *gin.Context) {
	var (
		book   Book
		exists bool
		id     string
	)

	id = c.Param("id")

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid data! Name, author, and description must be filled.",
		})
		return
	}

	for i, b := range BookDatas {
		if b.Id == id {
			book.Id = BookDatas[i].Id
			BookDatas[i] = book
			exists = true
			break
		}
	}

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Book not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book updated successfully",
		"data":    book,
	})
}

func DeleteBook(c *gin.Context) {
	var (
		exists    bool
		id        string
		bookIndex int
	)

	id = c.Param("id")

	for i, b := range BookDatas {
		if b.Id == id {
			bookIndex = i
			exists = true
			break
		}
	}

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Book not found",
		})
		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	c.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})
}
