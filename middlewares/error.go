package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moneybackward/backend/utils/errors"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				switch e := err.Err.(type) {
				case *errors.ValidationError:
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": e.Error()})
					return

				case *errors.ConflictError:
					c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": e.Error()})
					return

				case *errors.NotFoundError:
					c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": e.Error()})
					return

				case *errors.UnauthorizedError:
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": e.Error()})
					return

				default:
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
					return
				}
			}
		}
	}
}
