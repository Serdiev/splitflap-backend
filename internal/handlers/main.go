package handlers

import (
	"context"
	config "splitflap-backend/configs"
	"splitflap-backend/internal/sender"
	"splitflap-backend/internal/spotify"
	"splitflap-backend/internal/stocks"
)

var cfg = config.New()
var App Application

func CreateService(c context.Context) *Application {
	return &Application{
		Context: c,
		Sender:  GetSender(),
		Spotify: spotify.NewNoopSpotifyClient(), // gets replaced with real client once we login to spotify
		Stocks:  stocks.NewAvanzaClient(),
		State:   Idle,
	}
}

func GetSender() MessageSender {
	usbSender := sender.NewUsbSerialSender()
	if usbSender != nil {
		return usbSender
	}

	return sender.NewNoopSender()
}
