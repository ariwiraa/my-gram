// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/ariwiraa/my-gram/config"
	"github.com/ariwiraa/my-gram/handler"
	"github.com/ariwiraa/my-gram/repository"
	"github.com/ariwiraa/my-gram/repository/impl"
	"github.com/ariwiraa/my-gram/routes"
	"github.com/ariwiraa/my-gram/usecase"
	impl2 "github.com/ariwiraa/my-gram/usecase/impl"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

// Injectors from injector.go:

func initializedLikesHandler() handler.UserLikesPhotosHandler {
	db := config.InitializeDB()
	userLikesPhotoRepository := impl.NewUserLikesPhotoRepository(db)
	photoRepository := impl.NewPhotoRepository(db)
	userRepository := repository.NewUserRepository(db)
	userLikesPhotosUsecase := impl2.NewUserLikesPhotosUsecase(userLikesPhotoRepository, photoRepository, userRepository)
	validate := validator.New()
	userLikesPhotosHandler := handler.NewUserLikesPhotosHandler(userLikesPhotosUsecase, validate)
	return userLikesPhotosHandler
}

func initializedUserHandler() handler.UserHandler {
	db := config.InitializeDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	authenticationRepository := impl.NewAuthenticationRepositoryImpl(db)
	authenticationUsecase := impl2.NewAuthenticationUsecaseImpl(authenticationRepository)
	validate := validator.New()
	userHandler := handler.NewUserHandler(userUsecase, authenticationUsecase, validate)
	return userHandler
}

func initializedPhotoHandler() handler.PhotoHandler {
	db := config.InitializeDB()
	photoRepository := impl.NewPhotoRepository(db)
	commentRepository := impl.NewCommentRepository(db)
	tagRepository := impl.NewTagRepositoryImpl(db)
	photoTagsRepository := impl.NewPhotoTagsRepositoryImpl(db)
	photoUsecase := impl2.NewPhotoUsecase(photoRepository, commentRepository, tagRepository, photoTagsRepository)
	validate := validator.New()
	photoHandler := handler.NewPhotoHandler(photoUsecase, validate)
	return photoHandler
}

func initializedCommentHandler() handler.CommentHandler {
	db := config.InitializeDB()
	commentRepository := impl.NewCommentRepository(db)
	photoRepository := impl.NewPhotoRepository(db)
	commentUsecase := impl2.NewCommentUsecase(commentRepository, photoRepository)
	validate := validator.New()
	commentHandler := handler.NewCommentHandler(commentUsecase, validate)
	return commentHandler
}

func initializedFollowHandler() handler.FollowHandler {
	db := config.InitializeDB()
	followRepository := impl.NewFollowRepositoryImpl(db)
	userRepository := repository.NewUserRepository(db)
	followUsecase := impl2.NewFollowUsecaseImpl(followRepository, userRepository)
	followHandler := handler.NewFollowHandlerImpl(followUsecase)
	return followHandler
}

func InitializedServer() *gin.Engine {
	userHandler := initializedUserHandler()
	photoHandler := initializedPhotoHandler()
	commentHandler := initializedCommentHandler()
	userLikesPhotosHandler := initializedLikesHandler()
	followHandler := initializedFollowHandler()
	engine := routes.NewRouter(userHandler, photoHandler, commentHandler, userLikesPhotosHandler, followHandler)
	return engine
}

// injector.go:

var followsSet = wire.NewSet(impl.NewFollowRepositoryImpl, repository.NewUserRepository, impl2.NewFollowUsecaseImpl, handler.NewFollowHandlerImpl)

var authenticationSet = wire.NewSet(impl.NewAuthenticationRepositoryImpl, impl2.NewAuthenticationUsecaseImpl)

var userSet = wire.NewSet(repository.NewUserRepository, usecase.NewUserUsecase, authenticationSet, handler.NewUserHandler)

var photoSet = wire.NewSet(impl.NewPhotoRepository, impl.NewCommentRepository, impl.NewTagRepositoryImpl, impl.NewPhotoTagsRepositoryImpl, impl2.NewPhotoUsecase, handler.NewPhotoHandler)

var commentSet = wire.NewSet(impl.NewCommentRepository, impl.NewPhotoRepository, impl2.NewCommentUsecase, handler.NewCommentHandler)

var likesSet = wire.NewSet(impl.NewUserLikesPhotoRepository, impl.NewPhotoRepository, repository.NewUserRepository, impl2.NewUserLikesPhotosUsecase, handler.NewUserLikesPhotosHandler)
