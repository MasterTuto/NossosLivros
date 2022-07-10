package users

import (
	"crypto/sha256"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	FinishedLoans int    `json:"finishedLoans"`
	LoanLimit     int    `json:"loanLimit"`
}

func New(name, email, password string) *User {
	return &User{
		Name:      name,
		Email:     email,
		LoanLimit: 1,
		Password:  fmt.Sprintf("%x", sha256.Sum256([]byte(password))),
	}
}

func (u User) VerifyPassword(password string) bool {
	return u.Password == fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
}
