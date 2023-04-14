package handler

import (
	"net/http"

	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/helpers"
	"github.com/ariwiraa/my-gram/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	PostUserRegisterHandler(ctx *gin.Context)
	PostUserLoginHandler(ctx *gin.Context)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

// PostUserLoginHandler implements UserHandler
func (h *userHandler) PostUserLoginHandler(ctx *gin.Context) {
	var payload domain.UserLogin

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
	}

	loggedInUser, err := h.userUsecase.Login(payload)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	token := helpers.GenerateToken(loggedInUser.ID, loggedInUser.Email)
	helpers.SuccessResponse(ctx, http.StatusOK, gin.H{
		"access_token": token,
	})
}

// PostUserRegisterHandler implements UserHandler
func (h *userHandler) PostUserRegisterHandler(ctx *gin.Context) {
	var payload domain.UserRequest

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	newUser, err := h.userUsecase.Register(payload)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, http.StatusOK, gin.H{
		"id":       newUser.ID,
		"email":    newUser.Email,
		"username": newUser.Username,
	})
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: userUsecase}
}
