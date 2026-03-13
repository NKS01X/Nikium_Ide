package main

import (
	"log"

	"github.com/joho/godotenv"
	"nikium.com/api/cmd/server/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.POST("/login", auth.Login)
	log.Println("Starting API server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
