package loans

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

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

func LoansRouter(r *gin.Engine) {
	r.GET("/user/:userId/loans", getUserLoans)
	r.GET("/user/:userId/loans/:loanId", getUserLoanData)
	r.POST("/user/:userId/loans", lendBook)
	r.PATCH("/user/:userId/loans/:loanId", updateUserLoan)
}
