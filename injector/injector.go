//go:build wireinject
// +build wireinject

package injector

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
	repository.NewPhotoRepository, usecase.NewPhotoUsecase, handler.NewPhotoHandler,
)

func initializedUserHandler() handler.UserHandler {
	wire.Build(config.InitializeDB, validator.New, userSet)
	return nil
}

func initializedPhotoHandler() handler.PhotoHandler {
	wire.Build(config.InitializeDB, validator.New, photoSet)
	return nil
}

func InitializedServer() *gin.Engine {
	wire.Build(initializedUserHandler, initializedPhotoHandler, routes.NewRouter)
	return nil
}
