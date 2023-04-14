package middlewares

import (
	"net/http"

	"github.com/ariwiraa/my-gram/helpers"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}
		// menyimpan claim dari token
		c.Set("userData", verifyToken)
		c.Next()
	}
}
