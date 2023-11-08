package services

import (
	"sync"

	"github.com/google/uuid"
	"github.com/moneybackward/backend/models"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/repositories"
)

type TransactionService interface {
	Create(*dto.TransactionDTO) (*models.Transaction, error)
	FindAll(uuid.UUID) ([]models.Transaction, error)
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

func (transactionSvc *transactionService) Create(transaction *dto.TransactionDTO) (*models.Transaction, error) {
	transactionmodels, err := transaction.ToEntity()
	if err != nil {
		return nil, err
	}
	return transactionSvc.transactionRepository.Save(transactionmodels)
}

func (transactionSvc *transactionService) FindAll(noteId uuid.UUID) ([]models.Transaction, error) {
	return transactionSvc.transactionRepository.FindAll(noteId)
}

func (transactionSvc *transactionService) Delete(transactionId uuid.UUID) error {
	return transactionSvc.transactionRepository.Delete(transactionId)
}
