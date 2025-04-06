package handlers

import (
	"context"
	"splitflap-backend/internal/models"
	"splitflap-backend/internal/stocks"
	"splitflap-backend/internal/utils"
	"strings"

	gen "splitflap-backend/internal/generated"
	ws "splitflap-backend/internal/websocket"
)

type DisplayState int

const (
	Idle DisplayState = iota
	Stocks
	Spotify
	Clock
)

type Application struct {
	Context                     context.Context
	Sender                      MessageSender
	Spotify                     SpotifyClient
	Stocks                      StocksClient
	State                       DisplayState
	Ws                          ws.WebSocket
	ExternalLcdDisplayIpAddress string
	CurrentSplitflapText        string
	CurrentDisplayedImageUrl    string
}

type MessageSender interface {
	SendMessage(text string, sentFrom string) error
}

type StocksClient interface {
	GetStockInfo(s stocks.Stock) (*models.StockInfo, error)
	GetOHLCData(s stocks.Stock) (*models.OHLCData, error)
}

// Sets text with correct length (inserting spaces or truncating)
func (a *Application) SetState(state DisplayState) {
	a.State = state
}

// Sets mode to idle and sets text to empty
func (a *Application) SetToIdleState(sentBy string) {
	a.State = Idle
	_ = a.Sender.SendMessage(strings.Repeat(" ", cfg.Splitflap.ModuleCount), sentBy)
	_ = a.SendImage(utils.EmptyImage())
}

// Sets text with correct length (inserting spaces or truncating)
func (a *Application) IsIdleState() bool {
	return a.State == Idle
}

func (a *Application) HandleSplitflapState(state *gen.SplitflapState) {
	text := ""
	for _, module := range state.Modules {
		text += string(cfg.Splitflap.AlphabetESP32Order[module.FlapIndex])
	}

	// set text so we can read whenever we load
	newText := utils.MapForReading(text)
	if a.CurrentSplitflapText == newText {
		return
	}

	a.CurrentSplitflapText = newText

	a.Ws.BroadcastMessage(ws.ToBytes(CurrentTextResponse{CurrentText: a.CurrentSplitflapText}))
}
