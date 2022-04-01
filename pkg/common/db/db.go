package db

import (
	"log"

	"github.com/wbianchini/go-gin-api/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}
