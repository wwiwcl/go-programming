package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"slices"
)

type Book struct {
	// TODO: Finish struct
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Pages int    `json:"pages"`
}

var bookshelf = []Book{
	// TODO: Init bookshelf
	{1, "Blue Bird", 500},
}

func getBooks(c *gin.Context) {
    c.JSON(200, bookshelf)
}
func getBook(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
    for _, book := range bookshelf {
        if id == book.Id {
            c.JSON(200, book)
            return
        }
    }
    c.JSON(404, gin.H{"message": "book not found"})
}
func addBook(c *gin.Context) {
	var newBook Book
    err := c.BindJSON(&newBook)
    if err != nil {
        return
    }
    for _, book := range bookshelf {
        if newBook.Name == book.Name {
            c.JSON(409, gin.H{"message": "duplicate book name"})
            return
        }
    }
    bookId++
    newBook.Id = bookId
    bookshelf = append(bookshelf, newBook)
    c.JSON(201, newBook)
}
func deleteBook(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
    for i, book := range bookshelf {
        if id == book.Id {
            bookshelf = slices.Delete(bookshelf, i, i+1)
            c.JSON(204, nil)
            return
        }
    }
    c.JSON(204, nil)
}
func updateBook(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
    var newBook Book
    if err := c.BindJSON(&newBook); err != nil {
        return
    }
    for _, book := range bookshelf {
        if newBook.Name == book.Name {
            c.JSON(409, gin.H{"message": "duplicate book name"})
            return
        }
    }
    for i, book := range bookshelf {
        if id == book.Id {
            book.Name = newBook.Name
            book.Pages = newBook.Pages
            bookshelf[i] = book
            c.JSON(200, bookshelf[i])
            return
        }
    }
    c.JSON(404, gin.H{"message": "book not found"})
}

func main() {
	r := gin.Default()
	r.RedirectFixedPath = true

	// TODO: Add routes
	r.GET("/bookshelf", getBooks)
    r.GET("/bookshelf/:id", getBook)
    r.POST("/bookshelf", addBook)
    r.DELETE("/bookshelf/:id", deleteBook)
    r.PUT("/bookshelf/:id", updateBook)

	err := r.Run(":8087")
	if err != nil {
		return
	}
}