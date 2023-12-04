package handler

import (
	"github.com/ariwiraa/my-gram/domain/dtos/request"
	"net/http"

	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/helpers"
	"github.com/ariwiraa/my-gram/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler interface {
	PostUserRegisterHandler(ctx *gin.Context)
	PostUserLoginHandler(ctx *gin.Context)
	PutAccessTokenHandler(ctx *gin.Context)
	LogoutHandler(ctx *gin.Context)
}

type authHandler struct {
	authUsecase usecase.AuthenticationUsecase
	validate    *validator.Validate
}

// PutAccessTokenHandler implements AuthHandler.
func (h *authHandler) PutAccessTokenHandler(ctx *gin.Context) {
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

// Logout implements AuthHandler.
func (h *authHandler) LogoutHandler(ctx *gin.Context) {
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
// PostUserLoginHandler implements AuthHandler
func (h *authHandler) PostUserLoginHandler(ctx *gin.Context) {
	var payload request.UserLogin

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	loggedInUser, err := h.authUsecase.Login(payload)
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
// @Param register body dtos.UserRegister true "create account"
// @Success 200 {object} helpers.SuccessResult{data=domain.User,code=int,message=string}
// @Failure 400 {object} helpers.BadRequest{code=int,message=string}
// @Success 500 {object} helpers.InternalServerError{code=int,message=string}
// @Router /signup [post]
// PostUserRegisterHandler implements AuthHandler
func (h *authHandler) PostUserRegisterHandler(ctx *gin.Context) {
	var payload request.UserRegister

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

	newUser, err := h.authUsecase.Register(payload)
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

func NewAuthHandler(authUsecase usecase.AuthenticationUsecase, validate *validator.Validate) AuthHandler {
	return &authHandler{
		authUsecase: authUsecase,
		validate:    validate,
	}
}