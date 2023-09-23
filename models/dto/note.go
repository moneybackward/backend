package dto

import (
	"github.com/google/uuid"
	"github.com/moneybackward/backend/models"
)

type NoteDTO struct {
	Name   string    `json:"name" binding:"required"`
	UserId uuid.UUID `json:"user_id" binding:"required"`
}

func (dto *NoteDTO) ToEntity() (*models.Note, error) {
	u := &models.Note{
		Name:   dto.Name,
		UserId: dto.UserId,
	}

	return u, nil
}
