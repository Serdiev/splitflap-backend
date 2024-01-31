package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Action string

const (
	ShowStocks   Action = "ShowStocks"
	DisplayClock Action = "DisplayClock"
	// DisplayClock Action = "DisplayClock"
	// DisplayClock Action = "DisplayClock"
)

var actions = map[Action]ActionType{
	ShowStocks: {
		Label: "Show stocks",
	},
}

type ActionType struct {
	Label string `json:"label"`
}

func (a *Application) ExecuteAction(c *gin.Context) {
	var request = SendMessageRequest{}
	// Bind the JSON data from the request body to the 'playing' struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a.Sender.SendMessage(request.Message, "Execute action")

	c.JSON(http.StatusOK, request)
}

func (a *Application) GetActions(c *gin.Context) {
	var request = SendMessageRequest{}
	// Bind the JSON data from the request body to the 'playing' struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a.Sender.SendMessage(request.Message, "get actions")

	c.JSON(http.StatusOK, request)
}
