package dto

import "github.com/moneybackward/backend/models"

type CategoryDTO struct {
	Name     string `json:"name"`
	Priority int    `json:"priority"`
	NoteId   int    `json:"note_id"`
}

func (dto *CategoryDTO) ToEntity() (*models.Category, error) {
	u := &models.Category{
		Name:     dto.Name,
		Priority: dto.Priority,
		NoteId:   dto.NoteId,
	}

	return u, nil
}
