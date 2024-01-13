package dto

import (
	"github.com/google/uuid"
	"github.com/guregu/null"
	"github.com/moneybackward/backend/models"
)

type CategoryCreateDTO struct {
	Name      string   `json:"name"`
	Priority  int      `json:"priority"`
	Budget    *float64 `json:"budget,omitempty"`
	IsExpense *bool    `json:"is_expense"`
}

type CategoryUpdateDTO CategoryCreateDTO

type CategoryDTO struct {
	BaseDTO
	CategoryCreateDTO
	NoteId uuid.UUID `json:"note_id"`
}

type CategoryStatsDTO struct {
	CategoryDTO
	Total float64 `json:"total"`
	Count int     `json:"count"`
}

func (dto *CategoryCreateDTO) ToEntity() models.Category {
	nullableBudget := null.NewFloat(0, false)
	if dto.Budget != nil {
		nullableBudget = null.FloatFromPtr(dto.Budget)
	}
	return models.Category{
		Name:      dto.Name,
		Budget:    nullableBudget,
		Priority:  dto.Priority,
		IsExpense: dto.IsExpense,
	}
}

func (dto *CategoryUpdateDTO) ToEntity() models.Category {
	nullableBudget := null.NewFloat(0, false)
	if dto.Budget != nil {
		nullableBudget = null.FloatFromPtr(dto.Budget)
	}

	return models.Category{
		Name:      dto.Name,
		Budget:    nullableBudget,
		Priority:  dto.Priority,
		IsExpense: dto.IsExpense,
	}
}

func (dto *CategoryDTO) ToEntity() models.Category {
	nullableBudget := null.NewFloat(0, false)
	if dto.Budget != nil {
		nullableBudget = null.FloatFromPtr(dto.Budget)
	}

	return models.Category{
		NoteId:    dto.NoteId,
		Name:      dto.Name,
		Priority:  dto.Priority,
		IsExpense: dto.IsExpense,
		Budget:    nullableBudget,
	}
}

func (dto *CategoryDTO) FromEntity(category models.Category) {
	var budget *float64
	if category.Budget.Valid {
		budget = &category.Budget.Float64
	}

	dto.BaseDTO.FromEntity(&category.Base)
	dto.Name = category.Name
	dto.Priority = category.Priority
	dto.Budget = budget
	dto.NoteId = category.NoteId
	dto.IsExpense = category.IsExpense
}

func (dto *CategoryDTO) FromCreateDto(createDto CategoryCreateDTO) {
	dto.Name = createDto.Name
	dto.Priority = createDto.Priority
	dto.Budget = createDto.Budget
	dto.IsExpense = createDto.IsExpense
}
