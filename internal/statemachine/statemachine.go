package statemachine

import (
	"fmt"
	h "splitflap-backend/internal/handlers"
	"time"
)

func Initiate(app *h.Application) {
	ticker := time.NewTicker(time.Second)
	tick := 0
	for range ticker.C {
		tick++

		switch app.State {
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
			handleStocksState(app)
			break
		case h.Spotify:
			handleSpotifyState(app)
			break
		}
	}
}

func checkIfShouldSwapState(app *h.Application) {
	playing, err := app.Spotify.GetCurrentlyPlaying()
	fmt.Println(playing, err)
	fmt.Println("not playing anything?")
}

func handleSpotifyState(app *h.Application) bool {
	playing, _ := app.Spotify.GetCurrentlyPlaying()
	if playing == nil {
		app.ResetSplitFlapState()
		return false
	}

	app.SetSplitflapText("")
	app.SetState(h.Spotify)
	return true
}

func handleStocksState(app *h.Application) bool {
	app.Stocks.GetStockInfo("msft")
	return true
}

func isDivisibleBy(tick int, n int) bool {
	return tick%n == 0
}
