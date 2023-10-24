package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/services"
)

type TransactionController interface {
	List(ctx *gin.Context)
	Add(ctx *gin.Context)
	Delete(ctx *gin.Context)
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
// @Param transaction body dto.TransactionDTO true "transaction"
// @Success 201 {object} models.Transaction
// @Router /notes/{note_id}/transactions [post]
func (ctrl *transactionController) Add(ctx *gin.Context) {
	var input dto.TransactionDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := ctrl.transactionService.Create(&input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": transaction})
}

// @Summary List transactions
// @Tags transactions
// @Success 200 {object} []models.Transaction
// @Router /notes/{note_id}/transactions [get]
func (ctrl *transactionController) List(ctx *gin.Context) {
	noteId := uuid.MustParse(ctx.Param("noteId"))
	transactions, err := ctrl.transactionService.FindAll(noteId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": transactions})
}

// @Summary Delete a transaction
// @Tags transactions
// @Success 204 {object} nil
func (ctrl *transactionController) Delete(ctx *gin.Context) {
	transactionId := uuid.MustParse(ctx.Param("id"))
	err := ctrl.transactionService.Delete(transactionId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
