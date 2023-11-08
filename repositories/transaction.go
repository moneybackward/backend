package repositories

import (
	"github.com/google/uuid"
	"github.com/moneybackward/backend/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Save(*models.Transaction) (*models.Transaction, error)
	FindAll(uuid.UUID) ([]models.Transaction, error)
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

func (u *transactionRepository) Save(transaction *models.Transaction) (*models.Transaction, error) {
	u.DB.Create(&transaction)
	return transaction, u.DB.Error
}

func (u *transactionRepository) FindAll(noteId uuid.UUID) ([]models.Transaction, error) {
	var categories []models.Transaction
	err := u.DB.Where("note_id = ?", noteId).Find(&categories).Error
	return categories, err
}

func (u *transactionRepository) Delete(transactionId uuid.UUID) error {
	return u.DB.Delete(&models.Transaction{}, transactionId).Error
}
