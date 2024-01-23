package dto

import (
	"github.com/google/uuid"
	"github.com/moneybackward/backend/models"
)

type NoteDTO struct {
	BaseDTO
	Name         string    `json:"name" binding:"required"`
	CurrencyCode string    `json:"currency_code" binding:"required"`
	UserId       uuid.UUID `json:"user_id" binding:"required"`
}

type NoteCreateDTO struct {
	Name         string `json:"name" binding:"required"`
	CurrencyCode string `json:"currency_code" binding:"required"`
}

type NoteUpdateDTO NoteCreateDTO

func (dto *NoteDTO) ToEntity() (*models.Note, error) {
	u := &models.Note{
		Name:         dto.Name,
		UserId:       dto.UserId,
		CurrencyCode: dto.CurrencyCode,
	}

	return u, nil
}

func (dto *NoteCreateDTO) ToEntity() (*models.Note, error) {
	u := &models.Note{
		Name:         dto.Name,
		CurrencyCode: dto.CurrencyCode,
	}

	return u, nil
}

func (dto *NoteDTO) FromEntity(note *models.Note) {
	dto.BaseDTO.FromEntity(&note.Base)
	dto.Name = note.Name
	dto.UserId = note.UserId
	dto.CurrencyCode = note.CurrencyCode
}
