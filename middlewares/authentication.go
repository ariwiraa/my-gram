package middlewares

import (
	"net/http"
	"strings"

	"github.com/ariwiraa/my-gram/helpers"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.Request.Header.Get("Authorization")
		if headerToken == "" {
			helpers.FailResponse(c, http.StatusUnauthorized, "header not provide. Please login")
			c.Abort()
			return
		}

		bearer := strings.HasPrefix(headerToken, "Bearer")
		if !bearer {
			helpers.FailResponse(c, http.StatusUnauthorized, "invalid header type")
			c.Abort()
			return
		}

		stringToken := strings.Split(headerToken, " ")[1]

		verifyToken, err := helpers.VerifyToken(stringToken)
		if err != nil {
			helpers.FailResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		// menyimpan claim dari token
		c.Set("userData", verifyToken)
		c.Next()
	}
}
