package v1

import (
	"github.com/gin-gonic/gin"
)

func RegisterV1Routes(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	{
		RegisterUserRoutes(v1)
	}
}
