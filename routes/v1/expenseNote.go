package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/moneybackward/backend/controllers"
)

func RegisterExpenseNoteRoutes(router *gin.RouterGroup) {
	expenseNoteController := controllers.NewExpenseNoteController()
	expenseNotes := router.Group("/expense-notes")
	{
		expenseNotes.GET("", expenseNoteController.List)
		expenseNotes.POST("", expenseNoteController.Add)
	}
}
