package services

import (
	"sync"

	"github.com/google/uuid"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/repositories"
	"github.com/rs/zerolog/log"
)

type NoteService interface {
	Create(uuid.UUID, *dto.NoteCreateDTO) (*dto.NoteDTO, error)
	Find(uuid.UUID) (*dto.NoteDTO, error)
	FindAll() ([]dto.NoteDTO, error)
	FindUserNotes(uuid.UUID) ([]dto.NoteDTO, error)
	IsBelongsToUser(uuid.UUID, uuid.UUID) bool
	Delete(uuid.UUID) error
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

func (noteSvc *noteService) Create(userId uuid.UUID, note *dto.NoteCreateDTO) (*dto.NoteDTO, error) {
	result, err := noteSvc.noteRepository.Save(userId, note)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	return result, nil
}

func (noteSvc *noteService) Find(noteId uuid.UUID) (*dto.NoteDTO, error) {
	return noteSvc.noteRepository.Find(noteId)
}

func (noteSvc *noteService) FindAll() ([]dto.NoteDTO, error) {
	return noteSvc.noteRepository.FindAll()
}

func (noteSvc *noteService) FindUserNotes(userId uuid.UUID) ([]dto.NoteDTO, error) {
	return noteSvc.noteRepository.FindUserNotes(userId)
}

func (noteSvc *noteService) Delete(noteId uuid.UUID) error {
	return noteSvc.noteRepository.Delete(noteId)
}

func (noteSvc *noteService) IsBelongsToUser(noteId uuid.UUID, userId uuid.UUID) bool {
	note, error := noteSvc.Find(noteId)
	if error != nil {
		return false
	}
	return note.UserId == userId
}
