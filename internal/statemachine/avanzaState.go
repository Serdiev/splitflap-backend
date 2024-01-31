package statemachine

import (
	"splitflap-backend/internal/stocks"
	"splitflap-backend/internal/utils"
	"time"
)

var nextRun = utils.GetTodayAt(18, 0).Add(time.Hour * 24)

func (s *StateHandler) initAvanzaHandler() bool {
	// only display once per day (and not within 1 hr of starting app)
	if time.Now().Before(nextRun) {
		return false
	}

	nextRun = utils.GetTodayAt(18, 0).Add(time.Hour * 24)
	go s.handleStocksState()
	return true
}

func (s *StateHandler) handleStocksState() {
	for _, stock := range stocks.TRACKED_STOCKS {
		info, err := s.App.Stocks.GetStockInfo(stock)
		if err != nil {
			return
		}

		text := utils.NewText()
		text.TopLeft(info.CompanyName)
		text.BottomRight(info.GetToday())

		s.App.Sender.SendMessage(text.GetText(), "set stock")

		time.Sleep(10 * time.Second)
	}

	s.App.SetToIdleState("stock ending")
}
