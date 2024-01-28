package statemachine

import (
	"fmt"
	"splitflap-backend/internal/stocks"
	"splitflap-backend/internal/utils"
	"time"
)

var hasDisplayedToday = false

func (s *StateHandler) initAvanzaHandler() bool {
	// only display once per day
	if hasDisplayedToday {
		return false
	}

	// automatically print
	if time.Now().Before(utils.GetTodayAt(18, 0)) {
		return false
	}

	hasDisplayedToday = true
	go s.handleStocksState()
	return true
}

func (s *StateHandler) handleStocksState() bool {
	for _, stock := range stocks.TRACKED_STOCKS {
		info, err := s.App.Stocks.GetStockInfo(stock)
		fmt.Println(info, err)
		if err != nil {
			fmt.Println(err)
		}

		text := utils.NewText()
		text.TopLeft(info.CompanyName)

		// text.BottomLeft(info.YTD)
	}

	return true
}
