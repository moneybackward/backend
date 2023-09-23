package repositories

import (
	"github.com/moneybackward/backend/models"
	"gorm.io/gorm"
)

type NoteRepository interface {
	Save(Note *models.Note) (*models.Note, error)
	FindAll() ([]models.Note, error)
	FindUserNotes(userId int) ([]models.Note, error)
	Delete(Note *models.Note) error
	Migrate() error
}

type noteRepository struct {
	DB *gorm.DB
}

func NewNoteRepository() NoteRepository {
	return &noteRepository{
		DB: models.DB,
	}
}

func (noteRepo *noteRepository) Save(note *models.Note) (*models.Note, error) {
	noteRepo.DB.Create(&note)
	return note, noteRepo.DB.Error
}

func (noteRepo *noteRepository) FindAll() ([]models.Note, error) {
	var notes []models.Note
	err := noteRepo.DB.Find(&notes).Error
	return notes, err
}

func (noteRepo *noteRepository) FindUserNotes(userId int) ([]models.Note, error) {
	var notes []models.Note
	err := noteRepo.DB.Where("user_id = ?", userId).Find(&notes).Error
	return notes, err
}

func (noteRepo *noteRepository) Delete(note *models.Note) error {
	return noteRepo.DB.Delete(&note).Error
}

func (noteRepo *noteRepository) Migrate() error {
	return noteRepo.DB.AutoMigrate(&models.Note{})
}
