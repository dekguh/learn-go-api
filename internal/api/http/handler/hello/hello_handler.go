package hello

import "github.com/gin-gonic/gin"

type HelloHandler struct{}

func (h *HelloHandler) HelloRoutes(r *gin.Engine) {
	group := r.Group("/hello")
	{
		group.GET("/world", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello, World!",
			})
		})
	}
}
