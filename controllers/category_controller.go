package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/services"
)

type CategoryController interface {
	List(ctx *gin.Context)
	Add(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type categoryController struct {
	categoryService services.CategoryService
}

func NewCategoryController() CategoryController {
	categoryService := services.NewCategoryService()

	return &categoryController{
		categoryService: categoryService,
	}
}

// @Summary Add a category
// @Tags categories
// @Accept json
// @Param category body dto.CategoryDTO true "Category"
// @Success 201 {object} models.Category
// @Router /notes/:note_id/categories [post]
func (ctrl *categoryController) Add(ctx *gin.Context) {
	var input dto.CategoryCreateDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := ctrl.categoryService.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": category})
}

// @Summary List categories
// @Tags categories
// @Success 200 {object} []models.Category
// @Router /notes/:note_id/categories [get]
// @Param note_id path string true "Note ID"
func (ctrl *categoryController) List(ctx *gin.Context) {
	noteId := uuid.MustParse(ctx.Param("note_id"))
	categories, err := ctrl.categoryService.FindAll(noteId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": categories})
}

// @Summary Delete a category
// @Tags categories
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
