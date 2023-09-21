package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/services"
)

type NoteController interface {
	List(ctx *gin.Context)
	Add(ctx *gin.Context)
}

type noteController struct {
	noteService services.NoteService
}

func (noteCtrl *noteController) Add(ctx *gin.Context) {
	var input dto.NoteDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Note, err := noteCtrl.noteService.Create(&input)
	if err != nil {
		log.Panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": Note})
}

func (noteCtrl *noteController) List(ctx *gin.Context) {
	Notes, err := noteCtrl.noteService.FindAll()
	if err != nil {
		log.Panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": Notes})
}

func NewNoteController() NoteController {
	return &noteController{
		noteService: services.NewNoteService(),
	}
}
