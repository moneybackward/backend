package dto

import "github.com/moneybackward/backend/models/dao"

type NoteDTO struct {
	Name   string `json:"name" binding:"required"`
	UserId int    `json:"user_id" binding:"required"`
}

func (dto *NoteDTO) ToEntity() (*dao.NoteDAO, error) {
	u := &dao.NoteDAO{
		Name:   dto.Name,
		UserId: dto.UserId,
	}

	return u, nil
}
