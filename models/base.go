package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	Id        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time      `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP; constraint:OnUpdate:CURRENT_TIMESTAMP;"`
	DeletedAt gorm.DeletedAt `swaggerignore:"true" json:"deleted_at" gorm:"index"`
}

func (base *Base) BeforeCreate(db *gorm.DB) error {
	db.Statement.SetColumn("Id", uuid.New().String())
	return nil
}
