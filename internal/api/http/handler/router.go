package http

import (
	"github.com/dekguh/learn-go-api/internal/api/http/handler/hello"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	helloHandler := hello.HelloHandler{}
	helloHandler.HelloRoutes(r)

	return r
}
