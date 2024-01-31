package stocks

import (
	"fmt"
	"splitflap-backend/internal/models"
	"splitflap-backend/internal/utils"
)

type AvanzaClient struct {
}

func NewAvanzaClient() AvanzaClient {
	return AvanzaClient{}
}

func (c AvanzaClient) GetStockInfo(s Stock) (*models.StockInfo, error) {
	data, err := c.getStockData(s.AvanzaID)
	if err != nil {
		return nil, err
	}

	ytd := 100 * ((data.Quote.Last - data.HistoricalClosingPrices.StartOfYear) / data.HistoricalClosingPrices.StartOfYear)

	info := models.StockInfo{
		CompanyName: s.Company,
		DailyChange: data.Quote.ChangePercent,
		YTD:         ytd,
	}

	return &info, nil
}

func (c AvanzaClient) getStockData(avanzaID string) (*AvanzaResponse, error) {
	url := fmt.Sprintf("https://www.avanza.se/_api/market-guide/stock/%s", avanzaID)
	var series AvanzaResponse
	return utils.GetUrl(url, series)
}
