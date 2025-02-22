package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gauravlad21/book-management-system/commonutility"
	"github.com/gauravlad21/book-management-system/errors"
	"github.com/gauravlad21/book-management-system/external_resources/kafka"
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
	s.Cache.DeleteKey(ctx, commonutility.GetAllBooksKeyPrefix())

	kafka.PublishEvent("create", fmt.Sprint(book.ID), book.Title, book.Author, book.Year)
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

	bookJson, _ := json.Marshal(book)
	s.Cache.Set(ctx, &epredis.RedisKeyValue{Key: commonutility.GetCacheKey(id), Value: bookJson, ExpiryInMillis: 60 * 1000})

	return book, nil
}

func (s *ServiceStruct) ReadAllBooks(ctx context.Context, limit, offset int) []models.Book {

	cacheKey := commonutility.GetAllBooksKey(limit, offset)
	if cachedBooks, err := s.Cache.Get(ctx, cacheKey); err == nil {
		var books []models.Book
		json.Unmarshal([]byte(cachedBooks.(string)), &books)
		return books
	}

	books := s.DbOps.ReadAllBooks(ctx, limit, offset)

	booksJson, _ := json.Marshal(books)
	s.Cache.Set(ctx, &epredis.RedisKeyValue{Key: cacheKey, Value: booksJson, ExpiryInMillis: 60 * 1000})
	return books
}

func (s *ServiceStruct) UpdateBook(ctx context.Context, id string, book *models.Book) error {

	err := validateUpdateBook(id, book)
	if err != nil {
		return err
	}

	err = s.DbOps.UpdateBook(ctx, id, book)
	if err != nil {
		return err
	}

	s.Cache.DeleteKey(ctx, commonutility.GetAllBooksKeyPrefix(), commonutility.GetCacheKey(id))

	kafka.PublishEvent("update", fmt.Sprint(book.ID), book.Title, book.Author, book.Year)
	return nil
}

func (s *ServiceStruct) DeleteBook(ctx context.Context, id string) error {
	err := validateDeleteBook(id)
	if err != nil {
		return err
	}

	err = s.DbOps.DeleteBook(ctx, id)
	if err != nil {
		return err
	}

	s.Cache.DeleteKey(ctx, commonutility.GetAllBooksKeyPrefix(), commonutility.GetCacheKey(id))

	kafka.PublishEvent("delete", fmt.Sprint(id), "", "", 0)
	return nil
}
