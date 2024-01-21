package statemachine

import (
	"fmt"
	h "splitflap-backend/internal/handlers"
	"splitflap-backend/internal/utils"
	"time"
)

var stateHandlers = map[h.DisplayState]func(*h.Application) bool{
	h.Idle:    idleState,
	h.Spotify: handleSpotifyState,
}

func Initiate(app *h.Application) {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		busy := stateHandlers[app.State](app)

		if !busy {
			handleSpotifyState(app)
		}
	}
}

func idleState(app *h.Application) bool {
	return false
}

var lastCheckedSpotify = time.Now()

func handleSpotifyState(app *h.Application) bool {
	if !app.Spotify.IsLoggedIn() {
		return false
	}

	if time.Since(lastCheckedSpotify) < time.Second*5 {
		return true
	}

	lastCheckedSpotify = time.Now()

	fmt.Println("checking")
	playing, _ := app.Spotify.GetCurrentlyPlaying()
	fmt.Println(playing)
	if playing == nil {
		app.ResetSplitFlapState()
		return false
	}

	text := utils.NewText()
	text.TopLeft(playing.Song)
	text.BottomLeft(playing.Artist)
	app.Sender.SendMessage(text.GetText())
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
