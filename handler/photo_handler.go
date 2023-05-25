package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ariwiraa/my-gram/domain"
	"github.com/ariwiraa/my-gram/domain/dtos"
	"github.com/ariwiraa/my-gram/helpers"
	"github.com/ariwiraa/my-gram/usecase"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	var photo domain.Photo
	requestParam := ctx.Param("id")
	photoId, _ := strconv.Atoi(requestParam)
	photo.ID = uint(photoId)

	h.photoUsecase.Delete(photo)
	helpers.SuccessResponse(ctx, http.StatusOK, nil)
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
	requestParam := ctx.Param("id")
	photoId, _ := strconv.Atoi(requestParam)

	photo, err := h.photoUsecase.GetById(uint(photoId))
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
	}

	helpers.SuccessResponse(ctx, http.StatusOK, photo)

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
	photos, err := h.photoUsecase.GetAll()
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
	}

	helpers.SuccessResponse(ctx, http.StatusOK, photos)
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
	var payload dtos.PhotoRequest

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// ambil data dari formdata
	payload.Caption = ctx.PostForm("caption")
	// err := ctx.ShouldBind(&payload)
	// if err != nil {
	// 	helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
	// 	return
	// }

	// err = h.validate.Struct(payload)
	// if err != nil {
	// 	errorMessage := helpers.FormatValidationErrors(err)
	// 	helpers.FailResponse(ctx, http.StatusBadRequest, errorMessage)
	// 	return
	// }

	// ambil file yang diupload
	file, err := ctx.FormFile("photo_url")
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, "gagal mengambil file")
		return
	}

	//tentukan destinasi penyimpanan file
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	//save file yang diupload
	err = ctx.SaveUploadedFile(file, path)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	photo, err := h.photoUsecase.Create(payload, userID, path)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, http.StatusOK, photo)

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
	requestParam := ctx.Param("id")
	photoId, _ := strconv.Atoi(requestParam)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	var payload dtos.PhotoRequest
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

func NewPhotoHandler(photoUsecase usecase.PhotoUsecase, validate *validator.Validate) PhotoHandler {
	return &photoHandler{photoUsecase: photoUsecase, validate: validate}
}
