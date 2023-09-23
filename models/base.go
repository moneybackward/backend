package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	Id        uuid.UUID    `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

func (base *Base) BeforeCreate(db *gorm.DB) error {
	db.Statement.SetColumn("Id", uuid.New().String())
	return nil
}
