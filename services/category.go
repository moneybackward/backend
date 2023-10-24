package services

import (
	"sync"

	"github.com/google/uuid"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/repositories"
)

type CategoryService interface {
	Create(category dto.CategoryCreateDTO) (dto.CategoryDTO, error)
	FindAll(noteId uuid.UUID) ([]dto.CategoryDTO, error)
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

func (categorySvc *categoryService) Create(categoryCreateDto dto.CategoryCreateDTO) (dto.CategoryDTO, error) {
	return categorySvc.categoryRepository.Save(categoryCreateDto)
}

func (categorySvc *categoryService) FindAll(noteId uuid.UUID) ([]dto.CategoryDTO, error) {
	return categorySvc.categoryRepository.FindAll(noteId)
}

func (categorySvc *categoryService) Delete(categoryId uuid.UUID) error {
	return categorySvc.categoryRepository.Delete(categoryId)
}
