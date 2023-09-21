package dto

import "github.com/moneybackward/backend/models/dao"

type ExpenseNoteDTO struct {
	Name   string `json:"name" binding:"required"`
	UserId int    `json:"user_id" binding:"required"`
}

func (dto *ExpenseNoteDTO) ToEntity() (*dao.ExpenseNoteDAO, error) {
	u := &dao.ExpenseNoteDAO{
		Name:   dto.Name,
		UserId: dto.UserId,
	}

	return u, nil
}
