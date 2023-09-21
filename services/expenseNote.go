package services

import (
	"log"
	"sync"

	"github.com/moneybackward/backend/models/dao"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/repositories"
)

type ExpenseNoteService interface {
	Create(expenseNote *dto.ExpenseNoteDTO) (*dao.ExpenseNoteDAO, error)
	FindAll() ([]dao.ExpenseNoteDAO, error)
	FindUserExpenseNotes(userId int) ([]dao.ExpenseNoteDAO, error)
	Delete(expenseNote *dao.ExpenseNoteDAO) error
}

type expenseNoteService struct {
	expenseNoteRepository repositories.ExpenseNoteRepository
}

var expenseNoteServiceInstance *expenseNoteService
var expenseNoteOnce sync.Once

func NewExpenseNoteService() ExpenseNoteService {
	expenseNoteOnce.Do(func() {
		expenseNoteServiceInstance = &expenseNoteService{
			expenseNoteRepository: repositories.NewExpenseNoteRepository(),
		}
	})
	return expenseNoteServiceInstance
}

func (expenseNoteSvc *expenseNoteService) Create(expenseNote *dto.ExpenseNoteDTO) (*dao.ExpenseNoteDAO, error) {
	expenseNoteDao, err := expenseNote.ToEntity()
	if err != nil {
		log.Panic("Failed to convert expenseNote to DAO")
	}
	return expenseNoteSvc.expenseNoteRepository.Save(expenseNoteDao)
}

func (u *expenseNoteService) FindAll() ([]dao.ExpenseNoteDAO, error) {
	return u.expenseNoteRepository.FindAll()
}

func (u *expenseNoteService) FindUserExpenseNotes(userId int) ([]dao.ExpenseNoteDAO, error) {
	return u.expenseNoteRepository.FindUserExpenseNotes(userId)
}

func (u *expenseNoteService) Delete(expenseNote *dao.ExpenseNoteDAO) error {
	return u.expenseNoteRepository.Delete(expenseNote)
}
