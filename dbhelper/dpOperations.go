package dbhelper

import (
	"context"
	"log"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gauravlad21/book-management-system/errors"
	"github.com/gauravlad21/book-management-system/models"
)

type DbOperationsIF interface {
	CreateBook(ctx context.Context, book *models.Book) (int, error)
	ReadBook(ctx context.Context, id string) (*models.Book, error)
	ReadAllBooks(ctx context.Context, limit, offset int) []models.Book
	UpdateBook(ctx context.Context, id string, book *models.Book) error
	DeleteBook(ctx context.Context, id string) error
}

type DbOps struct {
	DB *gorm.DB
}

func New() DbOperationsIF {
	// host := viper.Get("db.host")
	// user := viper.Get("db.username")
	// pass := viper.Get("db.password")
	// dbname := viper.Get("db.dbname")
	// port := viper.Get("db.port")

	// dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", host, user, pass, dbname, port)
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	MigrateDB(db)
	return &DbOps{DB: db}
}

func (d *DbOps) CreateBook(ctx context.Context, book *models.Book) (int, error) {
	tx := d.DB.Create(&book)
	tx.Commit()
	return int(book.ID), nil
}

func (d *DbOps) ReadBook(ctx context.Context, id string) (*models.Book, error) {
	book := &models.Book{}
	tx := d.DB.First(book, "id = ?", id)
	if tx.Error != nil {
		return nil, errors.ErrNotFound
	}
	return book, nil
}

func (d *DbOps) ReadAllBooks(ctx context.Context, limit, offset int) []models.Book {
	var books []models.Book
	d.DB.Limit(limit).Offset(offset).Find(&books)
	return books
}

func (d *DbOps) UpdateBook(ctx context.Context, id string, book *models.Book) error {
	oldBook := &models.Book{}
	if err := d.DB.First(oldBook, "id = ? ", id).Error; err == nil {
		x, _ := strconv.Atoi(id)
		book.ID = uint(x)
		tx := d.DB.Save(book)
		if tx.Error != nil {
			return err
		}
		tx.Commit()
		return nil
	}
	return errors.ErrNotFound
}

func (d *DbOps) DeleteBook(ctx context.Context, id string) error {
	oldBook := &models.Book{}
	if err := d.DB.First(oldBook, id).Error; err == nil {
		tx := d.DB.Delete(oldBook)
		if tx.Error != nil {
			return err
		}
		tx.Commit()
		return nil
	}
	return errors.ErrNotFound
}
