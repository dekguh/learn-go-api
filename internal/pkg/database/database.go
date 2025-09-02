package database

import (
	"fmt"

	configs "github.com/dekguh/learn-go-api/internal/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(config *configs.Config, DB *gorm.DB) {
	connectionStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Password, config.Database.Name,
	)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionStr,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	*DB = *db
}
