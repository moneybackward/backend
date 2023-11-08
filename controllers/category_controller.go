package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/services"
	"github.com/rs/zerolog/log"
)

type CategoryController interface {
	List(*gin.Context)
	Add(*gin.Context)
	Delete(*gin.Context)
	SetBudget(*gin.Context)
}

type categoryController struct {
	categoryService services.CategoryService
	noteService     services.NoteService
}

func NewCategoryController() CategoryController {
	categoryService := services.NewCategoryService()
	noteService := services.NewNoteService()

	return &categoryController{
		categoryService: categoryService,
		noteService:     noteService,
	}
}

// @Summary Add a category
// @Tags categories
// @Accept json
// @Security BearerAuth
// @Router /notes/{note_id}/categories [post]
// @Param note_id path string true "Note ID"
// @Param category body dto.CategoryCreateDTO true "Category"
// @Success 201 {object} dto.CategoryDTO
func (ctrl *categoryController) Add(ctx *gin.Context) {
	userIdRaw, exists := ctx.Get("userId")
	userId := userIdRaw.(uuid.UUID)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}

	noteId := uuid.MustParse(ctx.Param("note_id"))

	var input dto.CategoryCreateDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isNoteBelongsToUser := ctrl.noteService.IsBelongsToUser(noteId, userId)
	if !isNoteBelongsToUser {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not belongs to user"})
	}

	category, err := ctrl.categoryService.Create(noteId, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": category})
}

// @Summary Set budget for a category
// @Tags categories
// @Accept json
// @Security BearerAuth
// @Router /notes/{note_id}/categories/{category_id}/budget [post]
// @Param note_id path string true "Note ID"
// @Param category_id path string true "Category ID"
// @Param category body dto.CategorySetBudgetDTO true "Category"
// @Success 200 {object} dto.CategoryDTO
func (ctrl *categoryController) SetBudget(ctx *gin.Context) {
	var input dto.CategorySetBudgetDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	noteId := uuid.MustParse(ctx.Param("note_id"))
	categoryId := uuid.MustParse(ctx.Param("category_id"))

	userIdRaw, exists := ctx.Get("userId")
	userId := userIdRaw.(uuid.UUID)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	isNoteBelongsToUser := ctrl.noteService.IsBelongsToUser(noteId, userId)
	if !isNoteBelongsToUser {
		ctx.AbortWithStatusJSON(http.StatusForbidden, "Note does not belong to the user")
	}

	isCategoryBelongsToNote := ctrl.categoryService.IsBelongsToNote(categoryId, noteId)
	if !isCategoryBelongsToNote {
		ctx.AbortWithStatusJSON(http.StatusForbidden, "Category does not belong to the note")
	}

	category, err := ctrl.categoryService.UpdateBudget(categoryId, input.Budget)
	if err != nil {
		log.Panic().Msg(err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{"data": category})
}

// @Summary List categories
// @Tags categories
// @Security BearerAuth
// @Router /notes/{note_id}/categories [get]
// @Param note_id path string true "Note ID"
// @Success 200 {object} []dto.CategoryDTO
func (ctrl *categoryController) List(ctx *gin.Context) {
	noteId := uuid.MustParse(ctx.Param("note_id"))
	categories, err := ctrl.categoryService.FindAllOfNote(noteId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": categories})
}

// @Summary Delete a category
// @Tags categories
// @Security BearerAuth
// @Success 204 {object} nil
func (ctrl *categoryController) Delete(ctx *gin.Context) {
	categoryId := uuid.MustParse(ctx.Param("id"))
	err := ctrl.categoryService.Delete(categoryId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
