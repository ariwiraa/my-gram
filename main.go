package main

import (
	"log"

	"github.com/ariwiraa/my-gram/config"
	"github.com/ariwiraa/my-gram/handler"
	"github.com/ariwiraa/my-gram/repository"
	"github.com/ariwiraa/my-gram/routes"
	"github.com/ariwiraa/my-gram/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	// port := os.Getenv("APP_PORT")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("gagal mengambil .env %v", err)
	}

	db := config.InitializeDB()
	validate := validator.New()
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase, validate)

	photoRepository := repository.NewPhotoRepository(db)
	photoUsecase := usecase.NewPhotoUsecase(photoRepository)
	photoHandler := handler.NewPhotoHandler(photoUsecase, validate)

	socialMediaRepository := repository.NewSocialMediaRepository(db)
	socialMediaUsecase := usecase.NewSocialMediaUsecase(socialMediaRepository)
	socialMediaHandler := handler.NewSocialMediaHandler(socialMediaUsecase)

	commentRepository := repository.NewCommentRepository(db)
	commentUsecase := usecase.NewCommentUsecase(commentRepository, photoRepository)
	commentHandler := handler.NewCommentHandler(commentUsecase)

	router := routes.NewRouter(userHandler, photoHandler, socialMediaHandler, commentHandler)

	router.Run(":8080")
}
