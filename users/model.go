package users

import (
	"crypto/sha256"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	name, email, password string
	finishedLoans         int
	loanLimit             int
}

func New(name, email, password string) *User {
	return &User{
		name:      name,
		email:     email,
		loanLimit: 1,
		password:  fmt.Sprintf("%x", sha256.Sum256([]byte(password))),
	}
}

func (u User) Name() string {
	return u.name
}

func (u User) Email() string {
	return u.email
}

func (u User) VerifyPassword(password string) bool {
	return u.password == fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
}
