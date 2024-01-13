package main

import (
	"github.com/ariwiraa/my-gram/config"
	"github.com/ariwiraa/my-gram/handler"
	"github.com/ariwiraa/my-gram/repository"
	repositoryImpl "github.com/ariwiraa/my-gram/repository/impl"
	"github.com/ariwiraa/my-gram/routes"
	usecaseImpl "github.com/ariwiraa/my-gram/usecase/impl"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func main() {
	cfg := config.InitializeConfig()

	db := config.InitializeDB(cfg)

	redis, err := config.ConnectRedis(cfg)
	if err != nil {
		panic(err)
	}

	cloudinary, err := config.ConnectCloudinary(cfg)
	if err != nil {
		panic(err)
	}

	router := newApp(db, redis, cloudinary)

	router.Run(":" + cfg.Server.Port)
}

func newApp(db *gorm.DB, client *redis.Client, cloudinary *cloudinary.Cloudinary) *gin.Engine {
	validate := validator.New()

	// Repository
	redisRepository := repositoryImpl.NewRedisRepositoryImpl(client)
	commentRepository := repositoryImpl.NewCommentRepository(db)
	photoRepository := repositoryImpl.NewPhotoRepository(db)
	userRepository := repository.NewUserRepository(db)
	userLikesPhotoRepository := repositoryImpl.NewUserLikesPhotoRepository(db)
	followRepository := repositoryImpl.NewFollowRepositoryImpl(db)
	authRepository := repositoryImpl.NewAuthenticationRepositoryImpl(db)
	tagRepository := repositoryImpl.NewTagRepositoryImpl(db)
	photoTagRepository := repositoryImpl.NewPhotoTagsRepositoryImpl(db)

	// Upload
	cloudinaryUsecase := usecaseImpl.NewCloudinaryImpl(*cloudinary)
	uploadFileUsecase := usecaseImpl.NewUploadFileImpl(cloudinaryUsecase)
	uploadFileHandler := handler.NewUploadFileHandler(uploadFileUsecase)

	// Photo Set
	photoUsecase := usecaseImpl.NewPhotoUsecase(
		photoRepository,
		commentRepository,
		tagRepository,
		photoTagRepository,
		userLikesPhotoRepository,
		userRepository,
		cloudinaryUsecase,
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
	authUsecase := usecaseImpl.NewAuthenticationUsecaseImpl(authRepository, userRepository, redisRepository)
	authHandler := handler.NewAuthHandler(authUsecase, validate)

	// Follow Set
	followUsecase := usecaseImpl.NewFollowUsecaseImpl(followRepository, userRepository)
	followHandler := handler.NewFollowHandlerImpl(followUsecase)

	routerHandler := routes.RouterHandler{
		UserHandler:       userHandler,
		PhotoHandler:      photoHandler,
		CommentHandler:    commentHandler,
		LikesHandler:      userLikesPhotosHandler,
		AuthHandler:       authHandler,
		FollowsHandler:    followHandler,
		UploadFileHandler: *uploadFileHandler,
	}

	router := routes.NewRouter(routerHandler)

	return router
}
