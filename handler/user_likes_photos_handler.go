package handler

import (
	"log"
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

	likes, err := h.likesUsecase.GetPhotosLikedByUserId(ctx.Request.Context(), userId)
	if err != nil {
		log.Printf("[GetPhotosLikedHandler, GetPhotosLikedByUserId] with error detail %v", err.Error())
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
		helpers.WithMessage("get photos liked success"),
		helpers.WithPayload(likes),
	).Send(ctx)
}

func (h *userLikesPhotosHandler) GetUsersWhoLikedPhotosHandler(ctx *gin.Context) {
	photoId := ctx.Param("id")

	users, err := h.likesUsecase.GetUsersWhoLikedPhotoByPhotoId(ctx.Request.Context(), photoId)
	if err != nil {
		log.Printf("[GetUsersWhoLikedPhotosHandler, GetUsersWhoLikedPhotoByPhotoId] with error detail %v", err.Error())
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
		helpers.WithMessage("get users who liked photo success"),
		helpers.WithPayload(users),
	).Send(ctx)
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

	likes, err := h.likesUsecase.LikeThePhoto(ctx.Request.Context(), photoId, userId)
	if err != nil {
		log.Printf("[PostLikesHandler, LikeThePhoto] with error detail %v", err.Error())
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
		helpers.WithMessage(likes),
	).Send(ctx)
}

func NewUserLikesPhotosHandler(likesUsecase usecase.UserLikesPhotosUsecase, validate *validator.Validate) UserLikesPhotosHandler {
	return &userLikesPhotosHandler{likesUsecase: likesUsecase, validate: validate}
}
