package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/moneybackward/backend/controllers"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	users := router.Group("/users")
	{
		users.GET("", controllers.ListUsers)
		users.POST("", controllers.CreateUser)
	}
}
