package main

import (
	"log"

	configs "github.com/dekguh/learn-go-api/internal/api/config"
	"github.com/gin-gonic/gin"
)

func main() {
	configYaml := configs.LoadConfig()

	log.Println("config:", configYaml.Application)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.Run(":8080")
}
