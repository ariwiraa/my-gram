package handler

import (
	"github.com/ariwiraa/my-gram/helpers"
	"github.com/ariwiraa/my-gram/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler interface {
	GetUserProfileHandler(ctx *gin.Context)
}

type userHandlerImpl struct {
	userUsecase usecase.UserUsecase
}

func (h *userHandlerImpl) GetUserProfileHandler(ctx *gin.Context) {
	username := ctx.Param("username")

	profileResponse, err := h.userUsecase.GetUserProfileByUsername(username)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, http.StatusOK, profileResponse)
}

func NewUserHandlerImpl(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandlerImpl{userUsecase: userUsecase}
}
