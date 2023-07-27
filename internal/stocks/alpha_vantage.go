package stocks

import (
	"fmt"
	config "splitflap-backend/configs"
	"splitflap-backend/internal/models"
	"splitflap-backend/internal/utils"
)

var cfg = config.New()

type AlphaVantageClient struct {
	baseUrl string
	apiKey  string
}

func NewNoopStockClient() AlphaVantageClient {
	return AlphaVantageClient{}
}

func NewStockClient() AlphaVantageClient {
	return AlphaVantageClient{
		baseUrl: cfg.ALPHA_VANTAGE_URL,
		apiKey:  cfg.ALPHA_VANTAGE_API_KEY,
	}
}

func (c AlphaVantageClient) GetStockInfo(symbol string) (*models.StockInfo, error) {
	series, err := c.getTimeSeries(MonthlyAdjusted, symbol)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	q, err := c.getQuote(symbol)

	info := models.StockInfo{
		Tick:        symbol,
		Company:     series.MetaData.Symbol,
		DailyChange: q.ChangePercent,
		YTD:         q.Change,
	}

	return &info, nil
}

// TimeSeries gets daily, weekly and monthly series
func (c AlphaVantageClient) getTimeSeries(function, symbol string) (*TimeSeriesData, error) {
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s", c.baseUrl, function, symbol, c.apiKey)
	var series TimeSeriesData
	return utils.GetUrlParseIfOK(url, series)
}

// get latest info of symbol
func (c AlphaVantageClient) getQuote(symbol string) (*Quote, error) {
	url := fmt.Sprintf("%s/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", c.baseUrl, symbol, c.apiKey)
	var quote QuoteResponse
	res, err := utils.GetUrlParseIfOK(url, quote)

	if err != nil {
		return nil, err
	}

	return &res.Quote, nil
}
