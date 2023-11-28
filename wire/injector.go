//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/ariwiraa/my-gram/config"
	"github.com/ariwiraa/my-gram/handler"
	"github.com/ariwiraa/my-gram/repository"
	repositoryImpl "github.com/ariwiraa/my-gram/repository/impl"
	"github.com/ariwiraa/my-gram/routes"
	"github.com/ariwiraa/my-gram/usecase"
	usecaseImpl "github.com/ariwiraa/my-gram/usecase/impl"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

var authenticationSet = wire.NewSet(
	repositoryImpl.NewAuthenticationRepositoryImpl, usecaseImpl.NewAuthenticationUsecaseImpl,
)

var userSet = wire.NewSet(
	repository.NewUserRepository, usecase.NewUserUsecase, authenticationSet, handler.NewUserHandler,
)

var photoSet = wire.NewSet(
	repository.NewPhotoRepository,
	repository.NewCommentRepository,
	repositoryImpl.NewTagRepositoryImpl,
	repositoryImpl.NewPhotoTagsRepositoryImpl,
	usecase.NewPhotoUsecase,
	handler.NewPhotoHandler,
)
var commentSet = wire.NewSet(
	repository.NewCommentRepository, repository.NewPhotoRepository, usecase.NewCommentUsecase, handler.NewCommentHandler,
)

var likesSet = wire.NewSet(repository.NewUserLikesPhotoRepository, repository.NewPhotoRepository, usecase.NewUserLikesPhotosUsecase, handler.NewUserLikesPhotosHandler)

func initializedLikesHandler() handler.UserLikesPhotosHandler {
	wire.Build(config.InitializeDB, validator.New, likesSet)
	return nil
}

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
	wire.Build(initializedUserHandler, initializedPhotoHandler, initializedCommentHandler, initializedLikesHandler, routes.NewRouter)
	return nil
}
