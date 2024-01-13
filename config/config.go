package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server     server
	Database   database
	JWT        jwtEnvironment
	Redis      RedisConfig
	Cloudinary CloudinaryConfig
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
		RedisConfig{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
		CloudinaryConfig{
			Name:      os.Getenv("CLOUDINARY_NAME"),
			APIKey:    os.Getenv("CLOUDINARY_API_KEY"),
			APISecret: os.Getenv("CLOUDINARY_API_SECRET"),
		},
	}

}
