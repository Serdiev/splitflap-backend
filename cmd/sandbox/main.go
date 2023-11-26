package main

import "splitflap-backend/internal/serial"

func main() {
	// c := context.Background()
	// app := handlers.CreateService(c)

	// statemachine.HandleStocksState(&app)

	sf := serial.NewSplitflap(nil)

	sf.SetText("abcdef", serial.ForceMovementNone)

	// s, err := app.Stocks.GetStockInfo(statemachine.TRACKED_STOCKS[0])

	// fmt.Println(s, err)
}
