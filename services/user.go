package services

import (
	"log/slog"
	"sync"

	"github.com/google/uuid"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/repositories"
	"github.com/moneybackward/backend/utils/errors"
	"github.com/moneybackward/backend/utils/token"
)

type UserService interface {
	Create(*dto.UserRegisterDTO) (*dto.UserDTO, error)
	FindAll() ([]dto.UserDTO, error)
	Find(uuid.UUID) (*dto.UserDTO, error)
	FindByEmail(string) (*dto.UserDTO, error)
	Delete(uuid.UUID) error
	Login(*dto.UserLoginDTO) (string, error)
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

	existingUser, _ := userSvc.FindByEmail(user.Email)
	if existingUser != nil {
		slog.Warn("Email already exist")
		return nil, &errors.ConflictError{Message: "Email already exist"}
	}

	createdUser, err := userSvc.userRepository.Save(user)
	if err != nil {
		slog.Error("Failed to save user")
		return nil, err
	}
	return createdUser, nil
}

func (userSvc *userService) FindByEmail(email string) (*dto.UserDTO, error) {
	userDTO, err := userSvc.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return userDTO, nil
}

func (userSvc *userService) FindAll() ([]dto.UserDTO, error) {
	users := []dto.UserDTO{}

	// Get all users from database
	usersDTO, err := userSvc.userRepository.FindAll()
	if err != nil {
		return users, err
	}

	return usersDTO, nil
}

func (userSvc *userService) Find(userId uuid.UUID) (*dto.UserDTO, error) {
	userDTO, err := userSvc.userRepository.Find(userId)
	if err != nil {
		return nil, err
	}
	return userDTO, nil
}

func (userSvc *userService) Delete(userId uuid.UUID) error {
	return userSvc.userRepository.Delete(userId)
}

func (userSvc *userService) Login(user *dto.UserLoginDTO) (string, error) {
	userDTO, err := userSvc.userRepository.FindByEmail(user.Email)
	if err != nil {
		return "", &errors.UnauthorizedError{Message: "Invalid email or password"}
	}

	err = userDTO.VerifyPassword(user.Password)
	if err != nil {
		slog.Error("Failed to verify password")
		return "", &errors.UnauthorizedError{Message: "Invalid email or password"}
	}

	token, err := token.GenerateToken(userDTO)
	if err != nil {
		slog.Error("Failed to generate token")
		return "", err
	}

	return token, nil
}
