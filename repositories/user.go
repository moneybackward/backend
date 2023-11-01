package repositories

import (
	"github.com/google/uuid"
	"github.com/moneybackward/backend/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *models.User) (*models.User, error)
	FindAll() ([]models.User, error)
	Find(userId uuid.UUID) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Delete(userId uuid.UUID) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		DB: models.DB,
	}
}

func (u *userRepository) Save(user *models.User) (*models.User, error) {
	u.DB.Create(&user)
	return user, u.DB.Error
}

func (u *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := u.DB.Find(&users).Error
	return users, err
}

func (userRepo *userRepository) Find(userId uuid.UUID) (*models.User, error) {
	var user models.User
	err := userRepo.DB.Where("id = ?", userId).First(&user).Error
	return &user, err
}

func (userRepo *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := userRepo.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (userRepo *userRepository) Delete(userId uuid.UUID) error {
	return userRepo.DB.Delete(&models.User{}, userId).Error
}
