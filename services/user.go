package services

import (
	"log"
	"sync"

	"github.com/moneybackward/backend/models"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/repositories"
)

type UserService interface {
	Create(user *dto.UserDTO) (*models.User, error)
	FindAll() ([]models.User, error)
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

func (userSvc *userService) Create(user *dto.UserDTO) (*models.User, error) {
	usermodels, err := user.ToEntity()
	if err != nil {
		log.Panic("Failed to convert user to ")
	}
	return userSvc.userRepository.Save(usermodels)
}

func (u *userService) FindAll() ([]models.User, error) {
	return u.userRepository.FindAll()
}
