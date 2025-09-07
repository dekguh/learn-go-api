package http

import (
	"os"
	"time"

	handlerTodo "github.com/dekguh/learn-go-api/internal/api/http/handler/todo"
	handlerUser "github.com/dekguh/learn-go-api/internal/api/http/handler/user"
	"github.com/dekguh/learn-go-api/internal/api/http/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("CORS_HOST")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == os.Getenv("CORS_HOST")
		},
		MaxAge: 12 * time.Hour,
	}))

	r.Use(middleware.EtagMiddleware())

	handlerUser.UserRoutes(r, db)
	handlerTodo.TodoRoutes(r, db)

	return r
}
