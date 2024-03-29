package handler

import (
	"log"
	"net/http"

	"github.com/ariwiraa/my-gram/domain/dtos/request"

	"github.com/ariwiraa/my-gram/helpers"
	"github.com/ariwiraa/my-gram/usecase"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type PhotoHandler interface {
	PostPhotoHandler(ctx *gin.Context)
	GetPhotosHandler(ctx *gin.Context)
	GetPhotosByUserIdHandler(ctx *gin.Context)
	GetPhotoHandler(ctx *gin.Context)
	PutPhotoHandler(ctx *gin.Context)
	DeletePhotoHandler(ctx *gin.Context)
}

type photoHandler struct {
	photoUsecase usecase.PhotoUsecase
	validate     *validator.Validate
}

// Deletephoto godoc
// @Summary Delete photo identified by the given id
// @Description Delete the photo corresponding to the input Id
// @Tags photo
// @Accept json
// @Produce json
// @Param id path int true "ID of the photo to be deleted"
// @Security JWT
// @Success 200 {object} helpers.SuccessResult{data=domain.Photo,code=int,message=string}
// @Failure 400 {object} helpers.BadRequest{code=int,message=string}
// @Success 500 {object} helpers.InternalServerError{code=int,message=string}
// @Router /photo/{id} [delete]
// DeletePhotoHandler implements PhotoHandler
func (h *photoHandler) DeletePhotoHandler(ctx *gin.Context) {

	photoId := ctx.Param("id")

	err := h.photoUsecase.Delete(ctx.Request.Context(), photoId)
	if err != nil {
		log.Printf("[DeletePhotoHandler, DeleteById] with error detail %v", err.Error())
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
		helpers.WithMessage("delete photo success"),
	).Send(ctx)
}

// Getphoto godoc
// @Summary Get Details for a given id
// @Description Get details of photo corresponding is the input Id
// @Tags photo
// @Accept json
// @Produce json
// @Param id path int true "ID of the photo"
// @Security JWT
// @Success 200 {object} helpers.SuccessResult{data=domain.Photo,code=int,message=string}
// @Failure 400 {object} helpers.BadRequest{code=int,message=string}
// @Success 500 {object} helpers.InternalServerError{code=int,message=string}
// @Router /photo/{id} [get]
// GetPhotoHandler implements PhotoHandler
func (h *photoHandler) GetPhotoHandler(ctx *gin.Context) {
	photoId := ctx.Param("id")

	photo, err := h.photoUsecase.GetById(ctx.Request.Context(), photoId)
	if err != nil {
		log.Printf("[GetPhotoHandler, GetById] with error detail %v", err.Error())
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
		helpers.WithMessage("get photo success"),
		helpers.WithPayload(photo),
	).Send(ctx)

}

// Getphotos godoc
// @Summary Get All photos
// @Description Get All photos
// @Tags photo
// @Accept json
// @Produce json
// @Param id path int true "ID of the photo"
// @Security JWT
// @Success 200 {object} helpers.SuccessResult{data=domain.Photo,code=int,message=string}
// @Failure 400 {object} helpers.BadRequest{code=int,message=string}
// @Success 500 {object} helpers.InternalServerError{code=int,message=string}
// @Router /photo [get]
// GetPhotosHandler implements PhotoHandler
func (h *photoHandler) GetPhotosHandler(ctx *gin.Context) {
	photos, err := h.photoUsecase.GetAll(ctx.Request.Context())
	if err != nil {
		log.Printf("[GetPhotosHandler, GetAll] with error detail %v", err.Error())
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
		helpers.WithMessage("get all photos success"),
		helpers.WithPayload(photos),
	).Send(ctx)
}

// Getphoto godoc
// @Summary Get Details for a given id
// @Description Get details of photo corresponding is the input Id
// @Tags photo
// @Accept json
// @Produce json
// @Param id path int true "ID of the photo"
// @Security JWT
// @Success 200 {object} helpers.SuccessResult{data=domain.Photo,code=int,message=string}
// @Failure 400 {object} helpers.BadRequest{code=int,message=string}
// @Success 500 {object} helpers.InternalServerError{code=int,message=string}
// @Router /photo/{id} [get]
// GetPhotoHandler implements PhotoHandler
func (h *photoHandler) GetPhotosByUserIdHandler(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["Id"].(float64))

	photo, err := h.photoUsecase.GetAllPhotosByUserId(ctx.Request.Context(), userID)
	if err != nil {
		log.Printf("[GetPhotosHandler, GetAll] with error detail %v", err.Error())
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
		helpers.WithMessage("get all photos success"),
		helpers.WithPayload(photo),
	).Send(ctx)

}

