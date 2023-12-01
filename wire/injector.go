//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/ariwiraa/my-gram/config"
	"github.com/ariwiraa/my-gram/handler"
	"github.com/ariwiraa/my-gram/repository"
	repositoryImpl "github.com/ariwiraa/my-gram/repository/impl"
	"github.com/ariwiraa/my-gram/routes"
	usecaseImpl "github.com/ariwiraa/my-gram/usecase/impl"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

var followsSet = wire.NewSet(
	repositoryImpl.NewFollowRepositoryImpl,
	repository.NewUserRepository,
	usecaseImpl.NewFollowUsecaseImpl,
	handler.NewFollowHandlerImpl,
)

var authenticationSet = wire.NewSet(
	repositoryImpl.NewAuthenticationRepositoryImpl,
	repository.NewUserRepository,
	usecaseImpl.NewAuthenticationUsecaseImpl,
	handler.NewAuthHandler,
)

var photoSet = wire.NewSet(
	repositoryImpl.NewPhotoRepository,
	repositoryImpl.NewCommentRepository,
	repositoryImpl.NewTagRepositoryImpl,
	repositoryImpl.NewPhotoTagsRepositoryImpl,
	usecaseImpl.NewPhotoUsecase,
	handler.NewPhotoHandler,
)
var commentSet = wire.NewSet(
	repositoryImpl.NewCommentRepository, repositoryImpl.NewPhotoRepository, usecaseImpl.NewCommentUsecase, handler.NewCommentHandler,
)

var likesSet = wire.NewSet(repositoryImpl.NewUserLikesPhotoRepository, repositoryImpl.NewPhotoRepository, repository.NewUserRepository, usecaseImpl.NewUserLikesPhotosUsecase, handler.NewUserLikesPhotosHandler)

func initializedLikesHandler() handler.UserLikesPhotosHandler {
	wire.Build(config.InitializeDB, validator.New, likesSet)
	return nil
}

func initializedAuthHandler() handler.AuthHandler {
	wire.Build(config.InitializeDB, validator.New, authenticationSet)
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

func initializedFollowHandler() handler.FollowHandler {
	wire.Build(config.InitializeDB, followsSet)
	return nil
}

func InitializedServer() *gin.Engine {
	wire.Build(initializedAuthHandler,
		initializedPhotoHandler,
		initializedCommentHandler,
		initializedLikesHandler,
		initializedFollowHandler,
		routes.NewRouter,
	)
	return nil
}
