package dto

import (
	"time"

	"github.com/moneybackward/backend/models"
	"github.com/moneybackward/backend/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserDTO struct {
	BaseDTO
	Name      string     `json:"name" binding:"required"`
	Email     string     `json:"email" binding:"required"`
	Password  string     `json:"-"`
	LastLogin *time.Time `json:"last_login"`
}

func (dto *UserDTO) ToEntity() (*models.User, error) {
	u := &models.User{
		Name:      dto.Name,
		Email:     dto.Email,
		Password:  dto.Password,
		LastLogin: dto.LastLogin,
	}

	return u, nil
}

func (dto *UserDTO) FromEntity(user *models.User) {
	dto.BaseDTO.FromEntity(&user.Base)
	dto.Name = user.Name
	dto.Email = user.Email
	dto.Password = user.Password
	dto.LastLogin = user.LastLogin
}

type UserRegisterDTO struct {
	Name                 string `json:"name" binding:"required"`
	Email                string `json:"email" binding:"required"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"`
}

func (dto *UserRegisterDTO) Validate() error {
	if dto.Password != dto.PasswordConfirmation {
		return &errors.ValidationError{
			Message: "Password and password confirmation must match",
		}
	}

	return nil
}

func (dto *UserRegisterDTO) ToEntity() (*models.User, error) {
	u := &models.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}

	return u, nil
}

func (dto *UserDTO) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(dto.Password), []byte(password))
}

type UserLoginDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
