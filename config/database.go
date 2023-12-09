package config

import (
	"fmt"
	"github.com/ariwiraa/my-gram/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type database struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

func InitializeDB(cnf *Config) *gorm.DB {

	psqlInfo := fmt.Sprintf(
		"host=%s "+
			"user=%s "+
			"password=%s "+
			"dbname=%s "+
			"port=%v "+
			"sslmode=disable",
		cnf.Database.Host,
		cnf.Database.Username,
		cnf.Database.Password,
		cnf.Database.Name,
		cnf.Database.Port,
	)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting database = ", err)
	}

	db.AutoMigrate(
		&domain.User{},
		&domain.Photo{},
		&domain.Comment{},
		&domain.UserLikesPhoto{},
		&domain.Authentication{},
		&domain.Tag{},
		&domain.Follow{},
	)

	return db
}
