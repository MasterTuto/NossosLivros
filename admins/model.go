package admins

import (
	"crypto/sha256"
	"fmt"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model

	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func New(user, password string) *Admin {
	return &Admin{
		User:     user,
		Password: fmt.Sprintf("%x", sha256.Sum256([]byte(password))),
	}
}
