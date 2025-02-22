package service

import (
	"context"
	"encoding/json"

	"github.com/gauravlad21/book-management-system/commonutility"
	"github.com/gauravlad21/book-management-system/errors"
	epredis "github.com/gauravlad21/book-management-system/external_resources/redis"
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

	if val, err := s.Cache.Get(ctx, commonutility.GetCacheKey(id)); err == nil {
		book := &models.Book{}
		json.Unmarshal([]byte(val.(string)), book)
		return book, nil
	}

	book, err := s.DbOps.ReadBook(ctx, id)
	if err != nil {
		return nil, errors.ErrInternal
	}

	booksJson, _ := json.Marshal(book)
	s.Cache.Set(ctx, &epredis.RedisKeyValue{Key: commonutility.GetCacheKey(id), Value: booksJson})

	return book, nil
}

func (s *ServiceStruct) ReadAllBooks(ctx context.Context) []models.Book {
	book := s.DbOps.ReadAllBooks(ctx)
	return book
}
