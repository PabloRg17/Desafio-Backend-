package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Author   string `json:"author"`
	Synopsis string `json:"synopsis"`
}

func CreateBook(c *gin.Context) {
	var book Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	query := `INSERT INTO books (title, category, author, synopsis) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err := DB.QueryRow(context.Background(), query, book.Title, book.Category, book.Author, book.Synopsis).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	book.ID = id
	c.JSON(http.StatusCreated, book)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book Book

	query := `SELECT id, title, category, author, synopsis FROM books WHERE id = $1`
	err := DB.QueryRow(context.Background(), query, id).Scan(&book.ID, &book.Title, &book.Category, &book.Author, &book.Synopsis)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func GetAllBooks(c *gin.Context) {
	rows, err := DB.Query(context.Background(), `SELECT id, title, category, author, synopsis FROM books`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Category, &book.Author, &book.Synopsis)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read rows"})
			return
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

func UpdateBook(c *gin.Context) {
	var book Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	query := `UPDATE books SET title = $1, category = $2, author = $3, synopsis = $4 WHERE id = $5`
	_, err := DB.Exec(context.Background(), query, book.Title, book.Category, book.Author, book.Synopsis, book.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	_, err := DB.Exec(context.Background(), `DELETE FROM books WHERE id = $1`, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
