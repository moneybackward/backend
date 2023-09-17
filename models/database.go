package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to migrate database")
	}

	DB = db
}
