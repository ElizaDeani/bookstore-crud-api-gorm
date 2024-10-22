package config

import (
	"post-test-mikti/pkg/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrasi skema database
	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
