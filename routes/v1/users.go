package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/moneybackward/backend/controllers"
	"github.com/moneybackward/backend/repositories"
	"github.com/moneybackward/backend/services"
	"gorm.io/gorm"
)

func RegisterUserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	users := router.Group("/users")
	{
		users.GET("", userController.ListUsers)
		users.POST("", userController.AddUser)
	}
}
