package helpers

import (
	"errors"
	"github.com/ariwiraa/my-gram/config"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var myAccessToken = []byte(config.LoadJwtConfig().GetTokenKey())
var myRefreshToken = []byte(config.LoadJwtConfig().GetRefreshKey())
var accessTokenExpiry = config.LoadJwtConfig().GetTokenExpiry()
var refreshTokenExpiry = config.LoadJwtConfig().GetRefreshExpiry()

type jwtConfig struct {
	cnf *config.Config
}

type claims struct {
	Id uint64
	jwt.StandardClaims
}

func NewAccessToken(id uint64) *claims {
	return &claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenExpiry).Unix(),
		},
	}
}

func NewRefreshToken(id uint64) *claims {
	return &claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refreshTokenExpiry).Unix(),
		},
	}
}

func (c *claims) GenerateAccessToken() string {
	parsetoken := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signedToken, err := parsetoken.SignedString(myAccessToken)
	if err != nil {
		log.Fatalf("error when generate access token: %s", err.Error())
	}

	return signedToken
}

func (c *claims) GenerateRefreshToken() string {

	parsetoken := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	signedToken, err := parsetoken.SignedString(myRefreshToken)
	if err != nil {
		log.Fatal("gagal membuat refresh token")
	}

	return signedToken
}

func VerifyToken(tokenString string) (interface{}, error) {
	errResponse := errors.New("sign in to proceed")

	token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}

		return []byte(myAccessToken), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}

func VerifyRefreshToken(token string) (*claims, error) {
	verifyToken, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return myRefreshToken, nil
	})
	if err != nil {
		return nil, err
	}

	if payload, ok := verifyToken.Claims.(*claims); ok && verifyToken.Valid {
		return payload, nil
	}

	return nil, errors.New("refresh token tidak valid")
}
