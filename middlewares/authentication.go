package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/moneybackward/backend/utils/token"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		u_id, err := token.ExtractTokenID(c)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		c.Set("userId", u_id)
		c.Next()
	}
}
