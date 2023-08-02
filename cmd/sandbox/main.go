package main

import (
	"context"
	"splitflap-backend/internal/handlers"
	"splitflap-backend/internal/statemachine"
)

func main() {
	c := context.Background()
	app := handlers.CreateService(c)

	statemachine.HandleStocksState(&app)
	// s, err := app.Stocks.GetStockInfo(statemachine.TRACKED_STOCKS[0])

	// fmt.Println(s, err)
}
