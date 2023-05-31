// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/ariwiraa/my-gram/config"
	"github.com/ariwiraa/my-gram/handler"
	"github.com/ariwiraa/my-gram/repository"
	"github.com/ariwiraa/my-gram/routes"
	"github.com/ariwiraa/my-gram/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

// Injectors from injector.go:

func initializedUserHandler() handler.UserHandler {
	db := config.InitializeDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	validate := validator.New()
	userHandler := handler.NewUserHandler(userUsecase, validate)
	return userHandler
}

func initializedPhotoHandler() handler.PhotoHandler {
	db := config.InitializeDB()
	photoRepository := repository.NewPhotoRepository(db)
	commentRepository := repository.NewCommentRepository(db)
	photoUsecase := usecase.NewPhotoUsecase(photoRepository, commentRepository)
	validate := validator.New()
	photoHandler := handler.NewPhotoHandler(photoUsecase, validate)
	return photoHandler
}

func initializedCommentHandler() handler.CommentHandler {
	db := config.InitializeDB()
	commentRepository := repository.NewCommentRepository(db)
	photoRepository := repository.NewPhotoRepository(db)
	commentUsecase := usecase.NewCommentUsecase(commentRepository, photoRepository)
	validate := validator.New()
	commentHandler := handler.NewCommentHandler(commentUsecase, validate)
	return commentHandler
}

func InitializedServer() *gin.Engine {
	userHandler := initializedUserHandler()
	photoHandler := initializedPhotoHandler()
	commentHandler := initializedCommentHandler()
	engine := routes.NewRouter(userHandler, photoHandler, commentHandler)
	return engine
}

// injector.go:

var userSet = wire.NewSet(repository.NewUserRepository, usecase.NewUserUsecase, handler.NewUserHandler)

var photoSet = wire.NewSet(repository.NewPhotoRepository, repository.NewCommentRepository, usecase.NewPhotoUsecase, handler.NewPhotoHandler)

var commentSet = wire.NewSet(repository.NewCommentRepository, repository.NewPhotoRepository, usecase.NewCommentUsecase, handler.NewCommentHandler)
