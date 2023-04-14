package handler

import (
	"net/http"
	"strconv"

	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/helpers"
	"github.com/ariwiraa/my-gram/usecase"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type PhotoHandler interface {
	PostPhotoHandler(ctx *gin.Context)
	GetPhotosHandler(ctx *gin.Context)
	GetPhotoHandler(ctx *gin.Context)
	PutPhotoHandler(ctx *gin.Context)
	DeletePhotoHandler(ctx *gin.Context)
}

type photoHandler struct {
	photoUsecase usecase.PhotoUsecase
}

// DeletePhotoHandler implements PhotoHandler
func (h *photoHandler) DeletePhotoHandler(ctx *gin.Context) {
	var photo domain.Photo
	requestParam := ctx.Param("id")
	photoId, _ := strconv.Atoi(requestParam)
	photo.ID = uint(photoId)

	h.photoUsecase.Delete(photo)
	helpers.SuccessResponse(ctx, http.StatusOK, nil)
}

// GetPhotoHandler implements PhotoHandler
func (h *photoHandler) GetPhotoHandler(ctx *gin.Context) {
	requestParam := ctx.Param("id")
	photoId, _ := strconv.Atoi(requestParam)

	photo, err := h.photoUsecase.GetById(uint(photoId))
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
	}

	helpers.SuccessResponse(ctx, http.StatusOK, photo)

}

// GetPhotosHandler implements PhotoHandler
func (h *photoHandler) GetPhotosHandler(ctx *gin.Context) {
	photos, err := h.photoUsecase.GetAll()
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
	}

	helpers.SuccessResponse(ctx, http.StatusOK, photos)
}

// CreatePhoto godoc
// @Summary Post Details
// @Description Post details of social media
// @Tags social media
// @Accept json
// @Produce json
// @Param photo body domain.PhotoRequest true "create social media"
// @Success 200 {object} helpers.SuccessResult{data=domain.Photo,code=int,message=string}
// @Failure 400 {object} helpers.BadRequest{code=int,message=string}
// @Success 500 {object} helpers.InternalServerError{code=int,message=string}
// @Router /socialmedia [post]
// PostPhotoHandler implements PhotoHandler
func (h *photoHandler) PostPhotoHandler(ctx *gin.Context) {
	var payload domain.PhotoRequest

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	photo, err := h.photoUsecase.Create(payload, userID)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, http.StatusOK, photo)

}

// PutPhotoHandler implements PhotoHandler
func (h *photoHandler) PutPhotoHandler(ctx *gin.Context) {
	requestParam := ctx.Param("id")
	photoId, _ := strconv.Atoi(requestParam)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	var payload domain.PhotoRequest
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	photo, err := h.photoUsecase.Update(payload, uint(photoId), userID)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, http.StatusOK, photo)

}

func NewPhotoHandler(photoUsecase usecase.PhotoUsecase) PhotoHandler {
	return &photoHandler{photoUsecase: photoUsecase}
}
