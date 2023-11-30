package handler

import (
	"net/http"

	"github.com/ariwiraa/my-gram/helpers"
	"github.com/ariwiraa/my-gram/usecase"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserLikesPhotosHandler interface {
	PostLikesHandler(ctx *gin.Context)
	GetUsersWhoLikedPhotosHandler(ctx *gin.Context)
	GetPhotosLikedHandler(ctx *gin.Context)
}

type userLikesPhotosHandler struct {
	likesUsecase usecase.UserLikesPhotosUsecase
	validate     *validator.Validate
}

func (h *userLikesPhotosHandler) GetPhotosLikedHandler(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["Id"].(float64))

	likes, err := h.likesUsecase.GetPhotosLikedByUserId(userId)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, http.StatusOK, likes)
}

func (h *userLikesPhotosHandler) GetUsersWhoLikedPhotosHandler(ctx *gin.Context) {
	photoId := ctx.Param("id")

	users, err := h.likesUsecase.GetUsersWhoLikedPhotoByPhotoId(photoId)
	if err != nil {
		helpers.FailResponse(ctx, 400, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, 200, users)
}

// LikePhoto godoc
// @Summary like photo
// @Description user can like the photo
// @Tags likes
// @Accept json
// @Produce json
// @Param id path int true "ID of the comment"
// @Security JWT
// @Success 200 {object} helpers.SuccessResult{data=string,code=int,message=string}
// @Failure 400 {object} helpers.BadRequest{code=int,message=string}
// @Success 500 {object} helpers.InternalServerError{code=int,message=string}
// @Router /photo/{id}/likes [post]
// PostLikesHandler implements UserLikesPhotosHandler
func (h *userLikesPhotosHandler) PostLikesHandler(ctx *gin.Context) {
	photoId := ctx.Param("id")

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["Id"].(float64))

	likes, err := h.likesUsecase.LikeThePhoto(photoId, userId)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, http.StatusCreated, likes)
}

func NewUserLikesPhotosHandler(likesUsecase usecase.UserLikesPhotosUsecase, validate *validator.Validate) UserLikesPhotosHandler {
	return &userLikesPhotosHandler{likesUsecase: likesUsecase, validate: validate}
}
