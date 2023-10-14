package repositories

import (
	"github.com/google/uuid"
	"github.com/moneybackward/backend/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Save(category *models.Category) (*models.Category, error)
	FindAll() ([]models.Category, error)
	Delete(categoryId uuid.UUID) error
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

func (u *categoryRepository) Delete(categoryId uuid.UUID) error {
	return u.DB.Delete(&models.Category{}, categoryId).Error
}
