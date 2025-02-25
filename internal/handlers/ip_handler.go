package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Boop struct {
	IpAddress string `json:"ip_address"`
	Secret    string `json:"secret"`
}

func (a *Application) UpdateESP32IPAddress(ctx *gin.Context, request Boop) {
	if request.Secret != "secret" {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	a.ExternalLCDDisplayIP = request.IpAddress

	ctx.Status(http.StatusNoContent)
}
