package models

import (
	"fmt"
)

type StockInfo struct {
	CompanyName string
	DailyChange float64
	YTD         float64
}

func (s *StockInfo) GetYTD() string {
	prefix := "+"
	if s.YTD < 0 {
		prefix = ""
	}
	return fmt.Sprintf("%s%.2f", prefix, s.YTD) + "%"
}

func (s *StockInfo) GetToday() string {
	prefix := "+"
	if s.DailyChange < 0 {
		prefix = ""
	}
	return fmt.Sprintf("%s%.2f", prefix, s.DailyChange) + "%"
}
