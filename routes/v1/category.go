package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/moneybackward/backend/controllers"
)

func RegisterCategoryRoutes(router *gin.RouterGroup) {
	categoryController := controllers.NewCategoryController()
	categories := router.Group("/categories")
	{
		categories.GET("", categoryController.List)
		categories.POST("", categoryController.Add)
	}
}
