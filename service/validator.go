package service

import (
	"time"

	"github.com/gauravlad21/book-management-system/errors"
	"github.com/gauravlad21/book-management-system/models"
)

func validateCreateBook(book *models.Book) error {
	if book == nil || book.Author == "" || book.Title == "" || book.Year < 1900 || book.Year > time.Now().Year() {
		return errors.ErrBadRequest
	}
	return nil
}
func validateUpdateBook(id string, book *models.Book) error {
	if book == nil || book.Author == "" || book.Title == "" || book.Year < 1900 || book.Year > time.Now().Year() || id == "0" || id == "" {
		return errors.ErrBadRequest
	}
	return nil
}

func validateDeleteBook(id string) error {
	if id == "" {
		return errors.ErrBadRequest
	}
	return nil
}
