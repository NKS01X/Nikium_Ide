package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"nikium.com/backend/internal/auth"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/api/auth/login", auth.Login)
	r.POST("/api/auth/logout", auth.Logout)

	log.Println("Starting backend on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
