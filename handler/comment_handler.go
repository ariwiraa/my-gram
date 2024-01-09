package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ariwiraa/my-gram/domain/dtos/request"

	"github.com/ariwiraa/my-gram/helpers"
	"github.com/ariwiraa/my-gram/usecase"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	validate       *validator.Validate
}

// DeleteComment godoc
// @Summary Delete comment identified by the given id
// @Description Delete the comment corresponding to the input Id
// @Tags comment
// @Accept json
// @Produce json
// @Param id path int true "ID of the comment to be deleted"
// @Security JWT
// @Success 200 {object} helpers.SuccessResult{data=domain.Comment,code=int,message=string}
// @Failure 400 {object} helpers.BadRequest{code=int,message=string}
// @Success 500 {object} helpers.InternalServerError{code=int,message=string}
// @Router /comment/{id} [delete]
// DeleteCommentHandler implements CommentHandler
func (h *commentHandler) DeleteCommentHandler(ctx *gin.Context) {
	photoId := ctx.Param("id")

	requestParam := ctx.Param("commentId")
	commentId, _ := strconv.Atoi(requestParam)

	h.commentUsecase.Delete(ctx.Request.Context(), uint(commentId), photoId)
	helpers.NewResponse(
		helpers.WithHttpCode(http.StatusOK),
		helpers.WithMessage("delete comment success"),
	).Send(ctx)
}

// GetComment godoc
// @Summary Get Details for a given id
// @Description Get details of comment corresponding is the input Id
// @Tags comment
// @Accept json
// @Produce json
// @Param id path int true "ID of the comment"
// @Security JWT
// @Success 200 {object} helpers.SuccessResult{data=domain.Comment,code=int,message=string}
// @Failure 400 {object} helpers.BadRequest{code=int,message=string}
// @Success 500 {object} helpers.InternalServerError{code=int,message=string}
// @Router /comment/{id} [get]
// GetCommentHandler implements CommentHandler
func (h *commentHandler) GetCommentHandler(ctx *gin.Context) {
	photoId := ctx.Param("id")
	requestParam := ctx.Param("commentId")
	commentId, _ := strconv.Atoi(requestParam)

	comment, err := h.commentUsecase.GetById(ctx.Request.Context(), uint(commentId), photoId)
	if err != nil {
		log.Printf("[GetCommentHandler, GetById] with error detail %v", err.Error())
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
		helpers.WithMessage("get comment success"),
		helpers.WithPayload(comment),
	).Send(ctx)

}

// GetComment godoc
// @Summary Get All comment
// @Description Get All comment
// @Tags comment
// @Accept json
// @Produce json
// @Security JWT
// @Success 200 {object} helpers.SuccessResult{data=domain.Comment,code=int,message=string}
// @Failure 400 {object} helpers.BadRequest{code=int,message=string}
// @Success 500 {object} helpers.InternalServerError{code=int,message=string}
// @Router /comment [get]
// GetCommentsHandler implements CommentHandler
func (h *commentHandler) GetCommentsHandler(ctx *gin.Context) {
	photoId := ctx.Param("id")

	comments, err := h.commentUsecase.GetAllCommentsByPhotoId(ctx.Request.Context(), photoId)

	if err != nil {
		log.Printf("[GetCommentsHandler, GetAllCommentsByPhotoId] with error detail %v", err.Error())
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
		helpers.WithMessage("get all comments success"),
		helpers.WithPayload(comments),
	).Send(ctx)
}

// CreateComment godoc
// @Summary Post Details
// @Description Post details of comment
// @Tags comment
// @Accept json
// @Produce json
// @Param comment body dtos.CommentRequest true "create comment"
// @Security JWT
// @Success 200 {object} helpers.SuccessResult{data=domain.Comment,code=int,message=string}
// @Failure 400 {object} helpers.BadRequest{code=int,message=string}
// @Success 500 {object} helpers.InternalServerError{code=int,message=string}
// @Router /comment [post]
// PostCommentHandler implements CommentHandler
func (h *commentHandler) PostCommentHandler(ctx *gin.Context) {
	var payload request.CommentRequest

	photoId := ctx.Param("id")
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["Id"].(float64))

	payload.UserId = userID
	payload.PhotoId = photoId

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		log.Printf("[PostCommentHandler, ShouldBindJSON] with error detail %v", err.Error())
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
		log.Printf("[PostCommentHandler, Struct] with error detail %v", err.Error())
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

	comment, err := h.commentUsecase.Create(ctx.Request.Context(), payload)
	if err != nil {
		log.Printf("[PostCommentHandler, Create] with error detail %v", err.Error())
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
		helpers.WithMessage("create comment success"),
		helpers.WithPayload(comment),
	).Send(ctx)

}

// UpdateComment godoc
// @Summary put Details
// @Description put details of comment
// @Tags comment
// @Accept json
// @Produce json
// @Param comment body dtos.CommentRequest true "create comment"
// @Param id path int true "ID of the comment"
// @Security JWT
// @Success 200 {object} helpers.SuccessResult{data=domain.Comment,code=int,message=string}
// @Failure 400 {object} helpers.BadRequest{code=int,message=string}
// @Success 500 {object} helpers.InternalServerError{code=int,message=string}
// @Router /comment/{id} [put]
// PutCommentHandler implements CommentHandler
func (h *commentHandler) PutCommentHandler(ctx *gin.Context) {
	photoId := ctx.Param("id")

	requestParam := ctx.Param("commentId")
	commentId, _ := strconv.Atoi(requestParam)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["Id"].(float64))

	var payload request.CommentRequest
	payload.PhotoId = photoId
	payload.UserId = userID

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		log.Printf("[PutCommentHandler, ShouldBindJSON] with error detail %v", err.Error())
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
		log.Printf("[PuttCommentHandler, Struct] with error detail %v", err.Error())
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

	comment, err := h.commentUsecase.Update(ctx.Request.Context(), payload, uint(commentId))
	if err != nil {
		log.Printf("[PutCommentHandler, Update] with error detail %v", err.Error())
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
		helpers.WithMessage("update comment success"),
		helpers.WithPayload(comment),
	).Send(ctx)

}

func NewCommentHandler(commentUsecase usecase.CommentUsecase, validate *validator.Validate) CommentHandler {
	return &commentHandler{commentUsecase: commentUsecase, validate: validate}
}
