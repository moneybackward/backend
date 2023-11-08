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
	List(*gin.Context)
	Detail(*gin.Context)
	Add(*gin.Context)
	Delete(*gin.Context)
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
// @Security BearerAuth
// @Router /notes [post]
// @Param note body dto.NoteCreateDTO true "Note"
// @Success 201 {object} models.Note
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

// @Summary Get a note
// @Tags notes
// @Security BearerAuth
// @Param note_id path string true "Note ID"
// @Router /notes/{note_id} [get]
// @Success 200 {object} dto.NoteDTO
func (noteCtrl *noteController) Detail(ctx *gin.Context) {
	noteId := uuid.MustParse(ctx.Param("note_id"))
	userIdRaw, exists := ctx.Get("userId")
	userId := userIdRaw.(uuid.UUID)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	note, err := noteCtrl.noteService.Find(noteId)
	if err != nil || note.UserId != userId {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": note})
}

// @Summary List notes
// @Tags notes
// @Security BearerAuth
// @Router /notes [get]
// @Success 200 {object} []dto.NoteDTO
func (noteCtrl *noteController) List(ctx *gin.Context) {
	userIdRaw, exists := ctx.Get("userId")
	userId := userIdRaw.(uuid.UUID)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	notes, err := noteCtrl.noteService.FindUserNotes(userId)
	if err != nil {
		log.Panic().Msg(err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{"data": notes})
}

// @Summary Delete a category
// @Tags categories
// @Security BearerAuth
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
