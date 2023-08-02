package stocks

const (
	// Intraday function
	Intraday = "TIME_SERIES_INTRADAY"
	// Daily function
	Daily = "TIME_SERIES_DAILY"
	// DailyAdjusted function
	DailyAdjusted = "TIME_SERIES_DAILY_ADJUSTED"
	// Weekly function
	Weekly = "TIME_SERIES_WEEKLY"
	// Monthly function
	Monthly = "TIME_SERIES_MONTHLY"
	// Monthly adjusted func
	MonthlyAdjusted = "TIME_SERIES_MONTHLY_ADJUSTED"
)

type TimeSeriesData struct {
	MetaData             MetaData              `json:"Meta Data"`
	WeeklyAdjustedSeries map[string]WeeklyData `json:"Weekly Adjusted Time Series"`
}

type MetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	TimeZone      string `json:"4. Time Zone"`
}

type WeeklyData struct {
	Open          string `json:"1. open"`
	High          string `json:"2. high"`
	Low           string `json:"3. low"`
	Close         string `json:"4. close"`
	AdjustedClose string `json:"5. adjusted close"`
	Volume        string `json:"6. volume"`
	Dividend      string `json:"7. dividend amount"`
}

type QuoteResponse struct {
	Quote Quote `json:"Global Quote"`
}

type Quote struct {
	Symbol           string `json:"01. symbol"`
	Open             string `json:"02. open"`
	High             string `json:"03. high"`
	Low              string `json:"04. low"`
	Price            string `json:"05. price"`
	Volume           string `json:"06. volume"`
	LatestTradingDay string `json:"07. latest trading day"`
	PreviousClose    string `json:"08. previous close"`
	Change           string `json:"09. change"`
	ChangePercent    string `json:"10. change percent"`
}
