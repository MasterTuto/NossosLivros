package main

import (
	"brenocs.dev/nossolivros/users"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	users.UserRoutes(r)
}
