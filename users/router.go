package users

import (
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userUpdatePatch = struct {
	name          string
	email         string
	password      string
	finishedLoans int
	loanLimit     int
}

func createUser(c *gin.Context) {
	varDb, found := c.Get("db")
	var db *gorm.DB = varDb.(*gorm.DB)

	if !found {
		c.JSON(500, gin.H{"error": "could not get db"})
		return
	}

	var user struct {
		name     string `binding:"required"`
		email    string `binding:"required"`
		password string `binding:"required"`
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "username, password or username are missing."})
		return
	}

	if len(user.password) < 8 {
		c.JSON(400, gin.H{
			"error": "Wrong password length",
		})
	}

	newUser := New(user.name, user.email, user.password)

	db.Create(&newUser)

	c.JSON(200, gin.H{
		"result": gin.H{
			"id":       newUser.ID,
			"username": newUser.Name,
			"email":    newUser.Email,
		},
	})
}

func logUserIn(c *gin.Context) {
	fmt.Println("Hello, world")
}

func updateUserData(c *gin.Context) {
	userDb, found := c.Get("db")
	db := userDb.(*gorm.DB)

	if !found {
		c.JSON(500, gin.H{"error": "could not get db"})
		return
	}

	userId := c.Param("userId")

	var user User

	db.First(&user, userId)

	if (user == User{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	var requestBody userUpdatePatch = userUpdatePatch{
		name:          "",
		email:         "",
		password:      "",
		finishedLoans: -1,
		loanLimit:     -1,
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "wrong request body"})
		return
	}

	updates := map[string]interface{}{}

	if requestBody.name != "" {
		updates["name"] = requestBody.name
	}

	if requestBody.email != "" {
		updates["email"] = requestBody.email
	}

	if requestBody.password != "" {
		updates["password"] = fmt.Sprintf("%x", sha256.Sum256([]byte(requestBody.password)))
	}

	if requestBody.finishedLoans != -1 {
		updates["finishedLoans"] = requestBody.finishedLoans
	}

	if requestBody.loanLimit != -1 {
		updates["loanLimit"] = requestBody.loanLimit
	}

	db.Model(&user).Updates(updates)
}

func UserRoutes(r *gin.Engine) {
	r.POST("/user/signup", createUser)
	r.POST("/user/signin", logUserIn)

	r.PATCH("/user/:userId", updateUserData)
}
