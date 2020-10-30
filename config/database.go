package config

import (
	"github.com/lrs-studies/poc-go-firebase-auth/api"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateDatabase() *gorm.DB {
	// Create db instance with gorm
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	// migrate our model for artist
	db.AutoMigrate(&api.Artist{})
	return db
}
