package books

import (
	"fmt"
	"net/http"
	"strconv"

	"brenocs.dev/nossolivros/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func getBooksFromUser(c *gin.Context) {
	userDb, found := c.Get("db")
	db := userDb.(*gorm.DB)

	if !found {
		c.JSON(500, gin.H{"error": "could not get db"})
		return
	}

	var user users.User
	userId := c.Param("userId")
	limit := c.Query("limit")
	page := c.Query("page")

	db.Take(&user, userId)

	if (user == users.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	limitNumber, err := strconv.Atoi(limit)
	if err != nil {
		limitNumber = 10
	}

	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		pageNumber = 1
	}

	var offset = limitNumber * (pageNumber - 1)

	var books []Stored
	var total int64
	db.Where("user_id = ?", user.ID).Limit(limitNumber).Offset(offset).Find(&books).Count(&total)

	c.JSON(http.StatusOK, gin.H{
		"result": books,
		"total":  total,
	})
}

func removeUserBookRegister(c *gin.Context) {
	readDb, found := c.Get("db")
	db := readDb.(*gorm.DB)

	if !found {
		c.JSON(500, gin.H{"error": "could not get db"})
		return
	}

	userId := c.Param("userId")
	bookId := c.Param("bookId")

	var user users.User
	db.Take(&user, userId)

	if (user == users.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	var book Stored
	db.Take(&book, bookId)

	if (book == Stored{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
		return
	}

	db.Delete(&book)
}

func registerBookToUser(c *gin.Context) {
	fmt.Println("Hello, world")
}

func BooksRouter(r *gin.Engine) {
	r.GET("/books", getBooks)
	r.GET("/books/:bookId", getBook)
	r.POST("/books", registerBook)
	r.PATCH("/books", updateBook)
	r.DELETE("/books", deleteBook)

	r.GET("/user/:userId/books", getBooksFromUser)
	r.DELETE("/user/:userId/books/:bookId", removeUserBookRegister)
	r.POST("/user/:userId/books", registerBookToUser)
}
