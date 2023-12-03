package services

import (
	"sync"

	"github.com/google/uuid"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/repositories"
)

type TransactionService interface {
	Create(noteId uuid.UUID, transactionCreate *dto.TransactionCreateDTO) (*dto.TransactionDTO, error)
	Update(transactionId uuid.UUID, transactionUpdate dto.TransactionUpdateDTO) (*dto.TransactionDTO, error)
	FindAllOfNote(noteId uuid.UUID) ([]dto.TransactionDTO, error)
	Find(transactionId uuid.UUID) (*dto.TransactionDTO, error)
	Delete(uuid.UUID) error
}

type transactionService struct {
	transactionRepository repositories.TransactionRepository
}

var transactionServiceInstance *transactionService
var transactionOnce sync.Once

func NewTransactionService() TransactionService {
	transactionOnce.Do(func() {
		transactionServiceInstance = &transactionService{
			transactionRepository: repositories.NewTransactionRepository(),
		}
	})
	return transactionServiceInstance
}

func (transactionSvc *transactionService) Create(noteId uuid.UUID, transactionCreate *dto.TransactionCreateDTO) (*dto.TransactionDTO, error) {
	return transactionSvc.transactionRepository.Save(noteId, transactionCreate)
}

func (transactionSvc *transactionService) Update(transactionId uuid.UUID, transactionUpdate dto.TransactionUpdateDTO) (*dto.TransactionDTO, error) {
	return transactionSvc.transactionRepository.Update(transactionId, transactionUpdate)
}

func (transactionSvc *transactionService) Find(transactionId uuid.UUID) (*dto.TransactionDTO, error) {
	return transactionSvc.transactionRepository.Find(transactionId)
}

func (transactionSvc *transactionService) FindAllOfNote(noteId uuid.UUID) ([]dto.TransactionDTO, error) {
	return transactionSvc.transactionRepository.FindAllOfNote(noteId)
}

func (transactionSvc *transactionService) Delete(transactionId uuid.UUID) error {
	return transactionSvc.transactionRepository.Delete(transactionId)
}
