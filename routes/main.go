package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/moneybackward/backend/routes/v1"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	api := router.Group("/api")
	{
		v1.RegisterV1Routes(api, db)
	}
}
