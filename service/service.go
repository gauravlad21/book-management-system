package service

import (
	"context"

	"github.com/gauravlad21/book-management-system/commonutility"
	"github.com/gauravlad21/book-management-system/errors"
	"github.com/gauravlad21/book-management-system/models"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func (s *ServiceStruct) Hello(ctx context.Context) string {
	commonutility.GetLogger().Info("Hello from service", zap.Any("someKey", "someValue"))
	return "hello from |" + viper.GetString("value") + "| charting-engine service"
}

func (s *ServiceStruct) CreateBook(ctx context.Context, book *models.Book) error {
	err := validateCreateBook(book)
	if err != nil {
		return err
	}

	id, err := s.DbOps.CreateBook(ctx, book)
	if err != nil {
		return errors.ErrInternal
	}
	if id == 0 {
		return errors.ErrNotCreated
	}
	return nil
}

func (s *ServiceStruct) ReadBook(ctx context.Context, id string) (*models.Book, error) {
	book, err := s.DbOps.ReadBook(ctx, id)
	if err != nil {
		return nil, errors.ErrInternal
	}
	return book, nil
}

func (s *ServiceStruct) ReadAllBooks(ctx context.Context) []models.Book {
	book := s.DbOps.ReadAllBooks(ctx)
	return book
}
