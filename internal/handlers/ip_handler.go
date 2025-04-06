package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Boop struct {
	IpAddress string `json:"ip_address"`
	Secret    string `json:"secret"`
}

func (a *Application) UpdateESP32IPAddress(ctx *gin.Context, request Boop) {
	if request.Secret != "boopedo" {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	fmt.Println("Received ip:", request.IpAddress)
	a.ExternalLcdDisplayIpAddress = request.IpAddress
	ctx.Status(http.StatusNoContent)
}
