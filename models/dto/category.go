package dto

import "github.com/moneybackward/backend/models"

type CategoryCreateDTO struct {
	Name     string `json:"name"`
	Priority int    `json:"priority"`
	NoteId   int    `json:"note_id"`
}

type CategoryDTO struct {
	BaseDTO
	CategoryCreateDTO
}

func (dto *CategoryCreateDTO) ToEntity() models.Category {
	return models.Category{
		Name:     dto.Name,
		Priority: dto.Priority,
		NoteId:   dto.NoteId,
	}
}

func (dto *CategoryDTO) FromEntity(category models.Category) {
	dto.BaseDTO.FromEntity(&category.Base)
	dto.Name = category.Name
	dto.Priority = category.Priority
	dto.NoteId = category.NoteId
}
