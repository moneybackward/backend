package repositories

import (
	"github.com/moneybackward/backend/models/dao"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *dao.UserDAO) (*dao.UserDAO, error)
	FindAll() ([]dao.UserDAO, error)
	Delete(user *dao.UserDAO) error
	Migrate() error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (u *userRepository) Save(user *dao.UserDAO) (*dao.UserDAO, error) {
	u.DB.Create(&user)
	return user, u.DB.Error
}

func (u *userRepository) FindAll() ([]dao.UserDAO, error) {
	var users []dao.UserDAO
	err := u.DB.Find(&users).Error
	return users, err
}

func (u *userRepository) Delete(user *dao.UserDAO) error {
	return u.DB.Delete(&user).Error
}

func (u *userRepository) Migrate() error {
	return u.DB.AutoMigrate(&dao.UserDAO{})
}
