package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ariwiraa/my-gram/domain/dtos/request"
	"github.com/ariwiraa/my-gram/helpers"
	"github.com/ariwiraa/my-gram/usecase"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type FollowHandler interface {
	PostFollowHandler(ctx *gin.Context)
	GetFollowersHandler(ctx *gin.Context)
	GetFollowingsHandler(ctx *gin.Context)
}

type followHandlerImpl struct {
	followUsecase usecase.FollowUsecase
}

func (h *followHandlerImpl) GetFollowersHandler(ctx *gin.Context) {
	username := ctx.Param("username")

	followers, err := h.followUsecase.GetFollowersByUsername(ctx.Request.Context(), username)
	if err != nil {
		log.Printf("[GetFollowersHandler, GetFollowersByUsername] with error detail %v", err.Error())
		myErr, ok := helpers.ErrorMapping[err.Error()]

		if !ok {
			myErr = helpers.ErrorGeneral
		}

		helpers.NewResponse(
			helpers.WithMessage(err.Error()),
			helpers.WithError(myErr),
		).Send(ctx)
		return
	}

	helpers.NewResponse(
		helpers.WithHttpCode(http.StatusOK),
		helpers.WithMessage("get follower success"),
		helpers.WithPayload(followers),
	).Send(ctx)
}

func (h *followHandlerImpl) GetFollowingsHandler(ctx *gin.Context) {
	username := ctx.Param("username")

	followings, err := h.followUsecase.GetFollowingsByUsername(ctx.Request.Context(), username)
	if err != nil {
		log.Printf("[GetFollowingsHandler, GetFollowingsByUsername] with error detail %v", err.Error())
		myErr, ok := helpers.ErrorMapping[err.Error()]

		if !ok {
			myErr = helpers.ErrorGeneral
		}

		helpers.NewResponse(
			helpers.WithMessage(err.Error()),
			helpers.WithError(myErr),
		).Send(ctx)
		return
	}

	helpers.NewResponse(
		helpers.WithHttpCode(http.StatusOK),
		helpers.WithMessage("get following uccess"),
		helpers.WithPayload(followings),
	).Send(ctx)
}

func (h *followHandlerImpl) PostFollowHandler(ctx *gin.Context) {
	params := ctx.Param("id")
	userIdFollowing, err := strconv.Atoi(params)
	if err != nil {
		log.Printf("[PostFollowHandler, Atoi] with error detail %v", err.Error())
		myErr := helpers.ErrorGeneral
		helpers.NewResponse(
			helpers.WithMessage(err.Error()),
			helpers.WithError(myErr),
			helpers.WithHttpCode(http.StatusInternalServerError),
		).Send(ctx)
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userIdFollower := uint(userData["Id"].(float64))

	payload := request.FollowRequest{
		UserIdFollowing: uint(userIdFollowing),
		UserIdFollower:  userIdFollower,
	}

	message, err := h.followUsecase.FollowUser(ctx.Request.Context(), payload)
	if err != nil {
		log.Printf("[GetFollowersHandler, GetFollowersByUsername] with error detail %v", err.Error())
		myErr, ok := helpers.ErrorMapping[err.Error()]

		if !ok {
			myErr = helpers.ErrorGeneral
		}

		helpers.NewResponse(
			helpers.WithMessage(err.Error()),
			helpers.WithError(myErr),
		).Send(ctx)
		return
	}

	helpers.NewResponse(
		helpers.WithHttpCode(http.StatusOK),
		helpers.WithMessage(message),
	).Send(ctx)
}

func NewFollowHandlerImpl(followUsecase usecase.FollowUsecase) FollowHandler {
	return &followHandlerImpl{followUsecase: followUsecase}
}
