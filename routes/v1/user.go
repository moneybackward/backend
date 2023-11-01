package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/moneybackward/backend/controllers"
	"github.com/moneybackward/backend/middlewares"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	userController := controllers.NewUserController()
	users := router.Group("/users")
	{
		users.GET("", middlewares.JwtAuthMiddleware(), userController.List)
	}
}
