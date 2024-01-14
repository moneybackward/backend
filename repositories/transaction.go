package repositories

import (
	"github.com/google/uuid"
	"github.com/moneybackward/backend/models"
	"github.com/moneybackward/backend/models/dto"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Save(noteId uuid.UUID, transactionCreate *dto.TransactionCreateDTO) (*dto.TransactionDTO, error)
	Update(transactionId uuid.UUID, transactionUpdate dto.TransactionUpdateDTO) (*dto.TransactionDTO, error)
	Find(transactionId uuid.UUID) (*dto.TransactionDTO, error)
	FindAllOfNote(noteId uuid.UUID) ([]dto.TransactionDTO, error)
	Delete(uuid.UUID) error
}

type transactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{
		DB: models.DB,
	}
}

func (u *transactionRepository) Save(noteId uuid.UUID, transactionCreate *dto.TransactionCreateDTO) (*dto.TransactionDTO, error) {
	transactionModel := transactionCreate.ToEntity()
	transactionModel.NoteId = noteId
	u.DB.Create(&transactionModel)

	transactionDto := dto.TransactionDTO{}
	transactionDto.FromEntity(transactionModel)
	return &transactionDto, u.DB.Error
}

func (u *transactionRepository) Update(transactionId uuid.UUID, transactionUpdate dto.TransactionUpdateDTO) (*dto.TransactionDTO, error) {
	var transaction models.Transaction
	err := u.DB.First(&transaction, transactionId).Error
	if err != nil {
		return nil, err
	}

	transaction.Label = transactionUpdate.Label
	transaction.Amount = transactionUpdate.Amount
	transaction.CategoryId = transactionUpdate.CategoryId
	transaction.IsExpense = transactionUpdate.IsExpense
	u.DB.Save(&transaction)
	transactionDto := dto.TransactionDTO{}
	transactionDto.FromEntity(&transaction)
	return &transactionDto, u.DB.Error
}

func (u *transactionRepository) Find(transactionId uuid.UUID) (*dto.TransactionDTO, error) {
	var transaction models.Transaction
	err := u.DB.First(&transaction, transactionId).Error
	if err != nil {
		return nil, err
	}

	transactionDto := dto.TransactionDTO{}
	transactionDto.FromEntity(&transaction)
	return &transactionDto, nil
}

func (u *transactionRepository) FindAllOfNote(noteId uuid.UUID) ([]dto.TransactionDTO, error) {
	var transactions []models.Transaction
	err := u.DB.Where("Transaction.note_id = ?", noteId).
		Joins("Category").
		Order("date DESC").
		Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	var transactionDtos []dto.TransactionDTO
	for _, transaction := range transactions {
		transactionDto := dto.TransactionDTO{}
		transactionDto.FromEntity(&transaction)
		transactionDtos = append(transactionDtos, transactionDto)
	}

	return transactionDtos, nil
}

func (u *transactionRepository) Delete(transactionId uuid.UUID) error {
	return u.DB.Delete(&models.Transaction{}, transactionId).Error
}
