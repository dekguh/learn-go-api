package main

import (
	"log"

	"github.com/dekguh/learn-go-api/cmd/api/docs"
	http "github.com/dekguh/learn-go-api/internal/api/http/handler"
	configs "github.com/dekguh/learn-go-api/internal/pkg/config"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	configYaml := configs.LoadConfig()
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "Hello World"
	docs.SwaggerInfo.Description = "Hello World docs"

	log.Println("run on port: ", configYaml.Server.Port)

	router := http.SetupRouter()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + configYaml.Server.Port)
}
