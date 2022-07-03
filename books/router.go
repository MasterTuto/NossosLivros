package books

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func getBooks(c *gin.Context) {
	fmt.Println("Hello, world")
}

func getBook(c *gin.Context) {
	fmt.Println("Hello, world")
}

func registerBook(c *gin.Context) {
	fmt.Println("Hello, world")
}

func updateBook(c *gin.Context) {
	fmt.Println("Hello, world")
}

func deleteBook(c *gin.Context) {
	fmt.Println("Hello, world")
}

func BooksRouter(r *gin.Engine) {
	r.GET("/books", getBooks)
	r.GET("/books/:bookId", getBook)
	r.POST("/books", registerBook)
	r.PATCH("/books", updateBook)
	r.DELETE("/books", deleteBook)
}
