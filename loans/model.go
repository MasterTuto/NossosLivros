package loans

import (
	"time"

	"brenocs.dev/nossolivros/books"
	"brenocs.dev/nossolivros/users"
	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model
	StoredBook            books.Stored `json:"storedBook"`
	Loaner                users.User   `json:"loander"`
	LoanTime              time.Time    `json:"loanTime"`
	ReturnDate            time.Time    `json:"returnDate"`
	BookRate              float64      `json:"bookRate"`
	BookRateDescription   string       `json:"bookRateDescription"`
	LoanerRate            float64      `json:"loanerRate"`
	LoanerRateDescription string       `json:"loanerRateDescription"`
}

func New(storedBook books.Stored, loaner users.User) *Loan {
	return &Loan{
		StoredBook:            storedBook,
		Loaner:                loaner,
		LoanTime:              time.Now(),
		ReturnDate:            time.Time{},
		BookRate:              0,
		BookRateDescription:   "",
		LoanerRate:            0,
		LoanerRateDescription: "",
	}
}
