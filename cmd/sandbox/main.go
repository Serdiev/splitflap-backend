package main

import (
	"context"
	"splitflap-backend/internal/handlers"
)

func main() {
	c := context.Background()
	app := handlers.CreateService(c)

	app.Stocks.GetStockInfo("AAPL")
}
