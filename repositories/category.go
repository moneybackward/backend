package repositories

import (
	"github.com/google/uuid"
	"github.com/moneybackward/backend/models"
	"github.com/moneybackward/backend/models/dto"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Save(category dto.CategoryCreateDTO) (*dto.CategoryDTO, error)
	Update(category dto.CategoryDTO) (*dto.CategoryDTO, error)
	Find(categoryId uuid.UUID) (*dto.CategoryDTO, error)
	FindAllOfNote(noteId uuid.UUID) ([]dto.CategoryDTO, error)
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

func (repo *categoryRepository) Save(categoryCreate dto.CategoryCreateDTO) (*dto.CategoryDTO, error) {
	category := categoryCreate.ToEntity()
	repo.DB.Create(&category)
	categoryDto := dto.CategoryDTO{}
	categoryDto.FromEntity(category)
	return &categoryDto, repo.DB.Error
}

func (u *categoryRepository) Update(category dto.CategoryDTO) (*dto.CategoryDTO, error) {
	categoryModel := category.ToEntity()
	err := u.DB.Save(&categoryModel).Error
	if err != nil {
		return nil, err
	}

	categoryDto := dto.CategoryDTO{}
	categoryDto.FromEntity(categoryModel)
	return &categoryDto, nil
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
