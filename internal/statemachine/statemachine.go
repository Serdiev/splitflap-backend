package statemachine

import (
	"fmt"
	config "splitflap-backend/configs"
	h "splitflap-backend/internal/handlers"
	"time"
)

var cfg = config.New()

func Initiate(app *h.Application) {
	tick := 0
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		switch app.State {
		case h.Idle:
			idleState(tick, app)
			break
		case h.Stocks:
			// stockStateHandler will go through all the stocks setup in TRACKED_STOCKS
			break
		case h.Spotify:
			handleSpotifyState(app)
			break
		}
	}

}

func idleState(tick int, app *h.Application) {
	tick = (tick + 1) % 10

	// only check if we should swap to spotifyState every 10 seconds
	if tick != 0 {
		return
	}

	if handleSpotifyState(app) {
		return
	}
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

func HandleStocksState(app *h.Application) bool {
	ticker := time.NewTicker(time.Second * 1)

	i := 0
	for range ticker.C {
		stock := TRACKED_STOCKS[i]
		info, err := app.Stocks.GetStockInfo(stock)
		fmt.Println(info, err)
		if err != nil {
			fmt.Println(err)
		}

		if i >= len(TRACKED_STOCKS) {
			break
		}

		i++
	}

	return true
}

func isDivisibleBy(tick int, n int) bool {
	return tick%n == 0
}
