package statemachine

import (
	"fmt"
	h "splitflap-backend/internal/handlers"
	"time"
)

var stateHandlers = map[h.DisplayState]func(*h.Application) bool{
	h.Idle:    idleState,
	h.Spotify: initSpotifyStateHandler,
}

func Initiate(app *h.Application) {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		if app.IsIdleState() {
			continue
		}

		for _, swapHandlerFunc := range stateHandlers {
			swapState := swapHandlerFunc(app)
			if swapState {
				break
			}
		}
	}
}

func idleState(app *h.Application) bool {
	app.SetToIdleState()
	return false
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
