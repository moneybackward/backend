package dto

import (
	"time"

	"github.com/moneybackward/backend/models"
)

type TransactionDTO struct {
	BaseDTO
	Label     string    `json:"label"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`

	NoteId     int `json:"note_id"`
	CategoryId int `json:"category_id"`
}

func (dto *TransactionDTO) ToEntity() (*models.Transaction, error) {
	u := &models.Transaction{
		Label:      dto.Label,
		Amount:     dto.Amount,
		NoteId:     dto.NoteId,
		CategoryId: dto.CategoryId,
		Timestamp:  dto.Timestamp,
	}

	return u, nil
}

func (dto *TransactionDTO) FromEntity(transaction *models.Transaction) {
	dto.BaseDTO.FromEntity(&transaction.Base)
	dto.Label = transaction.Label
	dto.Amount = transaction.Amount
	dto.NoteId = transaction.NoteId
	dto.CategoryId = transaction.CategoryId
	dto.Timestamp = transaction.Timestamp
}
