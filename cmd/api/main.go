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
	docs.SwaggerInfo.Host = configYaml.Openapi.Host
	docs.SwaggerInfo.BasePath = configYaml.Openapi.Basepath
	docs.SwaggerInfo.Version = configYaml.Openapi.Version
	docs.SwaggerInfo.Title = configYaml.Openapi.Title
	docs.SwaggerInfo.Description = configYaml.Openapi.Description

	log.Println("run on port: ", configYaml.Server.Port)

	router := http.SetupRouter()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + configYaml.Server.Port)
}
