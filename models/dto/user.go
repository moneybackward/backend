package dto

import "github.com/moneybackward/backend/models"

type UserDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (dto *UserDTO) ToEntity() (*models.User, error) {
	u := &models.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}

	return u, nil
}
