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
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	loggedInUser, err := h.userUsecase.Login(payload)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	token := helpers.GenerateToken(loggedInUser.ID, loggedInUser.Email)
	ctx.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})
}

// PostUserRegisterHandler implements UserHandler
func (h *userHandler) PostUserRegisterHandler(ctx *gin.Context) {
	var payload domain.UserRequest

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newUser, err := h.userUsecase.Register(payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":       newUser.ID,
		"email":    newUser.Email,
		"username": newUser.Username,
	})
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: userUsecase}
}
