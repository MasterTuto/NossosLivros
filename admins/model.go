package admins

import (
	"crypto/sha256"
	"fmt"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model

	user     string
	password string
}

func New(user, password string) *Admin {
	return &Admin{
		user:     user,
		password: fmt.Sprintf("%x", sha256.Sum256([]byte(password))),
	}
}
