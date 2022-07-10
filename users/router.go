package users

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func createUser(c *gin.Context) {
	varDb, found := c.Get("db")
	var db *gorm.DB = varDb.(*gorm.DB)

	if !found {
		c.JSON(500, gin.H{"error": "could not get db"})
		return
	}

	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "username, password or username are missing."})
		return
	}

	if len(user.Password) < 8 {
		c.JSON(400, gin.H{
			"error": "Missing username, email or password",
		})
	}

	db.Create(user)

	c.JSON(200, gin.H{
		"result": gin.H{
			"id":       user.ID,
			"username": user.Name,
			"email":    user.Email,
		},
	})
}

func logUserIn(c *gin.Context) {
	fmt.Println("Hello, world")
}

func updateUserData(c *gin.Context) {
	fmt.Println("Hello, world")
}

func getBooksFromUser(c *gin.Context) {
	fmt.Println("Hello, world")
}

func updateUserBook(c *gin.Context) {
	fmt.Println("Hello, world")
}

func removeUserBookRegister(c *gin.Context) {
	fmt.Println("Hello, world")
}

func registerBookToUser(c *gin.Context) {
	fmt.Println("Hello, world")
}

func getUserLoans(c *gin.Context) {
	fmt.Println("Hello, world")
}

func getUserLoanData(c *gin.Context) {
	fmt.Println("Hello, world")
}

func lendBook(c *gin.Context) {
	fmt.Println("Hello, world")
}

func updateUserLoan(c *gin.Context) {
	fmt.Println("Hello, world")
}

func UserRoutes(r *gin.Engine) {
	r.POST("/user/signup", createUser)
	r.POST("/user/signin", logUserIn)

	r.PATCH("/user", updateUserData)

	r.GET("/user/:userId/books", getBooksFromUser)
	r.PATCH("/user/:userId/books/:bookId", updateUserBook)
	r.DELETE("/user/:userId/books/:bookId", removeUserBookRegister)
	r.POST("/user/:userId/books", registerBookToUser)

	r.GET("/user/:userId/loans", getUserLoans)
	r.GET("/user/:userId/loans/:loanId", getUserLoanData)
	r.POST("/user/:userId/loans", lendBook)
	r.PATCH("/user/:userId/loans/:loanId", updateUserLoan)

}
