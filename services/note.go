package services

import (
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/moneybackward/backend/models"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/repositories"
)

type NoteService interface {
	Create(userId uuid.UUID, note *dto.NoteCreateDTO) (*models.Note, error)
	FindAll(userId uuid.UUID) ([]models.Note, error)
	FindUserNotes(userId uuid.UUID) ([]models.Note, error)
	Delete(noteId uuid.UUID) error
}

type noteService struct {
	noteRepository repositories.NoteRepository
}

var noteServiceInstance *noteService
var noteOnce sync.Once

func NewNoteService() NoteService {
	noteOnce.Do(func() {
		noteServiceInstance = &noteService{
			noteRepository: repositories.NewNoteRepository(),
		}
	})
	return noteServiceInstance
}

func (noteSvc *noteService) Create(userId uuid.UUID, note *dto.NoteCreateDTO) (*models.Note, error) {
	notemodels, err := note.ToEntity()
	notemodels.UserId = userId
	if err != nil {
		log.Panic("Failed to convert Note to ")
	}
	return noteSvc.noteRepository.Save(notemodels)
}

func (noteSvc *noteService) FindAll(userId uuid.UUID) ([]models.Note, error) {
	return noteSvc.noteRepository.FindAll(userId)
}

func (noteSvc *noteService) FindUserNotes(userId uuid.UUID) ([]models.Note, error) {
	return noteSvc.noteRepository.FindUserNotes(userId)
}

func (noteSvc *noteService) Delete(noteId uuid.UUID) error {
	return noteSvc.noteRepository.Delete(noteId)
}
