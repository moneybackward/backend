package models

import "github.com/google/uuid"

type Note struct {
	Base
	Name   string    `json:"name"`
	UserId uuid.UUID `json:"user_id"`
	User   User      `gorm:"foreignKey:UserId"`

	Categories   []Category    `gorm:"foreignKey:NoteId"`
	Transactions []Transaction `gorm:"foreignKey:NoteId"`
}
