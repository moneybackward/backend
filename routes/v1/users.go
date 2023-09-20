package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/moneybackward/backend/controllers"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	userController := controllers.NewUserController()
	users := router.Group("/users")
	{
		users.GET("", userController.ListUsers)
		users.POST("", userController.AddUser)
	}
}
