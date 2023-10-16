package services

import (
	"sync"

	"github.com/google/uuid"
	"github.com/moneybackward/backend/models"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/repositories"
)

type CategoryService interface {
	Create(category *dto.CategoryDTO) (*models.Category, error)
	FindAll(noteId uuid.UUID) ([]models.Category, error)
	Delete(categoryId uuid.UUID) error
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

func (categorySvc *categoryService) FindAll(noteId uuid.UUID) ([]models.Category, error) {
	return categorySvc.categoryRepository.FindAll(noteId)
}

func (categorySvc *categoryService) Delete(categoryId uuid.UUID) error {
	return categorySvc.categoryRepository.Delete(categoryId)
}
