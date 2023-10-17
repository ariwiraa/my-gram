package handler

import (
	"net/http"

	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/domain/dtos"
	"github.com/ariwiraa/my-gram/helpers"
	"github.com/ariwiraa/my-gram/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler interface {
	PostUserRegisterHandler(ctx *gin.Context)
	PostUserLoginHandler(ctx *gin.Context)
	PutAccessTokenHandler(ctx *gin.Context)
	LogoutHandler(ctx *gin.Context)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
	authUsecase usecase.AuthenticationUsecase
	validate    *validator.Validate
}

// PutAccessTokenHandler implements UserHandler.
func (h *userHandler) PutAccessTokenHandler(ctx *gin.Context) {
	var payload domain.Authentication

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.authUsecase.ExistsByRefreshToken(payload.RefreshToken)
	if err != nil {
		return
	}

	claims, err := helpers.VerifyRefreshToken(payload.RefreshToken)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	accessToken := helpers.NewAccessToken(claims.Id).GenerateAccessToken()

	helpers.SuccessResponse(ctx, http.StatusOK, gin.H{
		"access_token": accessToken,
	})

}

// Logout implements UserHandler.
func (h *userHandler) LogoutHandler(ctx *gin.Context) {
	var payload domain.Authentication

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.authUsecase.Delete(payload.RefreshToken)
	if err != nil {
		return
	}

	helpers.SuccessResponse(ctx, http.StatusOK, nil)

}

// UserLogin godoc
// @Summary User Login
// @Description user logs in
// @Tags user
// @Accept json
// @Produce json
// @Param login body dtos.UserLogin true "logged in"
// @Success 200 {object} helpers.SuccessResult{data=string,code=int,message=string}
// @Failure 400 {object} helpers.BadRequest{code=int,message=string}
// @Success 500 {object} helpers.InternalServerError{code=int,message=string}
// @Router /signin [post]
// PostUserLoginHandler implements UserHandler
func (h *userHandler) PostUserLoginHandler(ctx *gin.Context) {
	var payload dtos.UserLogin

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	loggedInUser, err := h.userUsecase.Login(payload)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	accessToken := helpers.NewAccessToken(uint64(loggedInUser.ID)).GenerateAccessToken()
	refreshToken := helpers.NewRefreshToken(uint64(loggedInUser.ID)).GenerateRefreshToken()

	err = h.authUsecase.Add(refreshToken)
	if err != nil {
		return
	}

	helpers.SuccessResponse(ctx, http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// UserRegister godoc
// @Summary User Register
// @Description user registers in the form provided
// @Tags user
// @Accept json
// @Produce json
// @Param register body dtos.UserRequest true "create account"
// @Success 200 {object} helpers.SuccessResult{data=domain.User,code=int,message=string}
// @Failure 400 {object} helpers.BadRequest{code=int,message=string}
// @Success 500 {object} helpers.InternalServerError{code=int,message=string}
// @Router /signup [post]
// PostUserRegisterHandler implements UserHandler
func (h *userHandler) PostUserRegisterHandler(ctx *gin.Context) {
	var payload dtos.UserRequest

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.validate.Struct(payload)
	if err != nil {
		errorMessage := helpers.FormatValidationErrors(err)
		helpers.FailResponse(ctx, http.StatusBadRequest, errorMessage)
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

func NewUserHandler(userUsecase usecase.UserUsecase, authUsecase usecase.AuthenticationUsecase, validate *validator.Validate) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
		authUsecase: authUsecase,
		validate:    validate,
	}
}
