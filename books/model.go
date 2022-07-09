package books

import (
	"time"

	"brenocs.dev/nossolivros/users"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	title       string
	author      string
	release     time.Time
	description string
	cover       string
}

type Stored struct {
	gorm.Model
	user          users.User
	book          Book
	storedAt      time.Time
	timesLoaned   int
	isBeingLoaned bool
}

func NewStored(book Book, user users.User) *Stored {
	return &Stored{
		user:          user,
		book:          book,
		timesLoaned:   0,
		isBeingLoaned: false,
		storedAt:      time.Now(),
	}
}
