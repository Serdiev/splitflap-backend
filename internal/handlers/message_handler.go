package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SendMessageRequest struct {
	Message string `json:"message"`
}

func (a *Application) SendMessage(c *gin.Context, request SendMessageRequest) {
	a.Sender.SendMessage(request.Message, "sendmessage handler")
	c.JSON(http.StatusOK, request)
}
