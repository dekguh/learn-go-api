package http

import (
	handlerTodo "github.com/dekguh/learn-go-api/internal/api/http/handler/todo"
	handlerUser "github.com/dekguh/learn-go-api/internal/api/http/handler/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	handlerUser.UserRoutes(r, db)
	handlerTodo.TodoRoutes(r, db)

	return r
}
