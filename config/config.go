package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Server   server
	Database database
	JWT      jwtEnvironment
}

type server struct {
	Host string
	Port string
}

func InitializeConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error when load env %s", err.Error())
	}

	return &Config{
		server{
			Host: os.Getenv("APP_HOST"),
			Port: os.Getenv("APP_PORT"),
		},
		database{
			Host:     os.Getenv("PG_HOST"),
			Port:     os.Getenv("PG_PORT"),
			Username: os.Getenv("PG_USERNAME"),
			Password: os.Getenv("PG_PASSWORD"),
			Name:     os.Getenv("PG_NAME"),
		},
		jwtEnvironment{
			JWTTokenKey:   os.Getenv("TOKEN_KEY"),
			JWTRefreshKey: os.Getenv("REFRESH_KEY"),
			TokenExpiry:   os.Getenv("TOKEN_EXPIRY"),
			RefreshExpiry: os.Getenv("REFRESH_EXPIRY"),
		},
	}

}
