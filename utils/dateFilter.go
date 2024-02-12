package utils

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type DateFilter struct {
	Start *time.Time
	End   *time.Time
}

func (dateFilter *DateFilter) WhereBetween(query *gorm.DB, key string) *gorm.DB {
	if dateFilter.Start != nil {
		query = query.Where(fmt.Sprintf("%s >= ?", key), *dateFilter.Start)
	}

	if dateFilter.End != nil {
		query = query.Where(fmt.Sprintf("%s <= ?", key), *dateFilter.End)
	}
	return query
}

func newDateFilter(start *time.Time, end *time.Time) DateFilter {
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

func parseDate(rawString string) (*time.Time, error) {
	layout := "2006-01-02"
	parsedDateStart, err := time.Parse(layout, rawString)
	if err != nil {
		return nil, err
	}

	return &parsedDateStart, nil
}

func alignDate(date *time.Time, isStart bool) *time.Time {
	if date == nil {
		return nil
	}

	// make the date start at 00:00:00
	if isStart {
		*date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	} else {
		// make the date end at 23:59:59
		*date = time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, date.Location())
	}

	return date
}

func NewDateFilter(dateStartRaw string, dateEndRaw string) DateFilter {
	var dateStart *time.Time = nil
	var dateEnd *time.Time = nil
	var err error = nil

	if dateStartRaw != "" {
		dateStart, err = parseDate(dateStartRaw)
		if err == nil {
			dateStart = alignDate(dateStart, true)
		}
	}

	if dateEndRaw != "" {
		dateEnd, err = parseDate(dateEndRaw)
		if err == nil {
			dateEnd = alignDate(dateEnd, false)
		}
	}

	return newDateFilter(dateStart, dateEnd)
}
