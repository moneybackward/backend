package repositories

import (
	"github.com/google/uuid"
	"github.com/moneybackward/backend/models"
	"github.com/moneybackward/backend/models/dto"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Save(uuid.UUID, dto.CategoryCreateDTO) (*dto.CategoryDTO, error)
	Update(categoryId uuid.UUID, categoryUpdateDto dto.CategoryUpdateDTO) (*dto.CategoryDTO, error)
	Find(uuid.UUID) (*dto.CategoryDTO, error)
	FindAllOfNote(uuid.UUID) ([]dto.CategoryDTO, error)
	Delete(uuid.UUID) error
}

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{
		DB: models.DB,
	}
}

func (repo *categoryRepository) Save(noteId uuid.UUID, categoryCreate dto.CategoryCreateDTO) (*dto.CategoryDTO, error) {
	category := categoryCreate.ToEntity()
	category.NoteId = noteId

	repo.DB.Create(&category)
	categoryDto := dto.CategoryDTO{}
	categoryDto.FromEntity(category)
	return &categoryDto, repo.DB.Error
}

func (u *categoryRepository) Update(categoryId uuid.UUID, categoryUpdateDto dto.CategoryUpdateDTO) (*dto.CategoryDTO, error) {
	var category models.Category
	err := u.DB.First(&category, categoryId).Error
	if err != nil {
		return nil, err
	}

	category.Name = categoryUpdateDto.Name
	category.Priority = categoryUpdateDto.Priority
	category.Budget = categoryUpdateDto.Budget
	u.DB.Save(&category)
	categoryDto := dto.CategoryDTO{}
	categoryDto.FromEntity(category)
	return &categoryDto, u.DB.Error
}

func (u *categoryRepository) Find(categoryId uuid.UUID) (*dto.CategoryDTO, error) {
	var category models.Category
	err := u.DB.First(&category, categoryId).Error
	if err != nil {
		return nil, err
	}

	categoryDto := dto.CategoryDTO{}
	categoryDto.FromEntity(category)
	return &categoryDto, nil
}

func (u *categoryRepository) FindAllOfNote(noteId uuid.UUID) ([]dto.CategoryDTO, error) {
	var categories []models.Category
	err := u.DB.Where("note_id = ?", noteId).Find(&categories).Error
	if err != nil {
		return nil, err
	}

	var categoryDtos []dto.CategoryDTO
	for _, category := range categories {
		categoryDto := dto.CategoryDTO{}
		categoryDto.FromEntity(category)
		categoryDtos = append(categoryDtos, categoryDto)
	}

	return categoryDtos, nil
}

func (u *categoryRepository) Delete(categoryId uuid.UUID) error {
	return u.DB.Delete(&models.Category{}, categoryId).Error
}
