package models

import "time"

type Transaction struct {
	Base
	Label     string    `json:"label"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`

	NoteId     int `json:"note_id"`
	CategoryId int `json:"category_id"`

	Note     Note     `gorm:"foreignKey:NoteId"`
	Category Category `gorm:"foreignKey:CategoryId"`
}
