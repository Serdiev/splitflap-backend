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

func (s *StateHandler) initSpotifyStateHandler() bool {
	if !s.App.Spotify.IsLoggedIn() || !s.App.Spotify.ShouldUpdateSplitFlap() {
		return false
	}

	if time.Since(lastCheckedSpotify) < time.Second*5 {
		return false
	}

	lastCheckedSpotify = time.Now()

	playing, err := s.App.Spotify.GetCurrentlyPlaying()
	if err != nil || playing == nil {
		return false
	}

	// set state to handle spotify and initiate spotify handler
	s.App.SetState(h.Spotify)
	go s.handleSpotifyState()
	return true
}

var lastSentImg = ""

func (s *StateHandler) handleSpotifyState() {
	ticker := time.NewTicker(5 * time.Second)

	for range ticker.C {
		if !s.App.Spotify.ShouldUpdateSplitFlap() {
			s.App.SetToIdleState("not playing spotify anymore")
			return
		}
		if s.App.State != h.Spotify {
			s.App.SetToIdleState("something took precedence over spotify")
			return
		}

		playing, _ := s.App.Spotify.GetCurrentlyPlaying()
		if playing == nil {
			s.App.SetToIdleState("playing nil")
			return
		}

		msg := getPlayingText(playing)
		s.App.Sender.SendMessage(msg, "spotify playing")

		if lastSentImg != playing.Image64PixelUrl {
			s.App.SendImage(playing.Image64PixelUrl)
			// lastSentImg = playing.Image64PixelUrl
		}
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
