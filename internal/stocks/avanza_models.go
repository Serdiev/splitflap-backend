package stocks

type AvanzaResponse struct {
	OrderbookID             string                  `json:"orderbookId"`
	Name                    string                  `json:"name"`
	Isin                    string                  `json:"isin"`
	InstrumentID            string                  `json:"instrumentId"`
	Sectors                 []Sectors               `json:"sectors"`
	Tradable                string                  `json:"tradable"`
	Listing                 Listing                 `json:"listing"`
	HistoricalClosingPrices HistoricalClosingPrices `json:"historicalClosingPrices"`
	KeyIndicators           KeyIndicators           `json:"keyIndicators"`
	Quote                   AvanzaQuote             `json:"quote"`
	Type                    string                  `json:"type"`
}
type Sectors struct {
	SectorID   string `json:"sectorId"`
	SectorName string `json:"sectorName"`
}
type Listing struct {
	ShortName             string `json:"shortName"`
	TickerSymbol          string `json:"tickerSymbol"`
	CountryCode           string `json:"countryCode"`
	Currency              string `json:"currency"`
	MarketPlaceCode       string `json:"marketPlaceCode"`
	MarketPlaceName       string `json:"marketPlaceName"`
	MarketListName        string `json:"marketListName"`
	TickSizeListID        string `json:"tickSizeListId"`
	MarketTradesAvailable bool   `json:"marketTradesAvailable"`
}
type HistoricalClosingPrices struct {
	OneDay      float64 `json:"oneDay"`
	OneWeek     float64 `json:"oneWeek"`
	OneMonth    float64 `json:"oneMonth"`
	ThreeMonths float64 `json:"threeMonths"`
	StartOfYear float64 `json:"startOfYear"`
	OneYear     float64 `json:"oneYear"`
	ThreeYears  float64 `json:"threeYears"`
	FiveYears   float64 `json:"fiveYears"`
	TenYears    float64 `json:"tenYears"`
	Start       float64 `json:"start"`
	StartDate   string  `json:"startDate"`
}
type MarketCapital struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}
type EquityPerShare struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}
type TurnoverPerShare struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}
type EarningsPerShare struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}
type Dividend struct {
	ExDate       string  `json:"exDate"`
	Amount       float64 `json:"amount"`
	CurrencyCode string  `json:"currencyCode"`
	ExDateStatus string  `json:"exDateStatus"`
}
type NextReport struct {
	Date       string `json:"date"`
	ReportType string `json:"reportType"`
}
type PreviousReport struct {
	Date       string `json:"date"`
	ReportType string `json:"reportType"`
}
type KeyIndicators struct {
	NumberOfOwners        int              `json:"numberOfOwners"`
	ReportDate            string           `json:"reportDate"`
	DirectYield           float64          `json:"directYield"`
	Volatility            float64          `json:"volatility"`
	Beta                  float64          `json:"beta"`
	PriceEarningsRatio    float64          `json:"priceEarningsRatio"`
	PriceSalesRatio       float64          `json:"priceSalesRatio"`
	ReturnOnEquity        float64          `json:"returnOnEquity"`
	ReturnOnTotalAssets   float64          `json:"returnOnTotalAssets"`
	EquityRatio           float64          `json:"equityRatio"`
	CapitalTurnover       float64          `json:"capitalTurnover"`
	OperatingProfitMargin float64          `json:"operatingProfitMargin"`
	GrossMargin           float64          `json:"grossMargin"`
	NetMargin             float64          `json:"netMargin"`
	MarketCapital         MarketCapital    `json:"marketCapital"`
	EquityPerShare        EquityPerShare   `json:"equityPerShare"`
	TurnoverPerShare      TurnoverPerShare `json:"turnoverPerShare"`
	EarningsPerShare      EarningsPerShare `json:"earningsPerShare"`
	Dividend              Dividend         `json:"dividend"`
	DividendsPerYear      int              `json:"dividendsPerYear"`
	NextReport            NextReport       `json:"nextReport"`
	PreviousReport        PreviousReport   `json:"previousReport"`
}
type AvanzaQuote struct {
	Last                       float64 `json:"last"`
	Highest                    float64 `json:"highest"`
	Lowest                     float64 `json:"lowest"`
	Change                     float64 `json:"change"`
	ChangePercent              float64 `json:"changePercent"`
	TimeOfLast                 int64   `json:"timeOfLast"`
	TotalValueTraded           float64 `json:"totalValueTraded"`
	TotalVolumeTraded          int     `json:"totalVolumeTraded"`
	Updated                    int64   `json:"updated"`
	VolumeWeightedAveragePrice float64 `json:"volumeWeightedAveragePrice"`
}
