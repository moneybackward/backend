package services

import (
	"log"
	"sync"

	"github.com/moneybackward/backend/models/dao"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/repositories"
)

type UserService interface {
	Create(user *dto.UserDTO) (*dao.UserDAO, error)
	FindAll() ([]dao.UserDAO, error)
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

func (userSvc *userService) Create(user *dto.UserDTO) (*dao.UserDAO, error) {
	userDao, err := user.ToDAO()
	if err != nil {
		log.Panic("Failed to convert user to DAO")
	}
	return userSvc.userRepository.Save(userDao)
}

func (u *userService) FindAll() ([]dao.UserDAO, error) {
	return u.userRepository.FindAll()
}
