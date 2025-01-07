package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	ConnectToDB()
	defer DB.Close(context.Background())

	r := gin.Default()

	r.POST("/books", CreateBook)
	r.GET("/books/:id", GetBook)
	r.GET("/books", GetAllBooks)
	r.PUT("/books", UpdateBook)
	r.DELETE("/books/:id", DeleteBook)

	log.Println("Starting server on :8080...")
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Error starting the server:", err)
		os.Exit(1)
	}
}
