package dao

import (
	"gorm.io/gorm"
)

type UserDAO struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
