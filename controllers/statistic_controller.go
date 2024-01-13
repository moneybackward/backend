package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moneybackward/backend/services"
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
	// optional is_expense param
	if ctx.Query("is_expense") != "" {
		isExpenseRaw, err := strconv.ParseBool(ctx.Query("is_expense"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		isExpense = &isExpenseRaw
	}

	categoriesStats, err := ctrl.statisticService.Categories(noteId, isExpense)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get categories statistic")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": categoriesStats})
}
