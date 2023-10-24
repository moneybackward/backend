package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/moneybackward/backend/models"
)

type BaseDTO struct {
	Id        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (dto *BaseDTO) FromEntity(base models.Base) {
	dto.Id = base.Id
	dto.CreatedAt = base.CreatedAt
	dto.UpdatedAt = base.UpdatedAt
	dto.DeletedAt = nil
	if base.DeletedAt.Valid {
		dto.DeletedAt = &base.DeletedAt.Time
	}
}
