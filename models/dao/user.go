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

type UserRepository interface {
	Save(user *UserDAO) (*UserDAO, error)
	FindAll() ([]UserDAO, error)
	Delete(user *UserDAO) error
	Migrate() error
}
