package users

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func createUser(c *gin.Context) {
	fmt.Println("Create user!")
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
