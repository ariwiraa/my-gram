package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccesResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type BadRequestResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type InteralServerError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, SuccesResult{
		Code:    code,
		Message: "success",
		Data:    data,
	})
}

func FailResponse(ctx *gin.Context, code int, message string) {
	if code == http.StatusInternalServerError {
		ctx.JSON(code, InteralServerError{
			Code:    code,
			Message: message,
			Data:    nil,
		})
		return
	}

	ctx.JSON(code, BadRequestResult{
		Code:    code,
		Message: message,
		Data:    nil,
	})

}
