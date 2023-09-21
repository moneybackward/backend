package services

import (
	"log"
	"sync"

	"github.com/moneybackward/backend/models/dao"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/repositories"
)

type NoteService interface {
	Create(note *dto.NoteDTO) (*dao.NoteDAO, error)
	FindAll() ([]dao.NoteDAO, error)
	FindUserNotes(userId int) ([]dao.NoteDAO, error)
	Delete(note *dao.NoteDAO) error
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

func (noteSvc *noteService) Create(Note *dto.NoteDTO) (*dao.NoteDAO, error) {
	NoteDao, err := Note.ToEntity()
	if err != nil {
		log.Panic("Failed to convert Note to DAO")
	}
	return noteSvc.noteRepository.Save(NoteDao)
}

func (noteSvc *noteService) FindAll() ([]dao.NoteDAO, error) {
	return noteSvc.noteRepository.FindAll()
}

func (noteSvc *noteService) FindUserNotes(userId int) ([]dao.NoteDAO, error) {
	return noteSvc.noteRepository.FindUserNotes(userId)
}

func (noteSvc *noteService) Delete(note *dao.NoteDAO) error {
	return noteSvc.noteRepository.Delete(note)
}
