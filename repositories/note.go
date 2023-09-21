package repositories

import (
	"github.com/moneybackward/backend/models"
	"github.com/moneybackward/backend/models/dao"
	"gorm.io/gorm"
)

type NoteRepository interface {
	Save(Note *dao.NoteDAO) (*dao.NoteDAO, error)
	FindAll() ([]dao.NoteDAO, error)
	FindUserNotes(userId int) ([]dao.NoteDAO, error)
	Delete(Note *dao.NoteDAO) error
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

func (noteRepo *noteRepository) Save(note *dao.NoteDAO) (*dao.NoteDAO, error) {
	noteRepo.DB.Create(&note)
	return note, noteRepo.DB.Error
}

func (noteRepo *noteRepository) FindAll() ([]dao.NoteDAO, error) {
	var notes []dao.NoteDAO
	err := noteRepo.DB.Find(&notes).Error
	return notes, err
}

func (noteRepo *noteRepository) FindUserNotes(userId int) ([]dao.NoteDAO, error) {
	var notes []dao.NoteDAO
	err := noteRepo.DB.Where("user_id = ?", userId).Find(&notes).Error
	return notes, err
}

func (noteRepo *noteRepository) Delete(note *dao.NoteDAO) error {
	return noteRepo.DB.Delete(&note).Error
}

func (noteRepo *noteRepository) Migrate() error {
	return noteRepo.DB.AutoMigrate(&dao.NoteDAO{})
}
