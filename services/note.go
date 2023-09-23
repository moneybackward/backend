package services

import (
	"log"
	"sync"

	"github.com/moneybackward/backend/models"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/repositories"
)

type NoteService interface {
	Create(note *dto.NoteDTO) (*models.Note, error)
	FindAll() ([]models.Note, error)
	FindUserNotes(userId int) ([]models.Note, error)
	Delete(note *models.Note) error
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

func (noteSvc *noteService) Create(Note *dto.NoteDTO) (*models.Note, error) {
	Notemodels, err := Note.ToEntity()
	if err != nil {
		log.Panic("Failed to convert Note to ")
	}
	return noteSvc.noteRepository.Save(Notemodels)
}

func (noteSvc *noteService) FindAll() ([]models.Note, error) {
	return noteSvc.noteRepository.FindAll()
}

func (noteSvc *noteService) FindUserNotes(userId int) ([]models.Note, error) {
	return noteSvc.noteRepository.FindUserNotes(userId)
}

func (noteSvc *noteService) Delete(note *models.Note) error {
	return noteSvc.noteRepository.Delete(note)
}
