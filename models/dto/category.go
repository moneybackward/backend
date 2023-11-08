package dto

import (
	"github.com/google/uuid"
	"github.com/moneybackward/backend/models"
)

type CategoryCreateDTO struct {
	Name     string  `json:"name"`
	Priority int     `json:"priority"`
	Budget   float64 `json:"budget"`
}

type CategoryDTO struct {
	BaseDTO
	CategoryCreateDTO
	NoteId uuid.UUID `json:"note_id"`
}

func (dto *CategoryCreateDTO) ToEntity() models.Category {
	return models.Category{
		Name:     dto.Name,
		Priority: dto.Priority,
	}
}

func (dto *CategoryDTO) FromEntity(category models.Category) {
	dto.BaseDTO.FromEntity(&category.Base)
	dto.Name = category.Name
	dto.Priority = category.Priority
	dto.NoteId = category.NoteId
}

type CategorySetBudgetDTO struct {
	Budget float64 `json:"budget"`
}
