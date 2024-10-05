package main

import (
	"studio-app/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/classes", handlers.CreateClass)
	r.GET("/classes", handlers.ListOfClasses)
	r.POST("/bookings", handlers.BookClass)

	r.Run(":8080")
}
