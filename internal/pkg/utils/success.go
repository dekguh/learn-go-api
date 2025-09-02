package httputils

import "github.com/gin-gonic/gin"

type SuccessResponse struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"Example return message"`
	Data    interface{} `json:"data" example:"{ success: true }"`
}

func NewSuccessResponse(ctx *gin.Context, code int, message string, data interface{}) {
	success := SuccessResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
	ctx.JSON(code, success)
}
