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

type CommentHandler interface {
	PostCommentHandler(ctx *gin.Context)
	GetCommentsHandler(ctx *gin.Context)
	GetCommentHandler(ctx *gin.Context)
	PutCommentHandler(ctx *gin.Context)
	DeleteCommentHandler(ctx *gin.Context)
}

type commentHandler struct {
	commentUsecase usecase.CommentUsecase
}

// DeleteCommentHandler implements CommentHandler
func (h *commentHandler) DeleteCommentHandler(ctx *gin.Context) {
	var comment domain.Comment
	requestParam := ctx.Param("id")
	commentId, _ := strconv.Atoi(requestParam)
	comment.ID = uint(commentId)

	h.commentUsecase.Delete(comment)
	helpers.SuccessResponse(ctx, http.StatusOK, nil)
}

// GetCommentHandler implements CommentHandler
func (h *commentHandler) GetCommentHandler(ctx *gin.Context) {
	requestParam := ctx.Param("id")
	commentId, _ := strconv.Atoi(requestParam)

	comment, err := h.commentUsecase.GetById(uint(commentId))
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
	}

	helpers.SuccessResponse(ctx, http.StatusOK, comment)

}

// GetCommentsHandler implements CommentHandler
func (h *commentHandler) GetCommentsHandler(ctx *gin.Context) {
	comments, err := h.commentUsecase.GetAll()
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
	}

	helpers.SuccessResponse(ctx, http.StatusOK, comments)
}

// CreateComment godoc
// @Summary Post Details
// @Description Post details of social media
// @Tags social media
// @Accept json
// @Produce json
// @Param comment body domain.CommentRequest true "create social media"
// @Success 200 {object} helpers.SuccessResult{data=domain.Comment,code=int,message=string}
// @Failure 400 {object} helpers.BadRequest{code=int,message=string}
// @Success 500 {object} helpers.InternalServerError{code=int,message=string}
// @Router /socialmedia [post]
// PostCommentHandler implements CommentHandler
func (h *commentHandler) PostCommentHandler(ctx *gin.Context) {
	var payload domain.CommentRequest

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	comment, err := h.commentUsecase.Create(payload, userID)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, http.StatusOK, comment)

}

// PutCommentHandler implements CommentHandler
func (h *commentHandler) PutCommentHandler(ctx *gin.Context) {
	requestParam := ctx.Param("id")
	commentId, _ := strconv.Atoi(requestParam)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	var payload domain.CommentRequest
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	comment, err := h.commentUsecase.Update(payload, uint(commentId), userID)
	if err != nil {
		helpers.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, http.StatusOK, comment)

}

func NewCommentHandler(commentUsecase usecase.CommentUsecase) CommentHandler {
	return &commentHandler{commentUsecase: commentUsecase}
}
