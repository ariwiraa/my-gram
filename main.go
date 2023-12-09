package main

import (
	"github.com/ariwiraa/my-gram/config"
	"github.com/ariwiraa/my-gram/handler"
	"github.com/ariwiraa/my-gram/repository"
	repositoryImpl "github.com/ariwiraa/my-gram/repository/impl"
	"github.com/ariwiraa/my-gram/routes"
	usecaseImpl "github.com/ariwiraa/my-gram/usecase/impl"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

func main() {
	cfg := config.InitializeConfig()

	db := config.InitializeDB(cfg)
	router := newApp(db)

	router.Run(":" + cfg.Server.Port)
}

func newApp(db *gorm.DB) *gin.Engine {
	validate := validator.New()

	// Repository
	commentRepository := repositoryImpl.NewCommentRepository(db)
	photoRepository := repositoryImpl.NewPhotoRepository(db)
	userRepository := repository.NewUserRepository(db)
	userLikesPhotoRepository := repositoryImpl.NewUserLikesPhotoRepository(db)
	followRepository := repositoryImpl.NewFollowRepositoryImpl(db)
	authRepository := repositoryImpl.NewAuthenticationRepositoryImpl(db)
	tagRepository := repositoryImpl.NewTagRepositoryImpl(db)
	photoTagRepository := repositoryImpl.NewPhotoTagsRepositoryImpl(db)

	// Photo Set
	photoUsecase := usecaseImpl.NewPhotoUsecase(
		photoRepository,
		commentRepository,
		tagRepository,
		photoTagRepository,
		userLikesPhotoRepository,
		userRepository,
	)
	photoHandler := handler.NewPhotoHandler(photoUsecase, validate)

	// Comment set
	commentUsecase := usecaseImpl.NewCommentUsecase(commentRepository, photoRepository)
	commentHandler := handler.NewCommentHandler(commentUsecase, validate)

	// Like Photo set
	userLikesPhotosUsecase := usecaseImpl.NewUserLikesPhotosUsecase(userLikesPhotoRepository, photoRepository, userRepository)
	userLikesPhotosHandler := handler.NewUserLikesPhotosHandler(userLikesPhotosUsecase, validate)

	// User set
	userUsecase := usecaseImpl.NewUserUsecaseImpl(userRepository, photoRepository, followRepository)
	userHandler := handler.NewUserHandlerImpl(userUsecase)

	// Auth Set
	authUsecase := usecaseImpl.NewAuthenticationUsecaseImpl(authRepository, userRepository)
	authHandler := handler.NewAuthHandler(authUsecase, validate)

	// Follow Set
	followUsecase := usecaseImpl.NewFollowUsecaseImpl(followRepository, userRepository)
	followHandler := handler.NewFollowHandlerImpl(followUsecase)

	routerHandler := routes.RouterHandler{
		UserHandler:    userHandler,
		PhotoHandler:   photoHandler,
		CommentHandler: commentHandler,
		LikesHandler:   userLikesPhotosHandler,
		AuthHandler:    authHandler,
		FollowsHandler: followHandler,
	}

	router := routes.NewRouter(routerHandler)

	return router
}
