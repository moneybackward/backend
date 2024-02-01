package repositories

import (
	"github.com/google/uuid"
	"github.com/guregu/null"
	"github.com/moneybackward/backend/models"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/utils"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Save(uuid.UUID, dto.CategoryCreateDTO) (*dto.CategoryDTO, error)
	Update(categoryId uuid.UUID, categoryUpdateDto dto.CategoryUpdateDTO) (*dto.CategoryDTO, error)
	Find(uuid.UUID) (*dto.CategoryDTO, error)
	FindAllOfNote(noteId uuid.UUID, isExpense *bool) ([]dto.CategoryDTO, error)
	Delete(uuid.UUID) error
	GetStats(noteId uuid.UUID, isExpense *bool, dateFilter *utils.DateFilter) ([]dto.CategoryStatsDTO, error)
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

	nullableBudget := null.NewFloat(0, false)
	if categoryUpdateDto.Budget != nil {
		nullableBudget = null.FloatFromPtr(categoryUpdateDto.Budget)
	}

	category.Name = categoryUpdateDto.Name
	category.Priority = categoryUpdateDto.Priority
	category.Budget = nullableBudget
	category.IsExpense = categoryUpdateDto.IsExpense
	category.Color = categoryUpdateDto.Color
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

func (u *categoryRepository) FindAllOfNote(noteId uuid.UUID, isExpense *bool) ([]dto.CategoryDTO, error) {
	var categories []models.Category
	query := u.DB.Where("note_id = ?", noteId)
	if isExpense != nil {
		query = query.Where("is_expense = ?", *isExpense)
	}

	err := query.
		Order("priority ASC").
		Order("name ASC").
		Order("created_at DESC").
		Find(&categories).Error

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

func (u *categoryRepository) GetStats(noteId uuid.UUID, isExpense *bool, dateFilter *utils.DateFilter) ([]dto.CategoryStatsDTO, error) {
	query := u.DB.Table("categories").
		Select("categories.*, SUM(transactions.amount) as total, COUNT(transactions) as count, (CASE WHEN categories.budget IS NOT NULL AND categories.budget != 0 THEN SUM(transactions.amount) / categories.budget * 100 ELSE 0 END) as percentage").
		Where("categories.deleted_at IS NULL").
		Where("transactions.deleted_at IS NULL").
		Joins("INNER JOIN transactions ON transactions.category_id = categories.id").
		Where("categories.note_id = ?", noteId)

	if isExpense != nil {
		query = query.Where("categories.is_expense = ?", *isExpense)
	}

	if dateFilter != nil {
		query = dateFilter.WhereTransaction(query)
	}

	var categoryStatsDtos []dto.CategoryStatsDTO
	err := query.Group("categories.id").
		Order("percentage DESC").
		Find(&categoryStatsDtos).
		Error

	if err != nil {
		return nil, err
	}

	return categoryStatsDtos, nil
}
