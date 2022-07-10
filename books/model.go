package books

import (
	"time"

	"brenocs.dev/nossolivros/users"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Release     time.Time `json:"release"`
	Description string    `json:"description"`
	Cover       string    `json:"cover"`
}

type Stored struct {
	gorm.Model
	User          users.User `json:"user"`
	Book          Book       `json:"book"`
	StoredAt      time.Time  `json:"storedAt"`
	TimesLoaned   int        `json:"timesLoaned"`
	IsBeingLoaned bool       `json:"isBeingLoaned"`
}

func New(title, author, description, cover string, release time.Time) *Book {
	return &Book{
		Title:       title,
		Author:      author,
		Description: description,
		Cover:       cover,
		Release:     release,
	}
}

func NewStored(book Book, user users.User) *Stored {
	return &Stored{
		User:          user,
		Book:          book,
		TimesLoaned:   0,
		IsBeingLoaned: false,
		StoredAt:      time.Now(),
	}
}

func genBooks() []Book {
	books := make([]Book, 7)
	books[0] = Book{
		Title:       "In Search of Lost Time",
		Author:      "Marcel Proust",
		Description: "In Search of Lost Time is a novel by Marcel Proust, published in 1913. It is one of his best-known works, and is considered one of the most important works of literature.",
		Cover:       "https://images-na.ssl-images-amazon.com/images/I/51-X-Q-X-QL._SX329_BO1,204,203,200_.jpg",
		Release:     time.Date(1999, time.May, 19, 0, 0, 0, 0, time.UTC),
	}

	books[1] = Book{
		Title:       "O Corte Ingles",
		Author:      "Machado de Assis",
		Description: "O Corte Ingles is a novel by Brazilian writer Machado de Assis, published in 1891. It is one of his best-known works, and is considered one of the most important works of literature.",
		Cover:       "https://images-na.ssl-images-amazon.com/images/I/51-X-Q-X-QL._SX329_BO1,204,203,200_.jpg",
		Release:     time.Date(1891, time.May, 19, 0, 0, 0, 0, time.UTC),
	}

	books[2] = Book{
		Title:       "The Trial",
		Author:      "Franz Kafka",
		Description: "The Trial is a novel by Franz Kafka, published in 1925. It is one of his best-known works, and is considered one of the most important works of literature.",
		Cover:       "https://images-na.ssl-images-amazon.com/images/I/51-X-Q-X-QL._SX329_BO1,204,203,200_.jpg",
		Release:     time.Date(1925, time.May, 19, 0, 0, 0, 0, time.UTC),
	}

	books[3] = Book{
		Title:       "The Stranger",
		Author:      "Albert Camus",
		Description: "The Stranger is a novel by Albert Camus, published in 1913. It is one of his best-known works, and is considered one of the most important works of literature.",
		Cover:       "https://images-na.ssl-images-amazon.com/images/I/51-X-Q-X-QL._SX329_BO1,204,203,200_.jpg",
		Release:     time.Date(1913, time.May, 19, 0, 0, 0, 0, time.UTC),
	}

	books[4] = Book{
		Title:       "The Birth of the Freethought",
		Author:      "Friederich Nietzsche",
		Description: "The Birth of the Freethought is a novel by Friedrich Nietzsche, published in 1844. It is one of his best-known works, and is considered one of the most important works of literature.",
		Cover:       "https://images-na.ssl-images-amazon.com/images/I/51-X-Q-X-QL._SX329_BO1,204,203,200_.jpg",
		Release:     time.Date(1844, time.May, 19, 0, 0, 0, 0, time.UTC),
	}

	books[5] = Book{
		Title:       "The Complete Works of Jean-Paul Sartre",
		Author:      "Jean-Paul Sartre",
		Description: "The Complete Works of Jean-Paul Sartre is a collection of writings by French philosopher Jean-Paul Sartre, published in 1913. It is one of his best-known works, and is considered one of the most important works of literature.",
		Cover:       "https://images-na.ssl-images-amazon.com/images/I/51-X-Q-X-QL._SX329_BO1,204,203,200_.jpg",
		Release:     time.Date(1913, time.May, 19, 0, 0, 0, 0, time.UTC),
	}

	books[6] = Book{
		Author:      "Fiódor Dostoievski",
		Title:       "The Idiot",
		Description: "The Idiot is a novel by Fiodor Dostoievski, published in 1869. It is one of his best-known works, and is considered one of the most important works of literature.",
		Cover:       "https://images-na.ssl-images-amazon.com/images/I/51-X-Q-X-QL._SX329_BO1,204,203,200_.jpg",
		Release:     time.Date(1869, time.May, 19, 0, 0, 0, 0, time.UTC),
	}

	books[7] = Book{
		Author:      "Ernest Hemingway",
		Title:       "For Whom the Bell Tolls",
		Description: "For Whom the Bell Tolls is a novel by Ernest Hemingway, published in 1940. It is one of his best-known works, and is considered one of the most important works of literature.",
		Cover:       "https://images-na.ssl-images-amazon.com/images/I/51-X-Q-X-QL._SX329_BO1,204,203,200_.jpg",
		Release:     time.Date(1940, time.May, 19, 0, 0, 0, 0, time.UTC),
	}

	return books
}
