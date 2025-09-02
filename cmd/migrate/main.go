package main

import (
	"github.com/dekguh/learn-go-api/internal/api/model"
	configs "github.com/dekguh/learn-go-api/internal/pkg/config"
	"github.com/dekguh/learn-go-api/internal/pkg/database"
	"gorm.io/gorm"
)

func main() {
	configYaml := configs.LoadConfig()
	GormDB := gorm.DB{}
	database.InitDatabase(configYaml, &GormDB)

	GormDB.AutoMigrate(&model.User{})
}
