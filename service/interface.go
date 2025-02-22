package service

import (
	"context"

	"github.com/gauravlad21/book-management-system/dbhelper"
	epredis "github.com/gauravlad21/book-management-system/external_resources/redis"
	"github.com/gauravlad21/book-management-system/models"
)

type ServiceIF interface {
	Hello(ctx context.Context) string
	CreateBook(ctx context.Context, book *models.Book) error
	ReadBook(ctx context.Context, id string) (*models.Book, error)
	ReadAllBooks(ctx context.Context, limit, offset int) []models.Book
	DeleteBook(ctx context.Context, id string) error
	UpdateBook(ctx context.Context, id string, book *models.Book) error
}

type ServiceStruct struct {
	DbOps dbhelper.DbOperationsIF
	Cache epredis.CacheInterface
}

func New(dbOps dbhelper.DbOperationsIF, cache epredis.CacheInterface) ServiceIF {

	return &ServiceStruct{
		DbOps: dbOps,
		Cache: cache,
	}
}
