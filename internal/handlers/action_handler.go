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

func (a *Application) PostAction(c *gin.Context, request ActionType) {

	c.JSON(http.StatusOK, request)
}

func (a *Application) GetActions(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"actions": actions})
}
