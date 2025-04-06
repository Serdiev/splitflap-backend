package stocks

import (
	"fmt"
	"splitflap-backend/internal/models"
	"splitflap-backend/pkg/fluent"
	"time"
)

const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"

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

func (c AvanzaClient) GetOHLCData(s Stock) (*models.OHLCData, error) {
	data, err := c.getOHLCData(s.AvanzaID)
	if err != nil {
		return nil, err
	}

	if len(data.Ohlc) < 2 {
		return nil, fmt.Errorf("no data found for %s", s.Company)
	}

	from, err := time.Parse("2006-01-02", data.From)
	if err != nil {
		return nil, err
	}
	to, err := time.Parse("2006-01-02", data.To)
	if err != nil {
		return nil, err
	}

	d := models.OHLCData{
		CompanyName:          s.Company,
		From:                 from,
		To:                   to,
		PreviousClosingPrice: data.PreviousClosingPrice,
		OHLC:                 []models.OHLC{},
	}

	for _, ohlc := range data.Ohlc {
		d.OHLC = append(d.OHLC, models.OHLC{
			Timestamp:         time.UnixMilli(ohlc.Timestamp),
			Open:              ohlc.Open,
			Close:             ohlc.Close,
			High:              ohlc.High,
			Low:               ohlc.Low,
			TotalVolumeTraded: ohlc.TotalVolumeTraded,
		})
	}

	return &d, nil
}

func (c AvanzaClient) getOHLCData(avanzaID string) (*OHLCResponse, error) {
	url := fmt.Sprintf("https://www.avanza.se/_api/price-chart/stock/%s?timePeriod=this_year", avanzaID)
	var series *OHLCResponse
	err := fluent.Get(url).
		WithHeader("user-agent", userAgent).
		OnSuccess(func(bytes []byte) error {
			res, innerErr := fluent.Deserialize[OHLCResponse](bytes)
			series = res
			return innerErr
		}).
		Execute()

	return series, err
}
