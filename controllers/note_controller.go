package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/services"
)

type NoteController interface {
	List(ctx *gin.Context)
	Add(ctx *gin.Context)
	Delete(ctx *gin.Context)
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

// @Summary Add a note
// @Tags notes
// @Accept json
// @Param note body dto.NoteDTO true "Note"
// @Success 201 {object} models.Note
// @Router /notes [post]
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

// @Summary List notes
// @Tags notes
// @Success 201 {object} []models.Note
// @Router /notes [get]
// @Param userId query string true "User ID"
func (noteCtrl *noteController) List(ctx *gin.Context) {
	var userId uuid.UUID
	if ctx.Query("userId") != "" {
		userId = uuid.MustParse(ctx.Query("userId"))
	}

	notes, err := noteCtrl.noteService.FindAll(userId)
	if err != nil {
		log.Panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": notes})
}

// @Summary Delete a category
// @Tags categories
// @Success 204 {object} nil
func (noteCtrl *noteController) Delete(ctx *gin.Context) {
	noteId := uuid.MustParse(ctx.Param("id"))
	err := noteCtrl.noteService.Delete(noteId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
