package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SendMessageRequest struct {
	Message string `json:"message"`
}

func (a *Application) SendMessage(c *gin.Context, request SendMessageRequest) {
	_ = a.Sender.SendMessage(request.Message, "sendmessage handler")
	c.JSON(http.StatusOK, request)
}

type CurrentTextResponse struct {
	CurrentText string `json:"currentText"`
}

func (a *Application) GetCurrentMessage(c *gin.Context) {
	c.JSON(http.StatusOK, CurrentTextResponse{CurrentText: a.CurrentSplitflapText})
}
