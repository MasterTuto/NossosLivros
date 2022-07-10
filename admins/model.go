package admins

import (
	"crypto/sha256"
	"fmt"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model

	User     string `json:"user"`
	Password string `json:"password"`
}

func New(user, password string) *Admin {
	return &Admin{
		User:     user,
		Password: fmt.Sprintf("%x", sha256.Sum256([]byte(password))),
	}
}
