package handler

import (
	"net/http"

	"github.com/ariwiraa/my-gram/helpers"
	"github.com/ariwiraa/my-gram/usecase"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UploadFileHandler struct {
	uploadFileUsecase usecase.UploadFileUsecase
}

func NewUploadFileHandler(uploadFileUsecase usecase.UploadFileUsecase) *UploadFileHandler {
	return &UploadFileHandler{
		uploadFileUsecase: uploadFileUsecase,
	}
}

func (h *UploadFileHandler) UploadFileHandler(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["Id"].(float64))

	file, err := ctx.FormFile("file")
	if err != nil {
		helpers.NewResponse(
			helpers.WithHttpCode(http.StatusInternalServerError),
			helpers.WithError(err),
			helpers.WithMessage(err.Error()),
		).Send(ctx)
		return
	}

	url, err := h.uploadFileUsecase.Upload(ctx.Request.Context(), file, userID)
	if err != nil {
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
		helpers.WithHttpCode(http.StatusCreated),
		helpers.WithMessage("berhasil upload file"),
		helpers.WithPayload(url),
	).Send(ctx)
}
