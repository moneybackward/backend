package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/moneybackward/backend/controllers"
	"github.com/moneybackward/backend/middlewares"
)

func RegisterNoteRoutes(router *gin.RouterGroup) {
	noteController := controllers.NewNoteController()
	categoryController := controllers.NewCategoryController()
	transactionController := controllers.NewTransactionController()

	notes := router.Group("/notes")
	{
		notes.GET("", middlewares.JwtAuthMiddleware(), noteController.List)
		notes.GET("/:note_id", middlewares.JwtAuthMiddleware(), noteController.Detail)
		notes.POST("", middlewares.JwtAuthMiddleware(), noteController.Add)

		notes.GET("/:note_id/categories", middlewares.JwtAuthMiddleware(), categoryController.List)
		notes.POST("/:note_id/categories", middlewares.JwtAuthMiddleware(), categoryController.Add)
		notes.GET("/:note_id/categories/:category_id", middlewares.JwtAuthMiddleware(), categoryController.Detail)
		notes.PUT("/:note_id/categories/:category_id", middlewares.JwtAuthMiddleware(), categoryController.Update)

		notes.GET("/:note_id/transactions", middlewares.JwtAuthMiddleware(), transactionController.List)
		notes.POST("/:note_id/transactions", middlewares.JwtAuthMiddleware(), transactionController.Add)
	}
}
