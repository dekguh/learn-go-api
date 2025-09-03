package http

import (
	handler "github.com/dekguh/learn-go-api/internal/api/http/handler/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	handler.UserRoutes(r, db)

	return r
}
