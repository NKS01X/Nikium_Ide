package auth

import (
	"fmt"

	"nikium.com/api/cmd/server/utils"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username        string `json:"username"`
	Password_hashed string `json:"password"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(req.Username)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate token"})
		return
	}
	c.JSON(200, gin.H{"token": token})
	fmt.Println("Token created")
	fmt.Println(token)
}
