package services

import (
	"sync"

	"github.com/google/uuid"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/repositories"
	"github.com/rs/zerolog/log"
)

type CategoryService interface {
	Create(uuid.UUID, dto.CategoryCreateDTO) (*dto.CategoryDTO, error)
	Update(categoryId uuid.UUID, categoryUpdateDto dto.CategoryUpdateDTO) (*dto.CategoryDTO, error)
	Find(uuid.UUID) (*dto.CategoryDTO, error)
	FindAllOfNote(noteId uuid.UUID, isExpense *bool) ([]dto.CategoryDTO, error)
	IsBelongsToNote(uuid.UUID, uuid.UUID) bool
	Delete(uuid.UUID) error
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

func (categorySvc *categoryService) Create(noteId uuid.UUID, categoryCreateDto dto.CategoryCreateDTO) (*dto.CategoryDTO, error) {
	return categorySvc.categoryRepository.Save(noteId, categoryCreateDto)
}

func (categorySvc *categoryService) Update(categoryId uuid.UUID, categoryUpdateDto dto.CategoryUpdateDTO) (*dto.CategoryDTO, error) {
	return categorySvc.categoryRepository.Update(categoryId, categoryUpdateDto)
}

func (categorySvc *categoryService) Find(categoryId uuid.UUID) (*dto.CategoryDTO, error) {
	return categorySvc.categoryRepository.Find(categoryId)
}

func (categorySvc *categoryService) FindAllOfNote(noteId uuid.UUID, isExpense *bool) ([]dto.CategoryDTO, error) {
	return categorySvc.categoryRepository.FindAllOfNote(noteId, isExpense)
}

func (categorySvc *categoryService) IsBelongsToNote(categoryId uuid.UUID, noteId uuid.UUID) bool {
	category, err := categorySvc.Find(categoryId)
	if err != nil {
		log.Error().Msg("Category not found")
		return false
	}

	return category.NoteId == noteId
}

func (categorySvc *categoryService) Delete(categoryId uuid.UUID) error {
	return categorySvc.categoryRepository.Delete(categoryId)
}
