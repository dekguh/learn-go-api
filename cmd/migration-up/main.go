package main

import (
	"fmt"
	"log"

	configs "github.com/dekguh/learn-go-api/internal/pkg/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	configYaml := configs.LoadConfig()
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", configYaml.Database.Username, configYaml.Database.Password, configYaml.Database.Host, configYaml.Database.Port, configYaml.Database.Name)
	m, err := migrate.New(
		"file://db/migrations",
		dbUrl,
	)
	if err != nil {
		panic(err)
	}

	err = m.Up()
	if err != nil {
		panic(err)
	}

	log.Println("Migration up successfully")
}
