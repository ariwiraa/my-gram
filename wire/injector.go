//go:build wireinject
// +build wireinject

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

var userSet = wire.NewSet(
	repository.NewUserRepository, usecase.NewUserUsecase, handler.NewUserHandler,
)

var photoSet = wire.NewSet(
	repository.NewPhotoRepository, repository.NewCommentRepository, usecase.NewPhotoUsecase, handler.NewPhotoHandler,
)
var commentSet = wire.NewSet(
	repository.NewCommentRepository, repository.NewPhotoRepository, usecase.NewCommentUsecase, handler.NewCommentHandler,
)

func initializedUserHandler() handler.UserHandler {
	wire.Build(config.InitializeDB, validator.New, userSet)
	return nil
}

func initializedPhotoHandler() handler.PhotoHandler {
	wire.Build(config.InitializeDB, validator.New, photoSet)
	return nil
}

func initializedCommentHandler() handler.CommentHandler {
	wire.Build(config.InitializeDB, validator.New, commentSet)
	return nil
}

func InitializedServer() *gin.Engine {
	wire.Build(initializedUserHandler, initializedPhotoHandler, initializedCommentHandler, routes.NewRouter)
	return nil
}
