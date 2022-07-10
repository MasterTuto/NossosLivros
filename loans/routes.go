package loans

import (
	"fmt"
	"net/http"
	"strconv"

	"brenocs.dev/nossolivros/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getUserLoans(c *gin.Context) {
	readDb, found := c.Get("db")

	if !found {
		c.JSON(500, gin.H{"error": "could not get db"})
		return
	}

	db := readDb.(*gorm.DB)

	userId := c.Param("userId")
	limit := c.Param("limit")
	page := c.Param("page")

	var user users.User

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

	offset := limitNumber * (pageNumber - 1)

	var loans []Loan
	var total int64

	db.Where("user_id = ?", userId).Offset(offset).Limit(limitNumber).Find(&loans).Count(&total)

	c.JSON(http.StatusOK, gin.H{
		"result": loans,
		"total":  total,
	})
}

func getUserLoanData(c *gin.Context) {
	readDb, found := c.Get("db")

	if !found {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get db"})
		return
	}

	db := readDb.(*gorm.DB)

	userId := c.Param("userId")
	loanId := c.Param("loanId")

	var user users.User

	db.Take(&user, userId)

	if (user == users.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	var loan Loan

	db.Take(&loan, loanId)

	if (loan == Loan{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "loan not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": loan,
	})
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
