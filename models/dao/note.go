package dao

import (
	"gorm.io/gorm"
)

type NoteDAO struct {
	gorm.Model
	Name   string `json:"name"`
	UserId int    `json:"user_id"`
}