// CreatePhoto godoc
// @Summary Post Details
// @Description Post details of photo
// @Tags photo
// @Accept json
// @Produce json
// @Param photo body dtos.PhotoRequest true "create photo"
// @Security JWT
// @Success 200 {object} helpers.SuccessResult{data=domain.Photo,code=int,message=string}
// @Failure 400 {object} helpers.BadRequest{code=int,message=string}
// @Success 500 {object} helpers.InternalServerError{code=int,message=string}
// @Router /photo [post]
// PostPhotoHandler implements PhotoHandler
func (h *photoHandler) PostPhotoHandler(ctx *gin.Context) {
	var payload request.PhotoRequest

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["Id"].(float64))

	// ambil data dari formdata
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		log.Printf("[PostPhotoHandler, ShouldBindJSON] with error detail %v", err.Error())
		myErr := helpers.ErrorGeneral
		helpers.NewResponse(
			helpers.WithMessage(err.Error()),
			helpers.WithError(myErr),
			helpers.WithHttpCode(http.StatusInternalServerError),
		).Send(ctx)
		return
	}

	err = h.validate.Struct(payload)
	if err != nil {
		log.Printf("[PostPhotoHandler, Struct] with error detail %v", err.Error())
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

	photo, err := h.photoUsecase.Create(ctx.Request.Context(), payload, userID)
	if err != nil {
		log.Printf("[PostPhotoHandler, Create] with error detail %v", err.Error())
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
		helpers.WithMessage("create photo success"),
		helpers.WithPayload(photo),
	).Send(ctx)

}

// UpdatePhoto godoc
// @Summary put Details
// @Description put details of photo
// @Tags photo
// @Accept json
// @Produce json
// @Param photo body dtos.PhotoRequest true "create photo"
// @Param id path int true "ID of the photo"
// @Security JWT
// @Success 200 {object} helpers.SuccessResult{data=domain.Photo,code=int,message=string}
// @Failure 400 {object} helpers.BadRequest{code=int,message=string}
// @Success 500 {object} helpers.InternalServerError{code=int,message=string}
// @Router /photo/{id} [put]
// PutPhotoHandler implements PhotoHandler
func (h *photoHandler) PutPhotoHandler(ctx *gin.Context) {
	photoId := ctx.Param("id")

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["Id"].(float64))

	var payload request.UpdatePhotoRequest
	err := ctx.ShouldBindWith(&payload, binding.FormMultipart)
	if err != nil {
		log.Printf("[PutPhotoHandler, ShouldBindWith] with error detail %v", err.Error())
		myErr := helpers.ErrorGeneral
		helpers.NewResponse(
			helpers.WithMessage(err.Error()),
			helpers.WithError(myErr),
			helpers.WithHttpCode(http.StatusInternalServerError),
		).Send(ctx)
		return
	}

	err = h.validate.Struct(payload)
	if err != nil {
		log.Printf("[PutPhotoHandler, Struct] with error detail %v", err.Error())
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

	photo, err := h.photoUsecase.Update(ctx.Request.Context(), payload, photoId, userID)
	if err != nil {
		log.Printf("[PutPhotoHandler, Update] with error detail %v", err.Error())
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
		helpers.WithMessage("update photo success"),
		helpers.WithPayload(photo),
	).Send(ctx)

}

func NewPhotoHandler(photoUsecase usecase.PhotoUsecase, validate *validator.Validate) PhotoHandler {
	return &photoHandler{photoUsecase: photoUsecase, validate: validate}
}
