package loans

import (
	"time"

	"brenocs.dev/nossolivros/books"
	"brenocs.dev/nossolivros/users"
	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model
	storedBook            books.Stored
	loaner                users.User
	loanTime              time.Time
	returnDate            time.Time
	bookRate              float64
	bookRateDescription   string
	loanerRate            float64
	loanerRateDescription string
}

func New(storedBook books.Stored, loaner users.User) *Loan {
	return &Loan{
		storedBook:            storedBook,
		loaner:                loaner,
		loanTime:              time.Now(),
		returnDate:            time.Time{},
		bookRate:              0,
		bookRateDescription:   "",
		loanerRate:            0,
		loanerRateDescription: "",
	}
}
