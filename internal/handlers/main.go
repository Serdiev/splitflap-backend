package handlers

import (
	"context"
	config "splitflap-backend/configs"
	"splitflap-backend/internal/models"
	"splitflap-backend/internal/sender"
	"splitflap-backend/internal/spotify"
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
		sender:  sender.NewMQTTSender(cfg.MQTT_TOPIC),
		Spotify: spotify.NewNoopSpotifyClient(),
		State:   Idle,
	}
}

type Application struct {
	Context context.Context
	sender  MessageSender
	Spotify SpotifyClient
	State   DisplayState
}

type MessageSender interface {
	SendMessage(text string) error
}

type SpotifyClient interface {
	IsLoggedIn() bool
	GetCurrentlyPlaying() (*models.SpotifyIsPlaying, error)
}

// Sets text with correct length (inserting spaces or truncating)
func (a *Application) SetSplitflapText(text string) error {
	preparedText := text
	diff := len(text) - cfg.SPLITFLAP_MODULE_COUNT

	// truncate to module_count
	if diff > 0 {
		preparedText = text[:cfg.SPLITFLAP_MODULE_COUNT]
	}

	// adds spaces to module_count
	if diff < 0 {
		preparedText += strings.Repeat(" ", -1*diff)
	}

	return a.sender.SendMessage(preparedText)
}

// Sets text with correct length (inserting spaces or truncating)
func (a *Application) SetState(state DisplayState) {
	a.State = state
}

// Sets mode to idle and sets text to empty
func (a *Application) ResetSplitFlapState() {
	a.State = Idle
	a.sender.SendMessage(strings.Repeat(" ", cfg.SPLITFLAP_MODULE_COUNT))
}
