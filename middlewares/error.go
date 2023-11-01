package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moneybackward/backend/utils"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				switch e := err.Err.(type) {
				case *utils.ValidationError:
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": e.Error()})
					return
				default:
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
					return
				}
			}
		}
	}
}
