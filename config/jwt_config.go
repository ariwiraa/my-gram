package config

import (
	"fmt"
	"strconv"
	"time"
)

type jwtEnvironment struct {
	JWTTokenKey   string
	JWTRefreshKey string
	TokenExpiry   string
	RefreshExpiry string
}

type jwtConfig struct {
	cfg Config
}

func LoadJwtConfig() *jwtConfig {
	cnf := InitializeConfig()
	return &jwtConfig{
		cfg: *cnf,
	}
}
func (c *jwtConfig) GetTokenKey() string {
	return c.cfg.JWT.JWTTokenKey
}

func (c *jwtConfig) GetRefreshKey() string {
	return c.cfg.JWT.JWTRefreshKey
}

func (c *jwtConfig) GetTokenExpiry() time.Duration {
	convertTokenExpiryToInt, err := strconv.Atoi(c.cfg.JWT.TokenExpiry)
	if err != nil {
		fmt.Errorf("error converting token expiry %s ", err.Error())
	}
	tokenExpiry := time.Duration(convertTokenExpiryToInt) * time.Minute
	return tokenExpiry
}

func (c *jwtConfig) GetRefreshExpiry() time.Duration {
	convertRefreshExpiryToInt, err := strconv.Atoi(c.cfg.JWT.RefreshExpiry)
	if err != nil {
		fmt.Errorf("error converting refresh expiry %s ", err.Error())
	}
	refreshExpiry := time.Duration(convertRefreshExpiryToInt) * time.Minute
	return refreshExpiry
}
