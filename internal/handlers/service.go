package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"splitflap-backend/internal/logger"
	"splitflap-backend/internal/models"
	"splitflap-backend/internal/stocks"
	"splitflap-backend/internal/utils"
	"splitflap-backend/pkg/fluent"
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
	Context              context.Context
	Sender               MessageSender
	Spotify              SpotifyClient
	Stocks               StocksClient
	State                DisplayState
	CurrentSplitflapText string
	Ws                   ws.WebSocket
	ExternalLCDDisplayIP string
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

func (a *Application) SendImage(url string) error {

	img, err := utils.ConvertUrlToImage(url)
	if err != nil {
		logger.Error().Err(err).Msg("failed to convert url to image")
		return err
	}

	bytes, err := json.Marshal(img)
	if err != nil {
		logger.Error().Err(err).Msg("failed to get image bytes")
		return err
	}

	err = fluent.Post("http://192.168.1.231:8080/image", bytes).
		OnSuccess(func(bytes []byte) error {
			return nil
		}).
		OnError(func(bytes []byte) error {
			return errors.New(string(bytes))
		}).
		Execute()

	if err != nil {
		logger.Error().Err(err).Msg("failed to send image")
		return err
	}

	return nil
}

type ImageRequest struct {
	Image [][]Color `json:"image"`
}

type Color struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
	A uint8 `json:"a"`
}
