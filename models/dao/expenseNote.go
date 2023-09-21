package dao

import (
	"gorm.io/gorm"
)

type ExpenseNoteDAO struct {
	gorm.Model
	Name   string `json:"name"`
	UserId int    `json:"user_id"`
}
