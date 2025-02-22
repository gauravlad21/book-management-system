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
