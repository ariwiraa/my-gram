package handler

import (
	"net/http"
	"strconv"

	"github.com/ariwiraa/my-gram/helpers"
	"github.com/ariwiraa/my-gram/usecase"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserLikesPhotosHandler interface {
	PostLikesHandler(ctx *gin.Context)
}

type userLikesPhotosHandler struct {
	likesUsecase usecase.UserLikesPhotosUsecase
	validate     *validator.Validate
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
	requestParam := ctx.Param("id")
	photoId, _ := strconv.Atoi(requestParam)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	likes, err := h.likesUsecase.LikeThePhoto(uint(photoId), userId)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, http.StatusCreated, likes)
}

func NewUserLikesPhotosHandler(likesUsecase usecase.UserLikesPhotosUsecase, validate *validator.Validate) UserLikesPhotosHandler {
	return &userLikesPhotosHandler{likesUsecase: likesUsecase, validate: validate}
}
