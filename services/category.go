package services

import (
	"sync"

	"github.com/moneybackward/backend/models"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/repositories"
)

type CategoryService interface {
	Create(category *dto.CategoryDTO) (*models.Category, error)
	FindAll() ([]models.Category, error)
}

type categoryService struct {
	categoryRepository repositories.CategoryRepository
}

var categoryServiceInstance *categoryService
var categoryOnce sync.Once

func NewCategoryService() CategoryService {
	categoryOnce.Do(func() {
		categoryServiceInstance = &categoryService{
			categoryRepository: repositories.NewCategoryRepository(),
		}
	})
	return categoryServiceInstance
}

func (categorySvc *categoryService) Create(category *dto.CategoryDTO) (*models.Category, error) {
	categorymodels, err := category.ToEntity()
	if err != nil {
		return nil, err
	}
	return categorySvc.categoryRepository.Save(categorymodels)
}

func (u *categoryService) FindAll() ([]models.Category, error) {
	return u.categoryRepository.FindAll()
}
