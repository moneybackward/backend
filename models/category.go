package models

import "github.com/google/uuid"

type Category struct {
	Base
	Name      string    `json:"name"`
	Priority  int       `json:"priority"`
	Budget    float64   `json:"budget"`
	NoteId    uuid.UUID `json:"note_id"`
	IsExpense bool      `json:"is_expense" gorm:"default:true; not null"`

	Note Note `gorm:"foreignKey:NoteId"`
}
