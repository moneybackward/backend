package dto

import (
	"github.com/google/uuid"
	"github.com/moneybackward/backend/models"
)

type CategoryCreateDTO struct {
	Name      string  `json:"name"`
	Priority  int     `json:"priority"`
	Budget    float64 `json:"budget"`
	IsExpense bool    `json:"is_expense"`
}

type CategoryUpdateDTO CategoryCreateDTO

type CategoryDTO struct {
	BaseDTO
	CategoryCreateDTO
	NoteId uuid.UUID `json:"note_id"`
}

func (dto *CategoryCreateDTO) ToEntity() models.Category {
	return models.Category{
		Name:      dto.Name,
		Budget:    dto.Budget,
		Priority:  dto.Priority,
		IsExpense: dto.IsExpense,
	}
}

func (dto *CategoryUpdateDTO) ToEntity() models.Category {
	return models.Category{
		Name:      dto.Name,
		Budget:    dto.Budget,
		Priority:  dto.Priority,
		IsExpense: dto.IsExpense,
	}
}

func (dto *CategoryDTO) ToEntity() models.Category {
	return models.Category{
		NoteId:    dto.NoteId,
		Name:      dto.Name,
		Priority:  dto.Priority,
		IsExpense: dto.IsExpense,
	}
}

func (dto *CategoryDTO) FromEntity(category models.Category) {
	dto.BaseDTO.FromEntity(&category.Base)
	dto.Name = category.Name
	dto.Priority = category.Priority
	dto.Budget = category.Budget
	dto.NoteId = category.NoteId
	dto.IsExpense = category.IsExpense
}

func (dto *CategoryDTO) FromCreateDto(createDto CategoryCreateDTO) {
	dto.Name = createDto.Name
	dto.Priority = createDto.Priority
	dto.Budget = createDto.Budget
	dto.IsExpense = createDto.IsExpense
}
