package models

import (
	"github.com/google/uuid"
	"github.com/guregu/null"
)

type Category struct {
	Base
	Name      string     `json:"name"`
	Priority  int        `json:"priority"`
	Budget    null.Float `json:"budget" gorm:"default:null"`
	NoteId    uuid.UUID  `json:"note_id"`
	IsExpense *bool      `json:"is_expense" gorm:"default:true; not null"`

	Note         Note          `gorm:"foreignKey:NoteId"`
	Transactions []Transaction `gorm:"foreignKey:CategoryId"`
}
