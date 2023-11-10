package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/services"
)

type TransactionController interface {
	List(*gin.Context)
	Add(*gin.Context)
	Detail(*gin.Context)
	Delete(*gin.Context)
}

type transactionController struct {
	transactionService services.TransactionService
}

func NewTransactionController() TransactionController {
	transactionService := services.NewTransactionService()

	return &transactionController{
		transactionService: transactionService,
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := ctrl.transactionService.Create(noteId, &input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": transaction})
}

// @Summary List transactions
// @Tags transactions
// @Security BearerAuth
// @Router /notes/{note_id}/transactions [get]
// @Param note_id path string true "Note ID"
// @Success 200 {object} []dto.TransactionDTO
func (ctrl *transactionController) List(ctx *gin.Context) {
	noteId := uuid.MustParse(ctx.Param("note_id"))
	transactions, err := ctrl.transactionService.FindAllOfNote(noteId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if transaction.NoteId != noteId {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not belongs to note"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": transaction})
}

// @Summary Delete a transaction
// @Tags transactions
// @Security BearerAuth
// @Success 204 {object} nil
func (ctrl *transactionController) Delete(ctx *gin.Context) {
	noteId := uuid.MustParse(ctx.Param("note_id"))
	transactionId := uuid.MustParse(ctx.Param("transaction_id"))

	transaction, err := ctrl.transactionService.Find(transactionId)
	if err != nil || transaction.NoteId != noteId {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.transactionService.Delete(transactionId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
