package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

type config struct {
	jwtTokenKey   string
	jwtRefreshKey string
	tokenExpiry   time.Duration
	refreshExpiry time.Duration
}

var jwtConfig = new(config)

func init() {
	intTokenExpiry, _ := strconv.Atoi(os.Getenv("TOKENEXPIRY"))
	intRefreshExpiry, _ := strconv.Atoi(os.Getenv("REFRESHEXPIRY"))
	jwtConfig.jwtTokenKey = os.Getenv("TOKENKEY")
	if jwtConfig.jwtTokenKey == "" {
		log.Println("Error: VARIABLE_NAME is not set or empty")
	}
	jwtConfig.jwtRefreshKey = os.Getenv("REFRESHKEY")
	jwtConfig.tokenExpiry = time.Duration(intTokenExpiry) * time.Minute
	jwtConfig.refreshExpiry = time.Duration(intRefreshExpiry) * time.Minute
}

func LoadJWTConfig() *config {
	return jwtConfig
}

func (jc *config) GetTokenExpiry() time.Duration {
	return jwtConfig.tokenExpiry
}

func (jc *config) GetJWTTokenKey() string {
	return jwtConfig.jwtTokenKey
}

func (jc *config) GetRefreshExpiry() time.Duration {
	return jwtConfig.refreshExpiry
}

func (jc *config) GetJWTRefreshKey() string {
	return jwtConfig.jwtRefreshKey
}
