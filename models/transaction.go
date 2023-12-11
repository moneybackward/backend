package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Base
	Label     string    `json:"label"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
	IsExpense bool      `json:"is_expense" gorm:"default:true"`

	NoteId     uuid.UUID `json:"note_id"`
	CategoryId uuid.UUID `json:"category_id"`

	Note     Note     `gorm:"foreignKey:NoteId"`
	Category Category `gorm:"foreignKey:CategoryId"`
}
