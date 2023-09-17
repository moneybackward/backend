package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/moneybackward/backend/routes/v1"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		v1.RegisterV1Routes(api)
	}
}
