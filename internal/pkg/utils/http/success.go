package httputils

import "github.com/gin-gonic/gin"

type SuccessResponse struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"Example return message"`
	Data    any    `json:"data"`
}

func NewSuccessResponse(ctx *gin.Context, code int, message string, data any) {
	success := SuccessResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
	ctx.JSON(code, success)
}
