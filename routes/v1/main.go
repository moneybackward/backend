package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterV1Routes(router *gin.RouterGroup, db *gorm.DB) {
	v1 := router.Group("/v1")
	{
		RegisterUserRoutes(v1, db)
	}
}
