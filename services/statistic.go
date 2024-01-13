package services

import (
	"sync"

	"github.com/google/uuid"
	"github.com/moneybackward/backend/models/dto"
)

type StatisticService interface {
	Categories(noteId uuid.UUID, isExpense *bool) ([]dto.CategoryStatsDTO, error)
}

type statisticService struct {
	categoryService CategoryService
}

var statisticServiceInstance *statisticService
var statisticOnce sync.Once

func NewStatisticService() StatisticService {
	statisticOnce.Do(func() {
		statisticServiceInstance = &statisticService{
			categoryService: NewCategoryService(),
		}
	})
	return statisticServiceInstance
}

func (statisticSvc *statisticService) Categories(noteId uuid.UUID, isExpense *bool) ([]dto.CategoryStatsDTO, error) {
	return statisticSvc.categoryService.GetStats(noteId, isExpense)
}
