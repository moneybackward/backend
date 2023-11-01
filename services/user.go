package services

import (
	"log/slog"
	"sync"

	"github.com/google/uuid"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/repositories"
)

type UserService interface {
	Create(user *dto.UserRegisterDTO) (*dto.UserDTO, error)
	FindAll() ([]dto.UserDTO, error)
	Find(userId uuid.UUID) (*dto.UserDTO, error)
	Delete(userId uuid.UUID) error
}

type userService struct {
	userRepository repositories.UserRepository
}

var userServiceInstance *userService
var userOnce sync.Once

func NewUserService() UserService {
	userOnce.Do(func() {
		userServiceInstance = &userService{
			userRepository: repositories.NewUserRepository(),
		}
	})
	return userServiceInstance
}

func (userSvc *userService) Create(user *dto.UserRegisterDTO) (*dto.UserDTO, error) {
	err := user.Validate()
	if err != nil {
		slog.Warn("Failed to validate user")
		return nil, err
	}

	usermodels, err := user.ToEntity()
	if err != nil {
		slog.Error("Failed to convert user to ")
		return nil, err
	}
	userModel, err := userSvc.userRepository.Save(usermodels)
	if err != nil {
		slog.Error("Failed to save user")
		return nil, err
	}
	createdUser := &dto.UserDTO{}
	err = createdUser.FromEntity(userModel)
	if err != nil {
		slog.Error("Failed to convert user to dto")
		return nil, err
	}
	return createdUser, nil
}

func (userSvc *userService) FindAll() ([]dto.UserDTO, error) {
	users := []dto.UserDTO{}

	// Get all users from database
	userModels, err := userSvc.userRepository.FindAll()
	if err != nil {
		return users, err
	}

	// Convert user models to user dtos
	for _, userModel := range userModels {
		userDTO := &dto.UserDTO{}
		err := userDTO.FromEntity(&userModel)
		if err != nil {
			slog.Error("Failed to convert user to dto")
			continue
		}
		users = append(users, *userDTO)
	}

	return users, nil
}

func (userSvc *userService) Find(userId uuid.UUID) (*dto.UserDTO, error) {
	userModel, err := userSvc.userRepository.Find(userId)
	if err != nil {
		return nil, err
	}
	userDTO := &dto.UserDTO{}
	err = userDTO.FromEntity(userModel)
	if err != nil {
		slog.Error("Failed to convert user to dto")
		return nil, err
	}
	return userDTO, nil
}

func (userSvc *userService) Delete(userId uuid.UUID) error {
	return userSvc.userRepository.Delete(userId)
}
