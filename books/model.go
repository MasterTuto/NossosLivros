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

func New(title, author, description, cover string, release time.Time) *Book {
	return &Book{
		title:       title,
		author:      author,
		description: description,
		cover:       cover,
		release:     release,
	}
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

func genBooks() []Book {
	books := make([]Book, 7)
	books[0] = Book{
		title:       "In Search of Lost Time",
		author:      "Marcel Proust",
		description: "In Search of Lost Time is a novel by Marcel Proust, published in 1913. It is one of his best-known works, and is considered one of the most important works of literature.",
		cover:       "https://images-na.ssl-images-amazon.com/images/I/51-X-Q-X-QL._SX329_BO1,204,203,200_.jpg",
		release:     time.Date(1999, time.May, 19, 0, 0, 0, 0, time.UTC),
	}

	books[1] = Book{
		title:       "O Corte Ingles",
		author:      "Machado de Assis",
		description: "O Corte Ingles is a novel by Brazilian writer Machado de Assis, published in 1891. It is one of his best-known works, and is considered one of the most important works of literature.",
		cover:       "https://images-na.ssl-images-amazon.com/images/I/51-X-Q-X-QL._SX329_BO1,204,203,200_.jpg",
		release:     time.Date(1891, time.May, 19, 0, 0, 0, 0, time.UTC),
	}

	books[2] = Book{
		title:       "The Trial",
		author:      "Franz Kafka",
		description: "The Trial is a novel by Franz Kafka, published in 1925. It is one of his best-known works, and is considered one of the most important works of literature.",
		cover:       "https://images-na.ssl-images-amazon.com/images/I/51-X-Q-X-QL._SX329_BO1,204,203,200_.jpg",
		release:     time.Date(1925, time.May, 19, 0, 0, 0, 0, time.UTC),
	}

	books[3] = Book{
		title:       "The Stranger",
		author:      "Albert Camus",
		description: "The Stranger is a novel by Albert Camus, published in 1913. It is one of his best-known works, and is considered one of the most important works of literature.",
		cover:       "https://images-na.ssl-images-amazon.com/images/I/51-X-Q-X-QL._SX329_BO1,204,203,200_.jpg",
		release:     time.Date(1913, time.May, 19, 0, 0, 0, 0, time.UTC),
	}

	books[4] = Book{
		title:       "The Birth of the Freethought",
		author:      "Friederich Nietzsche",
		description: "The Birth of the Freethought is a novel by Friedrich Nietzsche, published in 1844. It is one of his best-known works, and is considered one of the most important works of literature.",
		cover:       "https://images-na.ssl-images-amazon.com/images/I/51-X-Q-X-QL._SX329_BO1,204,203,200_.jpg",
		release:     time.Date(1844, time.May, 19, 0, 0, 0, 0, time.UTC),
	}

	books[5] = Book{
		title:       "The Complete Works of Jean-Paul Sartre",
		author:      "Jean-Paul Sartre",
		description: "The Complete Works of Jean-Paul Sartre is a collection of writings by French philosopher Jean-Paul Sartre, published in 1913. It is one of his best-known works, and is considered one of the most important works of literature.",
		cover:       "https://images-na.ssl-images-amazon.com/images/I/51-X-Q-X-QL._SX329_BO1,204,203,200_.jpg",
		release:     time.Date(1913, time.May, 19, 0, 0, 0, 0, time.UTC),
	}

	books[6] = Book{
		author:      "Fiódor Dostoievski",
		title:       "The Idiot",
		description: "The Idiot is a novel by Fiodor Dostoievski, published in 1869. It is one of his best-known works, and is considered one of the most important works of literature.",
		cover:       "https://images-na.ssl-images-amazon.com/images/I/51-X-Q-X-QL._SX329_BO1,204,203,200_.jpg",
		release:     time.Date(1869, time.May, 19, 0, 0, 0, 0, time.UTC),
	}

	books[7] = Book{
		author:      "Ernest Hemingway",
		title:       "For Whom the Bell Tolls",
		description: "For Whom the Bell Tolls is a novel by Ernest Hemingway, published in 1940. It is one of his best-known works, and is considered one of the most important works of literature.",
		cover:       "https://images-na.ssl-images-amazon.com/images/I/51-X-Q-X-QL._SX329_BO1,204,203,200_.jpg",
		release:     time.Date(1940, time.May, 19, 0, 0, 0, 0, time.UTC),
	}

	return books
}
