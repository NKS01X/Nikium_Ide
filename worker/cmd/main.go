package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"nikium.com/worker/internal/runner"
)

type CodeRunnerRequest struct {
	ID       uuid.UUID `json:"id"`
	UserId   uuid.UUID `json:"userId"`
	Code     string    `json:"code"`
	Input    string    `json:"input"`
	Language string    `json:"language"`
}

type CodeRunnerResponse struct {
	ID       uuid.UUID `json:"id"`
	UserId   uuid.UUID `json:"userId"`
	Output   string    `json:"output"`
	Error    string    `json:"error"`
	ExitCode int       `json:"exitCode"`
	Duration string    `json:"duration"`
}

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.POST("/run", func(c *gin.Context) {
		var req CodeRunnerRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		start := time.Now()

		output, exitCode, err := runner.RunCode(runner.CodeRunnerRequest{
			ID:       req.ID,
			UserId:   req.UserId,
			Code:     req.Code,
			Input:    req.Input,
			Language: req.Language,
		})
		duration := time.Since(start)

		resp := CodeRunnerResponse{
			ID:       req.ID,
			UserId:   req.UserId,
			Output:   output,
			ExitCode: exitCode,
			Duration: duration.String(),
		}

		if err != nil {
			resp.Error = err.Error()
		}

		c.JSON(http.StatusOK, resp)
	})

	log.Println("Starting worker on :8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
