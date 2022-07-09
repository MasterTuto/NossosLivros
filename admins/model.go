package admins

import "gorm.io/gorm"

type Admin struct {
	gorm.Model

	user     string
	password string
}

func NewAdmin(user, password string) *Admin {
	return &Admin{
		user:     user,
		password: password,
	}
}
