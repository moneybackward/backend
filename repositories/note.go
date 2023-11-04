package repositories

import (
	"github.com/google/uuid"
	"github.com/moneybackward/backend/models"
	"github.com/moneybackward/backend/models/dto"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type NoteRepository interface {
	Save(userId uuid.UUID, noteCreate *dto.NoteCreateDTO) (*dto.NoteDTO, error)
	Find(noteId uuid.UUID) (*dto.NoteDTO, error)
	FindAll() ([]dto.NoteDTO, error)
	FindUserNotes(userId uuid.UUID) ([]dto.NoteDTO, error)
	Delete(noteId uuid.UUID) error
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

func (noteRepo *noteRepository) Save(userId uuid.UUID, noteCreate *dto.NoteCreateDTO) (*dto.NoteDTO, error) {
	notemodels, err := noteCreate.ToEntity()
	notemodels.UserId = userId
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	noteRepo.DB.Create(&notemodels)
	note := &dto.NoteDTO{}
	note.FromEntity(notemodels)
	return note, noteRepo.DB.Error
}

func (noteRepo *noteRepository) Find(noteId uuid.UUID) (*dto.NoteDTO, error) {
	var note models.Note
	err := noteRepo.DB.First(&note, noteId).Error
	noteDto := &dto.NoteDTO{}
	noteDto.FromEntity(&note)
	return noteDto, err
}

func (noteRepo *noteRepository) FindAll() ([]dto.NoteDTO, error) {
	var result []dto.NoteDTO

	// get the notes
	var notes []models.Note
	err := noteRepo.DB.Find(&notes).Error

	// convert the models to DTOs
	for _, note := range notes {
		noteDto := &dto.NoteDTO{}
		noteDto.FromEntity(&note)
		result = append(result, *noteDto)
	}

	return result, err
}

func (noteRepo *noteRepository) FindUserNotes(userId uuid.UUID) ([]dto.NoteDTO, error) {
	var result []dto.NoteDTO

	// get the notes
	var notes []models.Note
	err := noteRepo.DB.Where("user_id = ?", userId).Find(&notes).Error

	// convert the models to DTOs
	for _, note := range notes {
		noteDto := &dto.NoteDTO{}
		noteDto.FromEntity(&note)
		result = append(result, *noteDto)
	}

	return result, err
}

func (noteRepo *noteRepository) Delete(noteId uuid.UUID) error {
	return noteRepo.DB.Delete(&dto.NoteDTO{}, noteId).Error
}

func (noteRepo *noteRepository) Migrate() error {
	return noteRepo.DB.AutoMigrate(&dto.NoteDTO{})
}
