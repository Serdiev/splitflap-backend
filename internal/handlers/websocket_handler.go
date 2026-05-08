package handlers

import (
	"splitflap-backend/internal/lcd_display"
	"splitflap-backend/internal/logger"

	"github.com/gin-gonic/gin"
)

type WebSocketRequest struct {
	Id SpotifyAccountId
}

func (a *Application) LcdWebsocketHandler(ctx *gin.Context, request WebSocketRequest) {
	spotifyClient, exists := a.SpotifyClients[request.Id]
	if !exists {
		logger.Warn().Str("id", string(request.Id)).Msg("WS handler: spotify client not found")
		ctx.AbortWithStatus(401)
		return
	}

	lcd, exists := a.LcdDisplays[request.Id]
	if !exists {
		a.LcdDisplays[request.Id] = lcd_display.NewLcdDisplay(string(request.Id), cfg.General.AllowedOrigins)
		logger.Warn().Str("id", string(request.Id)).Msg("WS handler: LCD not found, created new")
		ctx.AbortWithStatus(404)
		return
	}

	spotifyClient.RegisterHandler("update-lcd-image", lcd.HandleIsPlaying)

	logger.Info().Int("count", len(a.LcdDisplays)).Msg("WS handler: connecting LCD WebSocket")
	lcd.Ws.HandleWebSocket(ctx)
}
