package statemachine

import (
	"fmt"
	config "splitflap-backend/configs"
	h "splitflap-backend/internal/handlers"
	"splitflap-backend/internal/models"
	"splitflap-backend/internal/utils"
	"strings"
	"time"
)

var cfg = config.New()
var lastCheckedSpotify = time.Now()

func initSpotifyStateHandler(app *h.Application) bool {
	if !app.Spotify.IsLoggedIn() {
		return false
	}

	if time.Since(lastCheckedSpotify) < time.Second*5 {
		return false
	}

	lastCheckedSpotify = time.Now()

	playing, err := app.Spotify.GetCurrentlyPlaying()
	if err != nil || playing == nil {
		return false
	}

	// set state to handle spotify and initiate spotify handler
	app.SetState(h.Spotify)
	go handleSpotifyState(app)
	return true
}

func handleSpotifyState(app *h.Application) {
	ticker := time.NewTicker(5 * time.Second)

	for range ticker.C {
		playing, _ := app.Spotify.GetCurrentlyPlaying()
		if playing == nil {
			app.SetState(h.Idle)
			return
		}

		app.Sender.SendMessage(getPlayingText(playing))
	}
}

func getPlayingText(playing *models.SpotifyIsPlaying) string {
	text := utils.NewText()
	if len(playing.Song+" - "+playing.Artist) <= cfg.GetRowLength() {
		text.TopLeft(playing.Song + " - ")
		text.TopRight(playing.Artist)

		// text.BottomLeft(BottomSlider())
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

	passed := 0
	left := float32(cfg.GetRowLength()) * (float32(pct) / 100)

	finishedString := strings.Repeat("%", int(left))
	leftString := strings.Repeat("%", cfg.GetRowLength()-passed)
	msg := fmt.Sprintf("%s%s", finishedString, leftString)
	return msg
}
