package main

import (
	"log"

	http "github.com/dekguh/learn-go-api/internal/api/http/handler"
	configs "github.com/dekguh/learn-go-api/internal/pkg/config"
)

func main() {
	configYaml := configs.LoadConfig()

	log.Println("run on port: ", configYaml.Server.Port)

	router := http.SetupRouter()
	router.Run(":" + configYaml.Server.Port)
}
