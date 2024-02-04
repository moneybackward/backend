package models

import (
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Base
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	LastLogin *time.Time `json:"last_login"`
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	// trim spaces
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return nil
}
