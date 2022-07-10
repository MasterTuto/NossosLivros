package database

import (
	"brenocs.dev/nossolivros/admins"
	"brenocs.dev/nossolivros/books"
	"brenocs.dev/nossolivros/loans"
	"brenocs.dev/nossolivros/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "root:admin@tcp(mysql:3306)/nossolivros?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&users.User{})
	db.AutoMigrate(&admins.Admin{})

	db.AutoMigrate(&books.Book{})
	db.AutoMigrate(&books.Stored{})
	db.AutoMigrate(&loans.Loan{})
}

func Get() *gorm.DB {
	return db
}
