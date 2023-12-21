package db

import (
	"cc-auth-service/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// initializes the database
func InitDatabase() {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	db.AutoMigrate(&model.User{})
}
