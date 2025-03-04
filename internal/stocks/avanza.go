package stocks

import (
	"fmt"
	"splitflap-backend/internal/models"
	"splitflap-backend/pkg/fluent"
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

const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"

func (c AvanzaClient) getStockData(avanzaID string) (*AvanzaResponse, error) {
	url := fmt.Sprintf("https://www.avanza.se/_api/market-guide/stock/%s", avanzaID)
	var series *AvanzaResponse
	err := fluent.Get(url).
		WithHeader("user-agent", userAgent).
		OnSuccess(func(bytes []byte) error {
			res, innerErr := fluent.Deserialize[AvanzaResponse](bytes)
			series = res
			return innerErr
		}).
		Execute()

	return series, err
}
