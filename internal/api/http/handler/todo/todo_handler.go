package handler

import (
	"net/http"
	"strconv"

	"github.com/dekguh/learn-go-api/internal/api/http/middleware"
	"github.com/dekguh/learn-go-api/internal/api/model"
	"github.com/dekguh/learn-go-api/internal/api/repository"
	"github.com/dekguh/learn-go-api/internal/api/service"
	dbutils "github.com/dekguh/learn-go-api/internal/pkg/utils/database"
	httputils "github.com/dekguh/learn-go-api/internal/pkg/utils/http"
	"github.com/dekguh/learn-go-api/internal/pkg/validator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TodoHandler struct {
	service service.TodoService
}

type CreateTodoReq struct {
	Title       string `json:"title" binding:"required,min=3,max=128"`
	Description string `json:"description" binding:"required,min=3,max=255"`
}

func TodoRoutes(r *gin.Engine, db *gorm.DB) {
	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := NewTodoHandler(todoService)

	groupTodo := r.Group("/todos", middleware.Authentication())
	{
		groupTodo.POST("/create", func(ctx *gin.Context) {
			todoHandler.CreateTodo(ctx, db)
		})
		groupTodo.GET("/search", func(ctx *gin.Context) {
			todoHandler.FindAllTodos(ctx, db)
		})
		groupTodo.DELETE("/delete/:id", func(ctx *gin.Context) {
			todoHandler.DeleteTodoById(ctx, db)
		})
	}
}

func NewTodoHandler(service service.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

// @Summary Create todo
// @Description Create todo
// @Tags Todos
// @Accept json
// @Produce json
// @Param createTodoReq body CreateTodoReq true "Create todo request"
// @Success 200 {object} httputils.SuccessResponse{data=model.Todo}
// @Failure 400 {object} httputils.ErrorResponse
// @Router /todos/create [post]
func (handler *TodoHandler) CreateTodo(ctx *gin.Context, db *gorm.DB) {
	var json CreateTodoReq
	dbutils.SetCurrentUserId(db, ctx.GetUint("user_id"))
	if err := ctx.ShouldBindJSON(&json); err != nil {
		errors := validator.FormatValidationError(err)
		httputils.NewErrorResponse(ctx, http.StatusBadRequest, validator.JoinErrorValidation(errors))
		return
	}

	todo, err := handler.service.CreateTodo(&model.Todo{
		Title:       json.Title,
		Description: json.Description,
		Status:      "NOT_STARTED",
		UserID:      ctx.GetUint("user_id"),
	})
	if err != nil {
		httputils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	httputils.NewSuccessResponse(ctx, http.StatusOK, "success", todo)
}

// @Summary Find all todos
// @Description Find all todos
// @Tags Todos
// @Accept json
// @Produce json
// @Success 200 {object} httputils.SuccessResponse{data=[]model.Todo}
// @Failure 400 {object} httputils.ErrorResponse
// @Router /todos/search [get]
func (handler *TodoHandler) FindAllTodos(ctx *gin.Context, db *gorm.DB) {
	todos, err := handler.service.FindAllTodos()
	dbutils.SetCurrentUserId(db, ctx.GetUint("user_id"))
	if err != nil {
		httputils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	httputils.NewSuccessResponse(ctx, http.StatusOK, "success", todos)
}

// @Summary Delete todo by id
// @Description Delete todo by id
// @Tags Todos
// @Accept json
// @Produce json
// @Param id path string true "Todo id"
// @Success 200 {object} httputils.SuccessResponse{data=nil}
// @Failure 400 {object} httputils.ErrorResponse
// @Router /todos/delete/{id} [delete]
func (handler *TodoHandler) DeleteTodoById(ctx *gin.Context, db *gorm.DB) {
	dbutils.SetCurrentUserId(db, ctx.GetUint("user_id"))
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		httputils.NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID format")
		return
	}
	uintId := uint(id)

	dbutils.SetCurrentUserId(db, ctx.GetUint("user_id"))
	if err := handler.service.DeleteTodoById(uintId); err != nil {
		httputils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	httputils.NewSuccessResponse(ctx, http.StatusOK, "Success delete todo", nil)
}
