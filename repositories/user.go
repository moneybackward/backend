package repositories

import (
	"log/slog"

	"github.com/google/uuid"
	"github.com/moneybackward/backend/models"
	"github.com/moneybackward/backend/models/dto"
	"gorm.io/gorm"
)

var UserRepositoryInstance UserRepository

type UserRepository interface {
	Save(*dto.UserRegisterDTO) (*dto.UserDTO, error)
	FindAll() ([]dto.UserDTO, error)
	Find(uuid.UUID) (*dto.UserDTO, error)
	FindByEmail(string) (*dto.UserDTO, error)
	Delete(uuid.UUID) error
}

type userRepository struct {
	DB *gorm.DB
}

func newUserRepository() *userRepository {
	return &userRepository{
		DB: models.DB,
	}
}

func NewUserRepository() UserRepository {
	if UserRepositoryInstance == nil {
		UserRepositoryInstance = newUserRepository()
	}
	return UserRepositoryInstance
}

func (u *userRepository) Save(user *dto.UserRegisterDTO) (*dto.UserDTO, error) {
	userModel, err := user.ToEntity()
	if err != nil {
		slog.Error("Failed to convert user to ")
		return nil, err
	}
	u.DB.Create(&userModel)
	createdUser := &dto.UserDTO{}
	createdUser.FromEntity(userModel)
	return createdUser, u.DB.Error
}

func (u *userRepository) FindAll() ([]dto.UserDTO, error) {
	var result []dto.UserDTO

	var users []models.User
	err := u.DB.Find(&users).Error

	// convert the models to DTOs
	for _, user := range users {
		userDto := &dto.UserDTO{}
		userDto.FromEntity(&user)
		result = append(result, *userDto)
	}
	return result, err
}

func (userRepo *userRepository) Find(userId uuid.UUID) (*dto.UserDTO, error) {
	var user *models.User
	err := userRepo.DB.Where("id = ?", userId).First(&user).Error
	userDTO := &dto.UserDTO{}
	userDTO.FromEntity(user)
	return userDTO, err
}

func (userRepo *userRepository) FindByEmail(email string) (*dto.UserDTO, error) {
	var user *models.User
	err := userRepo.DB.Where("email = ?", email).First(&user).Error
	userDTO := &dto.UserDTO{}
	userDTO.FromEntity(user)
	return userDTO, err
}

func (userRepo *userRepository) Delete(userId uuid.UUID) error {
	return userRepo.DB.Delete(&models.User{}, userId).Error
}
