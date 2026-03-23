package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"nikium.com/backend/internal/auth"
)

type CodeSubmissionRequest struct {
	Code     string `json:"code" binding:"required"`
	Input    string `json:"input"`
	Language string `json:"language"`
}

type CodeSubmissionResponse struct {
	ID      uuid.UUID `json:"id"`
	Status  string    `json:"status"`
	Message string    `json:"message"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.POST("/api/auth/signup", auth.Signup)
	r.POST("/api/auth/login", auth.Login)
	r.POST("/api/auth/logout", auth.Logout)

	r.POST("/api/code/submit", func(c *gin.Context) {
		var req CodeSubmissionRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		submissionID := uuid.New()
		log.Printf("Code submission received: %s", submissionID)

		c.JSON(http.StatusAccepted, CodeSubmissionResponse{
			ID:      submissionID,
			Status:  "queued",
			Message: "Code submitted for execution",
		})
	})

	r.GET("/api/code/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":   "completed",
			"output":   "Hello, World!",
			"exitCode": 0,
		})
	})

	log.Println("Starting backend on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
