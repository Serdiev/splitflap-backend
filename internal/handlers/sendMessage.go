package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SendMessageRequest struct {
	Message string `json:"message"`
}

func (a Application) SendMessage(c *gin.Context) {
	var request = SendMessageRequest{}
	// Bind the JSON data from the request body to the 'playing' struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a.sender.SendMessage(request.Message)

	c.JSON(http.StatusOK, request)
}
