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
	userService services.UserService
}

func NewNoteController() NoteController {
	return &noteController{
		noteService: services.NewNoteService(),
		userService: services.NewUserService(),
	}
}

func (noteCtrl *noteController) Add(ctx *gin.Context) {
	var input dto.NoteDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := noteCtrl.userService.Find(input.UserId); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "message": "User not found"})
		return
	}

	note, err := noteCtrl.noteService.Create(&input)
	if err != nil {
		log.Panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": note})
}

func (noteCtrl *noteController) List(ctx *gin.Context) {
	notes, err := noteCtrl.noteService.FindAll()
	if err != nil {
		log.Panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": notes})
}
