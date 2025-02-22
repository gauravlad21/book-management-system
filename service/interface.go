package service

import (
	"context"

	"github.com/gauravlad21/book-management-system/dbhelper"
	"github.com/gauravlad21/book-management-system/models"
)

type ServiceIF interface {
	Hello(ctx context.Context) string
	CreateBook(ctx context.Context, book *models.Book) error
	ReadBook(ctx context.Context, id string) (*models.Book, error)
	ReadAllBooks(ctx context.Context) []models.Book
}

type ServiceStruct struct {
	DbOps dbhelper.DbOperationsIF
}

func New(dbOps dbhelper.DbOperationsIF) ServiceIF {

	return &ServiceStruct{
		DbOps: dbOps,
	}
}
