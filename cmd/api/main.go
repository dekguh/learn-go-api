package main

import (
	"github.com/dekguh/learn-go-api/cmd/api/docs"
	http "github.com/dekguh/learn-go-api/internal/api/http/handler"
	configs "github.com/dekguh/learn-go-api/internal/pkg/config"
	"github.com/dekguh/learn-go-api/internal/pkg/database"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func main() {
	configYaml := configs.LoadConfig()
	docs.SwaggerInfo.Host = configYaml.Openapi.Host
	docs.SwaggerInfo.BasePath = configYaml.Openapi.Basepath
	docs.SwaggerInfo.Version = configYaml.Openapi.Version
	docs.SwaggerInfo.Title = configYaml.Openapi.Title
	docs.SwaggerInfo.Description = configYaml.Openapi.Description

	GormDB := gorm.DB{}
	database.InitDatabase(configYaml, &GormDB)
	DBsql, err := GormDB.DB()
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	defer DBsql.Close()

	router := http.SetupRouter(&GormDB)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + configYaml.Server.Port)
}
