package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CurrentTextResponse struct {
	CurrentText string `json:"currentText"`
}

func (a *Application) GetCurrentMessage(c *gin.Context) {
	c.JSON(http.StatusOK, CurrentTextResponse{CurrentText: a.CurrentSplitflapText})
}
