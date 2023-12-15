package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/services"
	"github.com/rs/zerolog/log"
)

type CategoryController interface {
	List(*gin.Context)
	Add(*gin.Context)
	Detail(*gin.Context)
	Delete(*gin.Context)
	Update(*gin.Context)
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

// @Summary Create a category
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
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}

	noteId := uuid.MustParse(ctx.Param("note_id"))

	var input dto.CategoryCreateDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isNoteBelongsToUser := ctrl.noteService.IsBelongsToUser(noteId, userId)
	if !isNoteBelongsToUser {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Not belongs to user"})
	}

	category, err := ctrl.categoryService.Create(noteId, input)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": category})
}

// @Summary Get a category
// @Tags categories
// @Security BearerAuth
// @Router /notes/{note_id}/categories/{category_id} [get]
// @Param note_id path string true "Note ID"
// @Param category_id path string true "Category ID"
// @Success 200 {object} dto.CategoryDTO
func (ctrl *categoryController) Detail(ctx *gin.Context) {
	noteId := uuid.MustParse(ctx.Param("note_id"))
	categoryId := uuid.MustParse(ctx.Param("category_id"))

	category, err := ctrl.categoryService.Find(categoryId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if category.NoteId != noteId {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Note or category not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": category})
}

// @Summary Update a category
// @Tags categories
// @Accept json
// @Security BearerAuth
// @Router /notes/{note_id}/categories/{category_id} [put]
// @Param note_id path string true "Note ID"
// @Param category_id path string true "Category ID"
// @Param category body dto.CategoryUpdateDTO true "Category"
// @Success 200 {object} dto.CategoryDTO
func (ctrl *categoryController) Update(ctx *gin.Context) {
	userIdRaw, exists := ctx.Get("userId")
	userId := userIdRaw.(uuid.UUID)
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	var input dto.CategoryUpdateDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	noteId := uuid.MustParse(ctx.Param("note_id"))
	categoryId := uuid.MustParse(ctx.Param("category_id"))

	isNoteBelongsToUser := ctrl.noteService.IsBelongsToUser(noteId, userId)
	if !isNoteBelongsToUser {
		ctx.AbortWithStatusJSON(http.StatusForbidden, "Note does not belong to the user")
	}

	isCategoryBelongsToNote := ctrl.categoryService.IsBelongsToNote(categoryId, noteId)
	if !isCategoryBelongsToNote {
		ctx.AbortWithStatusJSON(http.StatusForbidden, "Category does not belong to the note")
	}

	category, err := ctrl.categoryService.Update(categoryId, input)
	if err != nil {
		log.Panic().Msg(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": category})
}

// @Summary List categories
// @Tags categories
// @Security BearerAuth
// @Router /notes/{note_id}/categories [get]
// @Param note_id path string true "Note ID"
// @Param is_expense query bool false "Is expense"
// @Success 200 {object} []dto.CategoryDTO
func (ctrl *categoryController) List(ctx *gin.Context) {
	userIdRaw, exists := ctx.Get("userId")
	userId := userIdRaw.(uuid.UUID)
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	var isExpense *bool = nil
	// optional is_expense param
	if ctx.Query("is_expense") != "" {
		isExpenseRaw, err := strconv.ParseBool(ctx.Query("is_expense"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		isExpense = &isExpenseRaw
	}

	noteId := uuid.MustParse(ctx.Param("note_id"))
	isNoteBelongsToUser := ctrl.noteService.IsBelongsToUser(noteId, userId)
	if !isNoteBelongsToUser {
		ctx.AbortWithStatusJSON(http.StatusForbidden, "Note does not belong to the user")
	}

	categories, err := ctrl.categoryService.FindAllOfNote(noteId, isExpense)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": categories})
}

// @Summary Delete a category
// @Tags categories
// @Security BearerAuth
// @Router /notes/{note_id}/categories/{category_id} [delete]
// @Param note_id path string true "Note ID"
// @Param category_id path string true "Category ID"
// @Success 204 {object} nil
func (ctrl *categoryController) Delete(ctx *gin.Context) {
	userIdRaw, exists := ctx.Get("userId")
	noteId := uuid.MustParse(ctx.Param("note_id"))
	categoryId := uuid.MustParse(ctx.Param("category_id"))

	userId := userIdRaw.(uuid.UUID)
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}
	note, err := ctrl.noteService.Find(noteId)
	if err != nil || note.UserId != userId {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	isCategoryBelongsToNote := ctrl.categoryService.IsBelongsToNote(categoryId, noteId)
	if !isCategoryBelongsToNote {
		ctx.AbortWithStatusJSON(http.StatusForbidden, "Category does not belong to the note")
	}

	err = ctrl.categoryService.Delete(categoryId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
