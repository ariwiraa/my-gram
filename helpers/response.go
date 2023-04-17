package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type BadRequest struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type InternalServerError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, SuccessResult{
		Code:    code,
		Message: "success",
		Data:    data,
	})
}

func FailResponse(ctx *gin.Context, code int, message string) {
	if code == http.StatusInternalServerError {
		ctx.JSON(code, InternalServerError{
			Code:    code,
			Message: message,
			Data:    nil,
		})
		return
	}

	ctx.JSON(code, BadRequest{
		Code:    code,
		Message: message,
		Data:    nil,
	})

}
