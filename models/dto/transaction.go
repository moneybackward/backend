package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/moneybackward/backend/models"
)

type TransactionDTO struct {
	BaseDTO
	Label     string    `json:"label"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`

	NoteId     uuid.UUID `json:"note_id"`
	CategoryId uuid.UUID `json:"category_id"`
}

type TransactionCreateDTO struct {
	Label      string    `json:"label"`
	Amount     float64   `json:"amount"`
	CategoryId uuid.UUID `json:"category_id"`
}

func (dto *TransactionCreateDTO) ToEntity() *models.Transaction {
	u := &models.Transaction{
		Label:      dto.Label,
		Amount:     dto.Amount,
		CategoryId: dto.CategoryId,
	}

	return u
}

func (dto *TransactionDTO) ToEntity() *models.Transaction {
	u := &models.Transaction{
		Label:      dto.Label,
		Amount:     dto.Amount,
		NoteId:     dto.NoteId,
		CategoryId: dto.CategoryId,
		Timestamp:  dto.Timestamp,
	}

	return u
}

func (dto *TransactionDTO) FromEntity(transaction *models.Transaction) {
	dto.BaseDTO.FromEntity(&transaction.Base)
	dto.Label = transaction.Label
	dto.Amount = transaction.Amount
	dto.NoteId = transaction.NoteId
	dto.CategoryId = transaction.CategoryId
	dto.Timestamp = transaction.Timestamp
}
