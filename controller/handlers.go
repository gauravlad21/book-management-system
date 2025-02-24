package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gauravlad21/book-management-system/commonutility"
	"github.com/gauravlad21/book-management-system/external_resources/kafka"
	"github.com/gauravlad21/book-management-system/models"

	"github.com/gin-gonic/gin"
)

func Hello(ctx *gin.Context) {
	commonutility.GetLogger().Info("Hello from API successful")
	msg := serviceRepo.Hello(commonutility.GetContext(ctx))
	ctx.JSON(200, msg)
}

// Create a Book
// @Summary Create a new book
// @Description Add a new book to the database
// @Tags Books
// @Accept  json
// @Produce  json
// @Param book body models.Book true "Book object"
// @Success 201 {object} models.Book
// @Failure 400 {object} map[string]string
// @Router /books [post]
// @host http://13.51.170.227:5002
func CreateBook(ctx *gin.Context) {
	book := &models.Book{}
	ctx.BindJSON(&book)
	err := serviceRepo.CreateBook(commonutility.GetContext(ctx), book)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, book)
}

// Get Book by ID
// @Summary Get book by ID
// @Description Retrieve a book by its ID
// @Tags Books
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {object} map[string]string
// @Router /books/{id} [get]
// @host http://13.51.170.227:5002
func ReadBook(ctx *gin.Context) {
	id := ctx.Param("id")
	book, err := serviceRepo.ReadBook(commonutility.GetContext(ctx), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	ctx.JSON(200, book)
}

// Get All Books
// @Summary Get all books
// @Description Retrieve a list of all books
// @Tags Books
// @Produce  json
// @Success 200 {array} models.Book
// @Router /books [get]
// @host http://13.51.170.227:5002
func ReadAllBooks(ctx *gin.Context) {
	limitStr := ctx.Request.URL.Query().Get("limit")
	offsetStr := ctx.Request.URL.Query().Get("offset")

	limit, err1 := strconv.Atoi(limitStr)
	offset, err2 := strconv.Atoi(offsetStr)
	if err1 != nil || err2 != nil {
		limit = 10
		offset = 0
	}

	book := serviceRepo.ReadAllBooks(commonutility.GetContext(ctx), limit, offset)
	ctx.JSON(200, book)
}

// Update a Book
// @Summary Update a book
// @Description Update an existing book's details
// @Tags Books
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Param book body models.Book true "Updated book object"
// @Success 200 {object} models.Book
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /books/{id} [put]
// @host http://13.51.170.227:5002
func UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	book := &models.Book{}
	ctx.BindJSON(&book)
	err := serviceRepo.UpdateBook(commonutility.GetContext(ctx), id, book)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, book)
}

// Delete a Book
// @Summary Delete a book
// @Description Delete a book by ID
// @Tags Books
// @Param id path int true "Book ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /books/{id} [delete]
// @host http://13.51.170.227:5002
func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	err := serviceRepo.DeleteBook(commonutility.GetContext(ctx), id)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"Message": fmt.Sprintf("Successfully Deleted id: %v", id)})
}

// Kafka Consumer
// @Summary Kafka Consumer
// @Description Kafka message received by Post, Put, Delete event
// @Tags Books
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /events [get]
// @host http://13.51.170.227:5002
func GetEvents(ctx *gin.Context) {
	messages := kafka.GetEvents(commonutility.GetContext(ctx))
	ctx.JSON(http.StatusOK, gin.H{"events": messages})
}
