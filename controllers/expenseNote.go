package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/services"
)

type ExpenseNoteController interface {
	List(ctx *gin.Context)
	Add(ctx *gin.Context)
}

type expenseNoteController struct {
	expenseNoteService services.ExpenseNoteService
}

func (ctrl *expenseNoteController) Add(ctx *gin.Context) {
	var input dto.ExpenseNoteDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expenseNote, err := ctrl.expenseNoteService.Create(&input)
	if err != nil {
		log.Panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": expenseNote})
}

func (expenseNoteCtrl *expenseNoteController) List(ctx *gin.Context) {
	expenseNotes, err := expenseNoteCtrl.expenseNoteService.FindAll()
	if err != nil {
		log.Panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": expenseNotes})
}

func NewExpenseNoteController() ExpenseNoteController {
	expenseNoteService := services.NewExpenseNoteService()

	return &expenseNoteController{
		expenseNoteService: expenseNoteService,
	}
}
