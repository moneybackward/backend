package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moneybackward/backend/repositories"
	"github.com/moneybackward/backend/utils/token"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		claims, err := token.ExtractClaims(c)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		u_id, _ := uuid.Parse(claims["user_id"].(string))

		_, error := repositories.UserRepositoryInstance.Find(u_id)
		if error != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": error.Error()})
			return
		}

		c.Set("claims", claims)
		c.Set("userId", u_id)
		c.Next()
	}
}
