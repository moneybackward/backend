package repositories

import (
	"github.com/moneybackward/backend/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Save(category *models.Category) (*models.Category, error)
	FindAll() ([]models.Category, error)
	Delete(category *models.Category) error
}

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{
		DB: models.DB,
	}
}

func (u *categoryRepository) Save(category *models.Category) (*models.Category, error) {
	u.DB.Create(&category)
	return category, u.DB.Error
}

func (u *categoryRepository) FindAll() ([]models.Category, error) {
	var categories []models.Category
	err := u.DB.Find(&categories).Error
	return categories, err
}

func (u *categoryRepository) Delete(category *models.Category) error {
	return u.DB.Delete(&category).Error
}
