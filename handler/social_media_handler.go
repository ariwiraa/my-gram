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

type SocialMediaHandler interface {
	PostSocialMediaHandler(ctx *gin.Context)
	GetSocialMediasHandler(ctx *gin.Context)
	GetSocialMediaHandler(ctx *gin.Context)
	PutSocialMediaHandler(ctx *gin.Context)
	DeleteSocialMediaHandler(ctx *gin.Context)
}

type socialMediaHandler struct {
	socialMediaUsecase usecase.SocialMediaUsecase
}

// DeleteSocialMediaHandler implements SocialMediaHandler
func (h *socialMediaHandler) DeleteSocialMediaHandler(ctx *gin.Context) {
	var socialMedia domain.SocialMedia
	requestParam := ctx.Param("id")
	socialMediaId, _ := strconv.Atoi(requestParam)
	socialMedia.ID = uint(socialMediaId)

	h.socialMediaUsecase.Delete(socialMedia)
	helpers.SuccessResponse(ctx, http.StatusOK, nil)
}

// GetSocialMediaHandler implements SocialMediaHandler
func (h *socialMediaHandler) GetSocialMediaHandler(ctx *gin.Context) {
	requestParam := ctx.Param("id")
	socialMediaId, _ := strconv.Atoi(requestParam)

	socialMedia, err := h.socialMediaUsecase.GetById(uint(socialMediaId))
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
	}

	helpers.SuccessResponse(ctx, http.StatusOK, socialMedia)

}

// GetSocialMediasHandler implements SocialMediaHandler
func (h *socialMediaHandler) GetSocialMediasHandler(ctx *gin.Context) {
	socialMedias, err := h.socialMediaUsecase.GetAll()
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
	}

	helpers.SuccessResponse(ctx, http.StatusOK, socialMedias)
}

// CreateSocialMedia godoc
// @Summary Post Details
// @Description Post details of social media
// @Tags social media
// @Accept json
// @Produce json
// @Param socialMedia body domain.SocialMediaRequest true "create social media"
// @Success 200 {object} helpers.SuccessResult{data=domain.SocialMedia,code=int,message=string}
// @Failure 400 {object} helpers.BadRequest{code=int,message=string}
// @Success 500 {object} helpers.InternalServerError{code=int,message=string}
// @Router /socialmedia [post]
// PostSocialMediaHandler implements SocialMediaHandler
func (h *socialMediaHandler) PostSocialMediaHandler(ctx *gin.Context) {
	var payload domain.SocialMediaRequest

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	socialMedia, err := h.socialMediaUsecase.Create(payload, userID)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, http.StatusOK, socialMedia)

}

// PutSocialMediaHandler implements SocialMediaHandler
func (h *socialMediaHandler) PutSocialMediaHandler(ctx *gin.Context) {
	requestParam := ctx.Param("id")
	socialMediaId, _ := strconv.Atoi(requestParam)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	var payload domain.SocialMediaRequest
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	socialMedia, err := h.socialMediaUsecase.Update(payload, uint(socialMediaId), userID)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, http.StatusOK, socialMedia)

}

func NewSocialMediaHandler(socialMediaUsecase usecase.SocialMediaUsecase) SocialMediaHandler {
	return &socialMediaHandler{socialMediaUsecase: socialMediaUsecase}
}
