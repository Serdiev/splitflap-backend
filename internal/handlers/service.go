package handlers

import (
	"context"
	"splitflap-backend/internal/models"
	"splitflap-backend/internal/stocks"
	"strings"
)

type DisplayState int

const (
	Idle DisplayState = iota
	Stocks
	Spotify
	Clock
)

type Application struct {
	Context context.Context
	Sender  MessageSender
	Spotify SpotifyClient
	Stocks  StocksClient
	State   DisplayState
}

type MessageSender interface {
	SendMessage(text string, sentFrom string) error
}

type StocksClient interface {
	GetStockInfo(s stocks.Stock) (*models.StockInfo, error)
}

// Sets text with correct length (inserting spaces or truncating)
func (a *Application) SetState(state DisplayState) {
	a.State = state
}

// Sets mode to idle and sets text to empty
func (a *Application) SetToIdleState(sentBy string) {
	a.State = Idle
	a.Sender.SendMessage(strings.Repeat(" ", cfg.Splitflap.ModuleCount), sentBy)
}

// Sets text with correct length (inserting spaces or truncating)
func (a *Application) IsIdleState() bool {
	return a.State == Idle
}
