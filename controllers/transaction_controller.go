package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/services"
	"github.com/moneybackward/backend/utils"
	"github.com/rs/zerolog/log"
)

type TransactionController interface {
	List(*gin.Context)
	Add(*gin.Context)
	Detail(*gin.Context)
	Delete(*gin.Context)
	Update(*gin.Context)
}

type transactionController struct {
	transactionService services.TransactionService
	noteService        services.NoteService
}

func NewTransactionController() TransactionController {
	transactionService := services.NewTransactionService()
	noteService := services.NewNoteService()

	return &transactionController{
		transactionService: transactionService,
		noteService:        noteService,
	}
}

// @Summary Add a transaction
// @Tags transactions
// @Accept json
// @Security BearerAuth
// @Router /notes/{note_id}/transactions [post]
// @Param note_id path string true "Note ID"
// @Param transaction body dto.TransactionCreateDTO true "transaction"
// @Success 201 {object} dto.TransactionDTO
func (ctrl *transactionController) Add(ctx *gin.Context) {
	noteId := uuid.MustParse(ctx.Param("note_id"))

	var input dto.TransactionCreateDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := ctrl.transactionService.Create(noteId, &input)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": transaction})
}

// @Summary List transactions
// @Tags transactions
// @Security BearerAuth
// @Router /notes/{note_id}/transactions [get]
// @Param note_id path string true "Note ID"
// @Param is_expense query bool false "Is expense"
// @Param date_start query string false "Date start"
// @Param date_end query string false "Date end"
// @Success 200 {object} []dto.TransactionDTO
func (ctrl *transactionController) List(ctx *gin.Context) {
	// optional is_expense param
	var isExpense *bool = nil
	if ctx.Query("is_expense") != "" {
		isExpenseRaw, err := strconv.ParseBool(ctx.Query("is_expense"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		isExpense = &isExpenseRaw
	}

	dateFilter := utils.NewDateFilter(ctx.Query("date_start"), ctx.Query("date_end"))
	noteId := uuid.MustParse(ctx.Param("note_id"))
	transactions, err := ctrl.transactionService.FindAllOfNote(noteId, isExpense, &dateFilter)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": transactions})
}

// @Summary Get transaction detail
// @Tags transactions
// @Security BearerAuth
// @Router /notes/{note_id}/transactions/{transaction_id} [get]
// @Param note_id path string true "Note ID"
// @Param transaction_id path string true "Transaction ID"
// @Success 200 {object} []dto.TransactionDTO
func (ctrl *transactionController) Detail(ctx *gin.Context) {
	noteId := uuid.MustParse(ctx.Param("note_id"))

	transactionId := uuid.MustParse(ctx.Param("transaction_id"))
	transaction, err := ctrl.transactionService.Find(transactionId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if transaction.NoteId != noteId {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not belongs to note"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": transaction})
}

// @Summary Set budget for a transaction
// @Tags transactions
// @Accept json
// @Security BearerAuth
// @Router /notes/{note_id}/transactions/{transaction_id} [put]
// @Param note_id path string true "Note ID"
// @Param transaction_id path string true "Transaction ID"
// @Param transaction body dto.TransactionUpdateDTO true "transaction"
// @Success 200 {object} dto.TransactionDTO
func (ctrl *transactionController) Update(ctx *gin.Context) {
	userIdRaw, exists := ctx.Get("userId")
	userId := userIdRaw.(uuid.UUID)
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	var input dto.TransactionUpdateDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	noteId := uuid.MustParse(ctx.Param("note_id"))
	transactionId := uuid.MustParse(ctx.Param("transaction_id"))

	note, err := ctrl.noteService.Find(noteId)
	if err != nil || note.UserId != userId {
		log.Panic().Msg(err.Error())
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	transaction, err := ctrl.transactionService.Find(transactionId)
	if err != nil || transaction.NoteId != noteId {
		log.Panic().Msg(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	transaction, err = ctrl.transactionService.Update(transactionId, input)
	if err != nil {
		log.Panic().Msg(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": transaction})
}

// @Summary Delete a transaction
// @Tags transactions
// @Security BearerAuth
// @Router /notes/{note_id}/transactions/{transaction_id} [delete]
// @Param note_id path string true "Note ID"
// @Param transaction_id path string true "Transaction ID"
// @Success 204 {object} nil
func (ctrl *transactionController) Delete(ctx *gin.Context) {
	// user validation
	userIdRaw, exists := ctx.Get("userId")
	userId := userIdRaw.(uuid.UUID)
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	noteId := uuid.MustParse(ctx.Param("note_id"))
	transactionId := uuid.MustParse(ctx.Param("transaction_id"))

	// note - user validation
	note, err := ctrl.noteService.Find(noteId)
	if err != nil || note.UserId != userId {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// transaction - note validation
	transaction, err := ctrl.transactionService.Find(transactionId)
	if err != nil || transaction.NoteId != noteId {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// delete
	err = ctrl.transactionService.Delete(transactionId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
