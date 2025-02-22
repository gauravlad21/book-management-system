package dbhelper

import (
	"context"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gauravlad21/book-management-system/errors"
	"github.com/gauravlad21/book-management-system/models"
	"github.com/spf13/viper"
)

type DbOperationsIF interface {
	CreateBook(ctx context.Context, book *models.Book) (int, error)
	ReadBook(ctx context.Context, id string) (*models.Book, error)
	ReadAllBooks(ctx context.Context) []models.Book
	// UpdateBook(ctx context.Context, book *models.Book)
	// DeleteBook(ctx context.Context, id int)
}

type DbOps struct {
	DB *gorm.DB
}

func New() DbOperationsIF {
	host := viper.Get("db.host")
	user := viper.Get("db.username")
	pass := viper.Get("db.password")
	dbname := viper.Get("db.dbname")
	port := viper.Get("db.port")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", host, user, pass, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return &DbOps{DB: db}
}

func (d *DbOps) CreateBook(ctx context.Context, book *models.Book) (int, error) {
	tx := d.DB.Create(&book)
	tx.Commit()
	return int(book.ID), nil
}

func (d *DbOps) ReadBook(ctx context.Context, id string) (*models.Book, error) {
	book := &models.Book{}
	d.DB.First(book, id)
	if fmt.Sprint(book.ID) != id {
		return nil, errors.ErrNotFound
	}
	return book, nil
}

func (d *DbOps) ReadAllBooks(ctx context.Context) []models.Book {
	var books []models.Book
	d.DB.Find(&books)
	return books
}

// func (d *DbOps) UpdateBook(ctx context.Context, id int, book *models.Book) {
// 	d.DB.Save(book)
// }

// func (d *DbOps) DeleteBook(ctx context.Context, id int) {

// }
