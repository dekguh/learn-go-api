package main

import (
	"log"

	"github.com/dekguh/learn-go-api/configs"
	"github.com/gin-gonic/gin"
)

func main() {
	configYaml, err := configs.LoadConfig("")
	if err != nil {
		log.Fatalln("failed to load config:", err)
	}

	log.Println("config:", configYaml)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.Run(":8080")
}
