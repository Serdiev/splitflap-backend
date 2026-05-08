package handlers

import (
	"context"
	"fmt"
	"strings"

	"splitflap-backend/internal/lcd_display"
	"splitflap-backend/internal/models"
	"splitflap-backend/internal/spotify"
	"splitflap-backend/internal/stocks"
	"splitflap-backend/internal/utils"

	gen "splitflap-backend/internal/generated"
	ws "splitflap-backend/internal/websocket"
)

type DisplayState int
type SpotifyAccountId string

const MainSpotifyAccountId SpotifyAccountId = "fredriks-spotify-account"

const (
	Idle DisplayState = iota
	Stocks
	Spotify
	Clock
)

type Application struct {
	Context                      context.Context
	Sender                       MessageSender
	SpotifyShouldUpdateSplitFlap bool
	Stocks                       StocksClient
	State                        DisplayState
	Ws                           ws.WebSocket
	ExternalLcdDisplayIpAddress  string
	CurrentSplitflapText         string
	CurrentDisplayedImageUrl     string
	SpotifyClients               map[SpotifyAccountId]*spotify.SpotifyClient
	LcdDisplays                  map[SpotifyAccountId]*lcd_display.Esp32LcdDisplay
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
}

// Checks if is in idle state
func (a *Application) IsIdleState() bool {
	return a.State == Idle
}

func (a *Application) isState(state DisplayState) bool {
	return a.State == state
}

func (a *Application) HandleSplitflapState(state *gen.SplitflapState) {
	var builder strings.Builder
	builder.Grow(len(state.Modules))
	for _, module := range state.Modules {
		builder.WriteByte(cfg.Splitflap.AlphabetESP32Order[module.FlapIndex])
	}
	text := builder.String()

	// set text so we can read whenever we load
	newText := utils.MapForReading(text)
	if a.CurrentSplitflapText == newText {
		return
	}

	a.CurrentSplitflapText = newText

	a.Ws.BroadcastMessageAsText(ws.ToBytes(CurrentTextResponse{CurrentText: a.CurrentSplitflapText}))
}

func (a *Application) SendIsPlayingTextToSplitflap(playing *models.SpotifyIsPlaying) {
	if (a.isState(Idle) || a.isState(Spotify)) && playing != nil {
		if !a.SpotifyShouldUpdateSplitFlap {
			return
		}

		a.SetState(Spotify)

		msg := getPlayingText(playing)
		_ = a.Sender.SendMessage(msg, "spotify playing")
	} else if a.isState(Spotify) && playing == nil {
		a.SetToIdleState("not playing spotify anymore")
	}
}

func getPlayingText(playing *models.SpotifyIsPlaying) string {
	text := utils.NewText()
	if len(playing.Song+" - "+playing.Artist) <= cfg.GetRowLength() {
		text.TopLeft(playing.Song + " - ")
		text.TopRight(playing.Artist)

		sliderText := BottomSlider(playing.PercentageLeft())
		text.BottomLeft(sliderText)
	} else {
		text.TopLeft(playing.Song)
		text.BottomLeft(playing.Artist)
	}

	return text.GetText()
}

func BottomSlider(percentage int) string {
	pct := percentage
	if percentage > 100 {
		pct = 100
	} else if percentage <= 0 {
		pct = 1
	}

	left := float32(cfg.GetRowLength()) * (float32(pct) / 100)
	finishedString := strings.Repeat("-", int(left))
	leftString := strings.Repeat("+", cfg.GetRowLength()-int(left))
	msg := fmt.Sprintf("%s%s", finishedString, leftString)
	return msg
}
