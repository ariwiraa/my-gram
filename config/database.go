package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ariwiraa/my-gram/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB() *gorm.DB {
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	dbname := os.Getenv("PGNAME")

	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting database = ", err)
	} else {
		log.Println("Successfully connected to database")
	}

	db.AutoMigrate(&domain.User{}, &domain.Photo{}, &domain.Comment{}, &domain.UserLikesPhoto{}, &domain.Authentication{})

	return db
}
