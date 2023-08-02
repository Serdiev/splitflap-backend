package models

type SpotifyIsPlaying struct {
	Song        string
	Artist      string
	SecondsLeft int
	TimeLeft    string
}

type StockInfo struct {
	CompanyName string
	DailyChange float64
	YTD         float64
}
