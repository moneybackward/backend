package models

import (
	"github.com/moneybackward/backend/models/dao"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&dao.UserDAO{})
	if err != nil {
		panic("failed to migrate database")
	}

	DB = db
	return DB
}
