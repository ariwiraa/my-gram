package handler

import (
	"log"
	"net/http"

	"github.com/ariwiraa/my-gram/helpers"
	"github.com/ariwiraa/my-gram/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetUserProfileHandler(ctx *gin.Context)
}

type userHandlerImpl struct {
	userUsecase usecase.UserUsecase
}

func (h *userHandlerImpl) GetUserProfileHandler(ctx *gin.Context) {
	username := ctx.Param("username")

	profileResponse, err := h.userUsecase.GetUserProfileByUsername(ctx.Request.Context(), username)
	if err != nil {

		log.Printf("[GetUserProfileHandler, GetUserProfileByUsername] with error detail %v", err.Error())
		errorMessage := helpers.FormatValidationErrors(err)

		myErr, ok := helpers.ErrorMapping[errorMessage.Error()]

		if !ok {
			myErr = helpers.ErrorGeneral
		}

		helpers.NewResponse(
			helpers.WithMessage(errorMessage.Error()),
			helpers.WithError(myErr),
			helpers.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
		return

	}

	helpers.NewResponse(
		helpers.WithHttpCode(http.StatusOK),
		helpers.WithMessage("get profile success"),
		helpers.WithPayload(profileResponse),
	).Send(ctx)
}

func NewUserHandlerImpl(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandlerImpl{userUsecase: userUsecase}
}
