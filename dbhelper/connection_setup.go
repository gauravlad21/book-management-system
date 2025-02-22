package dbhelper

import (
	"github.com/gauravlad21/book-management-system/models"
	"gorm.io/gorm"
)

var dbOps DbOperationsIF

func GetDbOps() DbOperationsIF {
	if dbOps == nil {
		dbOps = New()
	}
	return dbOps
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.Book{})
}
