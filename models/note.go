package models

type Note struct {
	Base
	Name   string `json:"name"`
	UserId int    `json:"user_id"`
	User   User   `gorm:"foreignKey:UserId"`
}
