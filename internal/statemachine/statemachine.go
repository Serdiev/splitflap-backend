package statemachine

import (
	"fmt"
	h "splitflap-backend/internal/handlers"
	"time"
)

func Initiate(app *h.Application) {
	var state = h.Idle

	ticker := time.NewTicker(time.Second)
	tick := 0
	for range ticker.C {
		tick++

		switch state {
		case h.Idle:
			if !isDivisibleBy(tick, 10) {
				continue
			}

			if handleSpotifyState(app) {
				continue
			}
			checkIfShouldSwapState(app)
			break
		case h.Stocks:
			getCurrentlyPlaying(app)
			break
		case h.Spotify:
			handleSpotifyState(app)
			break
		}
	}
}

func checkIfShouldSwapState(app *h.Application) {
	playing, err := app.Spotify.GetCurrentlyPlaying()
}

func handleSpotifyState(app *h.Application) bool {
	playing, err := app.Spotify.GetCurrentlyPlaying()
	if playing == nil {
		return false
	}

	app.SetSplitflapText("")
	app.SetState(h.Spotify)
	fmt.Println(playing)
	return true
}

func isDivisibleBy(tick int, n int) bool {
	return tick%n == 0
}
