package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/services"
	"github.com/rs/zerolog/log"
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
// @Param note body dto.NoteCreateDTO true "Note"
// @Success 201 {object} models.Note
// @Router /notes [post]
// @Security BearerAuth
func (noteCtrl *noteController) Add(ctx *gin.Context) {
	userIdRaw, exists := ctx.Get("userId")
	userId := userIdRaw.(uuid.UUID)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	var input dto.NoteCreateDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := noteCtrl.userService.Find(userId); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "message": "User not found"})
		return
	}

	note, err := noteCtrl.noteService.Create(userId, &input)
	if err != nil {
		log.Panic().Msg(err.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{"data": note})
}

// @Summary List notes
// @Tags notes
// @Success 201 {object} []models.Note
// @Router /notes [get]
// @Security BearerAuth
func (noteCtrl *noteController) List(ctx *gin.Context) {
	userIdRaw, exists := ctx.Get("userId")
	userId := userIdRaw.(uuid.UUID)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	notes, err := noteCtrl.noteService.FindAll(userId)
	if err != nil {
		log.Panic().Msg(err.Error())
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
