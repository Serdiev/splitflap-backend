package handlers

import (
	"net/http"

	"splitflap-backend/internal/logger"

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

	logger.Info().Str("ip", request.IpAddress).Msg("IP handler: received ESP32 IP")
	ctx.Status(http.StatusNoContent)
}
