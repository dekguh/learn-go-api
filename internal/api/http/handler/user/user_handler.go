package handler

import (
	"net/http"

	"github.com/dekguh/learn-go-api/internal/api/repository"
	"github.com/dekguh/learn-go-api/internal/api/service"
	httputils "github.com/dekguh/learn-go-api/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	service service.UserService
}

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	group := r.Group("/users")
	{
		group.GET("/detail/email/:email", userHandler.GetUserDetailByEmail)
	}
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (handler *UserHandler) GetUserDetailByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		httputils.NewErrorResponse(ctx, http.StatusBadRequest, "email is required")
		return
	}

	user, err := handler.service.GetUserByEmail(email)
	if err != nil {
		httputils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	httputils.NewSuccessResponse(ctx, http.StatusOK, "success", user)
}
