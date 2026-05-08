package handlers

import (
	"fmt"
	"splitflap-backend/internal/lcd_display"

	"github.com/gin-gonic/gin"
)

type WebSocketRequest struct {
	Id SpotifyAccountId
}

func (a *Application) LcdWebsocketHandler(ctx *gin.Context, request WebSocketRequest) {
	spotifyClient, exists := a.SpotifyClients[request.Id]
	if !exists {
		fmt.Println("Spotify client not found or not logged in")
		ctx.AbortWithStatus(401)
		return
	}

	lcd, exists := a.LcdDisplays[request.Id]
	if !exists {
		a.LcdDisplays[request.Id] = lcd_display.NewLcdDisplay(string(request.Id), cfg.General.AllowedOrigins)
		fmt.Println("Lcd not found")
		ctx.AbortWithStatus(404)
		return
	}

	spotifyClient.RegisterHandler("update-lcd-image", lcd.HandleIsPlaying)

	fmt.Println("Connecting LCD WebSocket")
	fmt.Printf("Connecting LCD WebSocket: %d", len(a.LcdDisplays))
	lcd.Ws.HandleWebSocket(ctx)
}
