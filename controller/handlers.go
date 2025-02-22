package controller

import (
	"fmt"
	"net/http"

	"github.com/gauravlad21/book-management-system/commonutility"
	"github.com/gauravlad21/book-management-system/models"

	"github.com/gin-gonic/gin"
)

func Hello(ctx *gin.Context) {
	commonutility.GetLogger().Info("Hello from API successful")
	msg := serviceRepo.Hello(commonutility.GetContext(ctx))
	ctx.JSON(200, msg)
}

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

func ReadBook(ctx *gin.Context) {
	id := ctx.Param("id")
	book, err := serviceRepo.ReadBook(commonutility.GetContext(ctx), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	ctx.JSON(200, book)
}

func ReadAllBooks(ctx *gin.Context) {
	book := serviceRepo.ReadAllBooks(commonutility.GetContext(ctx))
	ctx.JSON(200, book)
}

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

func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	err := serviceRepo.DeleteBook(commonutility.GetContext(ctx), id)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"Message": fmt.Sprintf("Successfully Deleted id: %v", id)})
}
