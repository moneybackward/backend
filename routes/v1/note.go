package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/moneybackward/backend/controllers"
)

func RegisterNoteRoutes(router *gin.RouterGroup) {
	noteController := controllers.NewNoteController()
	categoryController := controllers.NewCategoryController()
	notes := router.Group("/notes")
	{
		notes.GET("", noteController.List)
		notes.POST("", noteController.Add)

		notes.GET("/:noteId/categories", categoryController.List)
		notes.POST("/:noteId/categories", categoryController.Add)
	}
}
