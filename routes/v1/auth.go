package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/moneybackward/backend/controllers"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	userController := controllers.NewUserController()
	users := router.Group("/auth")
	{
		users.POST("/register", userController.Register)
		users.POST("/login", userController.Login)
	}
}
