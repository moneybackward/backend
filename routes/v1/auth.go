package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/moneybackward/backend/controllers"
	"github.com/moneybackward/backend/middlewares"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	userController := controllers.NewUserController()
	users := router.Group("/auth")
	{
		users.POST("/register", userController.Register)
		users.POST("/login", userController.Login)
		users.GET("/verify", middlewares.JwtAuthMiddleware(), userController.VerifyToken)
	}
}
