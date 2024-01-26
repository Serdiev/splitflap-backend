package handlers

import (
	"context"
	config "splitflap-backend/configs"
	"splitflap-backend/internal/models"
	"splitflap-backend/internal/sender"
	"splitflap-backend/internal/spotify"
	"splitflap-backend/internal/stocks"
	"strings"
)

var cfg = config.New()
var App Application

type DisplayState int

const (
	Idle DisplayState = iota
	Stocks
	Spotify
	Clock
)

func CreateService(c context.Context) Application {
	return Application{
		Context: c,
		Sender:  GetSender(),
		Spotify: spotify.NewNoopSpotifyClient(), // gets replaced with real client once we login to spotify
		Stocks:  stocks.NewAvanzaClient(),
		State:   Idle,
	}
}

type Application struct {
	Context context.Context
	Sender  MessageSender
	Spotify SpotifyClient
	Stocks  StocksClient
	State   DisplayState
}

func GetSender() MessageSender {
	usbSender := sender.NewUsbSerialSender()
	if usbSender != nil {
		return usbSender
	}

	return sender.NewNoopSender()
}

type MessageSender interface {
	SendMessage(text string) error
}

type StocksClient interface {
	GetStockInfo(s stocks.Stock) (*models.StockInfo, error)
}

// Sets text with correct length (inserting spaces or truncating)
func (a *Application) SetSplitflapText(text string) error {
	preparedText := text
	diff := len(text) - cfg.Splitflap.ModuleCount

	// truncate to module_count
	if diff > 0 {
		preparedText = text[:cfg.Splitflap.ModuleCount]
	}

	// adds spaces to module_count
	if diff < 0 {
		preparedText += strings.Repeat(" ", -1*diff)
	}

	return a.Sender.SendMessage(preparedText)
}

// Sets text with correct length (inserting spaces or truncating)
func (a *Application) SetState(state DisplayState) {
	a.State = state
}

// Sets mode to idle and sets text to empty
func (a *Application) SetToIdleState() {
	a.State = Idle
	a.SetSplitflapText(strings.Repeat("a", cfg.Splitflap.ModuleCount))
}

// Sets text with correct length (inserting spaces or truncating)
func (a *Application) IsIdleState() bool {
	return a.State == Idle
}
