package handler

import (
	"github.com/ariwiraa/my-gram/domain/dtos/request"
	"github.com/ariwiraa/my-gram/helpers"
	"github.com/ariwiraa/my-gram/usecase"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FollowHandler interface {
	PostFollowHandler(ctx *gin.Context)
}

type followHandlerImpl struct {
	followUsecase usecase.FollowUsecase
}

func (h *followHandlerImpl) PostFollowHandler(ctx *gin.Context) {
	params := ctx.Param("id")
	userIdFollowing, err := strconv.Atoi(params)
	if err != nil {
		helpers.FailResponse(ctx, 400, "failed to convert parameters")
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userIdFollower := uint(userData["Id"].(float64))

	payload := request.FollowRequest{
		UserIdFollowing: uint(userIdFollowing),
		UserIdFollower:  userIdFollower,
	}

	message, err := h.followUsecase.FollowUser(payload)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, http.StatusOK, message)
}

func NewFollowHandlerImpl(followUsecase usecase.FollowUsecase) FollowHandler {
	return &followHandlerImpl{followUsecase: followUsecase}
}
