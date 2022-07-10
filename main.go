package main

import (
	"brenocs.dev/nossolivros/admins"
	"brenocs.dev/nossolivros/books"
	"brenocs.dev/nossolivros/database"
	"brenocs.dev/nossolivros/loans"
	"brenocs.dev/nossolivros/users"
	"github.com/gin-gonic/gin"
)

func setDB(ctx *gin.Context) {
	ctx.Set("db", database.Get())
}

func main() {
	r := gin.Default()

	r.Use(setDB)

	users.UserRoutes(r)
	loans.LoansRouter(r)
	admins.AdminRoutes(r)
	books.BooksRouter(r)
}
