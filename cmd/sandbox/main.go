package main

import (
	"fmt"
	"splitflap-backend/internal/usb_serial"

	"go.bug.st/serial"
)

func main() {
	// c := context.Background()
	// app := handlers.CreateService(c)

	// statemachine.HandleStocksState(&app)

	a, b := serial.GetPortsList()

	fmt.Println(a, b)
	sf := usb_serial.NewSplitflap(nil)

	sf.SetText("abcdef", usb_serial.ForceMovementNone)

	// s, err := app.Stocks.GetStockInfo(statemachine.TRACKED_STOCKS[0])

	// fmt.Println(s, err)
}
