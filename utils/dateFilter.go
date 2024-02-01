package utils

import (
	"time"

	"gorm.io/gorm"
)

type DateFilter struct {
	Start *time.Time
	End   *time.Time
}

func (dateFilter *DateFilter) WhereTransaction(query *gorm.DB) *gorm.DB {
	if dateFilter.Start != nil {
		query = query.Where("transactions.date >= ?", *dateFilter.Start)
	}

	if dateFilter.End != nil {
		query = query.Where("transactions.date <= ?", *dateFilter.End)
	}
	return query
}

func NewDateFilter(start *time.Time, end *time.Time) DateFilter {
	// Get the current date and time
	now := time.Now()

	// get first day of month if start is not provided
	if start == nil {
		// Get the first day of the current month
		firstDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		start = &firstDay
	}

	// get today if end is not provided
	if end == nil {
		// Calculate the last day of the current month
		lastDay := time.Date(now.Year(), now.Month()+1, 0, 23, 59, 59, 0, now.Location())
		end = &lastDay
	}

	return DateFilter{
		Start: start,
		End:   end,
	}
}
