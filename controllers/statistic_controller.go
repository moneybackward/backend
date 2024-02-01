package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moneybackward/backend/services"
	"github.com/moneybackward/backend/utils"
	"github.com/rs/zerolog/log"
)

type StatisticController interface {
	Categories(*gin.Context)
}

type statisticController struct {
	noteService      services.NoteService
	statisticService services.StatisticService
}

func NewStatisticController() StatisticController {
	noteService := services.NewNoteService()
	statisticService := services.NewStatisticService()

	return &statisticController{
		noteService:      noteService,
		statisticService: statisticService,
	}
}

// @Summary Get categories statistic
// @Tags statistics
// @Accept json
// @Security BearerAuth
// @Router /notes/{note_id}/statistics/categories [get]
// @Param note_id path string true "Note ID"
// @Param is_expense query bool false "Is expense"
// @Param date_start query string false "Date start"
// @Param date_end query string false "Date end"
// @Success 200 {object} []dto.CategoryStatsDTO
func (ctrl *statisticController) Categories(ctx *gin.Context) {
	userIdRaw, exists := ctx.Get("userId")
	userId := userIdRaw.(uuid.UUID)
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	noteId := uuid.MustParse(ctx.Param("note_id"))
	isNoteBelongsToUser := ctrl.noteService.IsBelongsToUser(noteId, userId)
	if !isNoteBelongsToUser {
		ctx.AbortWithStatusJSON(http.StatusForbidden, "Note does not belong to the user")
	}

	var isExpense *bool = nil
	var dateFilter utils.DateFilter = utils.NewDateFilter(nil, nil)

	// optional is_expense param
	if ctx.Query("is_expense") != "" {
		isExpenseRaw, err := strconv.ParseBool(ctx.Query("is_expense"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		isExpense = &isExpenseRaw
	}

	if dateStartRaw := ctx.Query("date_start"); dateStartRaw != "" {
		layout := "2006-01-02"
		dateStart, err := time.Parse(layout, dateStartRaw)
		// make the date start at 00:00:00
		dateStart = time.Date(dateStart.Year(), dateStart.Month(), dateStart.Day(), 0, 0, 0, 0, dateStart.Location())
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		dateFilter.Start = &dateStart
	}

	if dateEndRaw := ctx.Query("date_end"); dateEndRaw != "" {
		layout := "2006-01-02"
		dateEnd, err := time.Parse(layout, dateEndRaw)
		// make the date end at 23:59:59
		dateEnd = time.Date(dateEnd.Year(), dateEnd.Month(), dateEnd.Day(), 23, 59, 59, 0, dateEnd.Location())
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		dateFilter.End = &dateEnd
	}

	categoriesStats, err := ctrl.statisticService.Categories(noteId, isExpense, &dateFilter)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get categories statistic")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": categoriesStats})
}
