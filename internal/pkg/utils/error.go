package httputils

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Example return message"`
}

func NewErrorResponse(ctx *gin.Context, code int, message string) {
	err := ErrorResponse{
		Code:    code,
		Message: message,
	}
	ctx.JSON(code, err)
}
