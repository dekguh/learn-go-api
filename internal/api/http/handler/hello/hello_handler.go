package hello

import (
	httputils "github.com/dekguh/learn-go-api/internal/pkg/utils"
	"github.com/gin-gonic/gin"
)

type HelloHandler struct{}

func (h *HelloHandler) HelloRoutes(r *gin.Engine) {
	group := r.Group("/hello")
	{
		group.GET("/world", func(c *gin.Context) {
			httputils.NewSuccessResponse(
				c,
				200,
				"Hello, World!",
				map[string]interface{}{"success": true},
			)
		})
	}
}
