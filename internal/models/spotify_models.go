package models

type SpotifyIsPlaying struct {
	Song        string
	Artist      string
	SecondsLeft int
	TimeLeft    string
}

type StockInfo struct {
	Tick        string
	Company     string
	DailyChange string
	YTD         string
}
