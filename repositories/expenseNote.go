package repositories

import (
	"github.com/moneybackward/backend/models"
	"github.com/moneybackward/backend/models/dao"
	"gorm.io/gorm"
)

type ExpenseNoteRepository interface {
	Save(expenseNote *dao.ExpenseNoteDAO) (*dao.ExpenseNoteDAO, error)
	FindAll() ([]dao.ExpenseNoteDAO, error)
	FindUserExpenseNotes(userId int) ([]dao.ExpenseNoteDAO, error)
	Delete(expenseNote *dao.ExpenseNoteDAO) error
	Migrate() error
}

type expenseNoteRepository struct {
	DB *gorm.DB
}

func NewExpenseNoteRepository() ExpenseNoteRepository {
	return &expenseNoteRepository{
		DB: models.DB,
	}
}

func (u *expenseNoteRepository) Save(expenseNote *dao.ExpenseNoteDAO) (*dao.ExpenseNoteDAO, error) {
	u.DB.Create(&expenseNote)
	return expenseNote, u.DB.Error
}

func (u *expenseNoteRepository) FindAll() ([]dao.ExpenseNoteDAO, error) {
	var expenseNotes []dao.ExpenseNoteDAO
	err := u.DB.Find(&expenseNotes).Error
	return expenseNotes, err
}

func (u *expenseNoteRepository) FindUserExpenseNotes(userId int) ([]dao.ExpenseNoteDAO, error) {
	var expenseNotes []dao.ExpenseNoteDAO
	err := u.DB.Where("user_id = ?", userId).Find(&expenseNotes).Error
	return expenseNotes, err
}

func (u *expenseNoteRepository) Delete(expenseNote *dao.ExpenseNoteDAO) error {
	return u.DB.Delete(&expenseNote).Error
}

func (u *expenseNoteRepository) Migrate() error {
	return u.DB.AutoMigrate(&dao.ExpenseNoteDAO{})
}
