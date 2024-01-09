package middlewares

import (
	"log"
	"strings"

	"github.com/ariwiraa/my-gram/helpers"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.Request.Header.Get("Authorization")
		if headerToken == "" {
			log.Printf("[Authentication, Get] header token is nil")

			helpers.NewResponse(
				helpers.WithMessage(helpers.ErrHeaderNotProvide.Error()),
				helpers.WithError(helpers.ErrHeaderNotProvide),
			).Send(c)

			c.Abort()
			return
		}

		bearer := strings.HasPrefix(headerToken, "Bearer")
		if !bearer {
			log.Printf("[Authentication, HasPrefix] invalid header type")

			helpers.NewResponse(
				helpers.WithMessage(helpers.ErrInvalidHeaderType.Error()),
				helpers.WithError(helpers.ErrInvalidHeaderType),
			).Send(c)

			c.Abort()
			return
		}

		stringToken := strings.Split(headerToken, " ")[1]

		verifyToken, err := helpers.VerifyToken(stringToken)
		if err != nil {
			log.Printf("[Authentication, VerifyToken] with error detail %v", err.Error())

			helpers.NewResponse(
				helpers.WithMessage(helpers.ErrTokenNotVerified.Error()),
				helpers.WithError(helpers.ErrTokenNotVerified),
			).Send(c)

			return
		}

		// menyimpan claim dari token
		c.Set("userData", verifyToken)
		c.Next()
	}
}
