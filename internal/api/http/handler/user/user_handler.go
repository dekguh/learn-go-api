package handler

import (
	"net/http"

	"github.com/dekguh/learn-go-api/internal/api/repository"
	"github.com/dekguh/learn-go-api/internal/api/service"
	httputils "github.com/dekguh/learn-go-api/internal/pkg/utils"
	"github.com/dekguh/learn-go-api/internal/pkg/validator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	service service.UserService
}

type RegisterUserReq struct {
	Name     string `json:"name" binding:"required,min=3,max=64"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	group := r.Group("/users")
	{
		group.GET("/detail/email/:email", userHandler.GetUserDetailByEmail)
		group.POST("/register", userHandler.RegisterUser)
	}
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// @Summary Get user detail by email
// @Description Get user detail based by user email
// @Tags Users
// @Accept json
// @Produce json
// @Param email path string true "User email"
// @Success 200 {object} httputils.SuccessResponse{data=model.User}
// @Failure 400 {object} httputils.ErrorResponse
// @Router /users/detail/email/{email} [get]
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

// @Summary Register user
// @Description Register user
// @Tags Users
// @Accept json
// @Produce json
// @Param registerUserReq body RegisterUserReq true "Register user request"
// @Success 200 {object} httputils.SuccessResponse{data=model.User}
// @Failure 400 {object} httputils.ErrorResponse
// @Router /users/register [post]
func (handler *UserHandler) RegisterUser(ctx *gin.Context) {
	var json RegisterUserReq
	if err := ctx.ShouldBindJSON(&json); err != nil {
		errors := validator.FormatValidationError(err)
		httputils.NewErrorResponse(ctx, http.StatusBadRequest, validator.JoinErrorValidation(errors))
		return
	}

	user, err := handler.service.RegisterUser(json.Email, json.Name, json.Password)
	if err != nil {
		httputils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	httputils.NewSuccessResponse(ctx, http.StatusOK, "success", user)
}
