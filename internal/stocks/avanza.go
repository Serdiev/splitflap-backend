package stocks

import (
	"fmt"
	"splitflap-backend/internal/models"
	"splitflap-backend/internal/utils"
)

type AvanzaClient struct {
	baseUrl string
}

func NewAvanzaClient() AvanzaClient {
	return AvanzaClient{
		baseUrl: cfg.ALPHA_VANTAGE_URL,
	}
}

func (c AvanzaClient) GetStockInfo(s Stock) (*models.StockInfo, error) {
	data, err := c.getStockData(s.AvanzaID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	ytd := (data.Quote.Last / data.HistoricalClosingPrices.StartOfYear) - 1

	info := models.StockInfo{
		CompanyName: s.Company,
		DailyChange: data.Quote.ChangePercent,
		YTD:         ytd,
	}

	return &info, nil
}

// TimeSeries gets daily, weekly and monthly series
func (c AvanzaClient) getStockData(avanzaID string) (*AvanzaResponse, error) {
	url := fmt.Sprintf("https://www.avanza.se/_api/market-guide/stock/%s", avanzaID)
	var series AvanzaResponse
	return utils.GetUrlParseIfOK(url, series)
}
