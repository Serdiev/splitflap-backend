package handlers

import (
	"context"
	"splitflap-backend/internal/lcd_display"
	"splitflap-backend/internal/sender"
	"splitflap-backend/internal/spotify"
	"splitflap-backend/internal/stocks"

	config "splitflap-backend/configs"
	gen "splitflap-backend/internal/generated"
	ws "splitflap-backend/internal/websocket"
)

var cfg = config.New()
var App Application

func CreateService(c context.Context) *Application {
	a := &Application{
		Context:                      c,
		Stocks:                       stocks.NewAvanzaClient(),
		State:                        Idle,
		Ws:                           *ws.NewWebsocket(),
		LcdDisplays:                  map[SpotifyAccountId]*lcd_display.Esp32LcdDisplay{},
		SpotifyClients:               map[SpotifyAccountId]*spotify.SpotifyClient{},
		SpotifyShouldUpdateSplitFlap: true,
	}

	a.Sender = GetSender(a.HandleSplitflapState)
	return a
}

func GetSender(handleState func(state *gen.SplitflapState)) MessageSender {
	usbSender := sender.NewUsbSerialSender(handleState)
	if usbSender != nil {
		return usbSender
	}

	return sender.NewNoopSender()
}
