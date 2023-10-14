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

func (ctrl *categoryController) Add(ctx *gin.Context) {
	var input dto.CategoryDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := ctrl.categoryService.Create(&input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": category})
}

func (ctrl *categoryController) List(ctx *gin.Context) {
	categories, err := ctrl.categoryService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": categories})
}

func (ctrl *categoryController) Delete(ctx *gin.Context) {
	categoryId := uuid.MustParse(ctx.Param("id"))
	err := ctrl.categoryService.Delete(categoryId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
