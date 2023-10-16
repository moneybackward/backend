package dto

import (
	"time"

	"github.com/moneybackward/backend/models"
)

type TransactionDTO struct {
	Label     string    `json:"label"`
	Amount    int       `json:"amount"`
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
