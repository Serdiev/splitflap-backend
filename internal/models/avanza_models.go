package models

import (
	"fmt"
	"time"
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

type OHLCData struct {
	CompanyName          string
	From                 time.Time
	To                   time.Time
	PreviousClosingPrice float64
	OHLC                 []OHLC
}

type OHLC struct {
	Timestamp         time.Time
	Open              float64
	Close             float64
	High              float64
	Low               float64
	TotalVolumeTraded int64
}

func (o *OHLCData) GetMaxValue() float64 {
	highest := o.OHLC[0].Close
	for _, ohlc := range o.OHLC {
		if ohlc.Close > highest {
			highest = ohlc.Close
		}
	}
	return highest
}

func (o *OHLCData) GetMinValue() float64 {
	lowest := o.OHLC[0].Close
	for _, ohlc := range o.OHLC {
		if ohlc.Close < lowest {
			lowest = ohlc.Close
		}
	}
	return lowest
}

func (o *OHLCData) StartValue() float64 {
	return o.OHLC[0].Close
}

func (o *OHLCData) EndValue() float64 {
	return o.OHLC[len(o.OHLC)-1].Close
}
