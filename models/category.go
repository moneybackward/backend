package models

type Category struct {
	Base
	Name     string `json:"name"`
	Priority int    `json:"priority"`
	NoteId   int    `json:"note_id"`
	Note     Note   `gorm:"foreignKey:NoteId"`
}
