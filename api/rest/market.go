package rest

import (
	"encoding/json"
	"net/http"

	"github.com/yitech/okex"
	requests "github.com/yitech/okex/requests/rest/market"
	responses "github.com/yitech/okex/responses/market"
)

// Market
//
// https://www.okx.com/docs-v5/en/#rest-api-market-data
type Market struct {
	client *ClientRest
}

// NewMarket returns a pointer to a fresh Market
func NewMarket(c *ClientRest) *Market {
	return &Market{c}
}

// GetTickers
// Retrieve the latest price snapshot, best bid/ask price, and trading volume in the last 24 hours.
//
// https://www.okx.com/docs-v5/en/#rest-api-market-data-get-tickers
func (c *Market) GetTickers(req requests.GetTickersRequest) (response responses.TickerResponse, err error) {
	p := "/api/v5/market/tickers"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// GetTicker
// Retrieve the latest price snapshot, best bid/ask price, and trading volume in the last 24 hours.
//
// https://www.okx.com/docs-v5/en/#rest-api-market-data-get-ticker
func (c *Market) GetTicker(req requests.GetTickerRequest) (response responses.TickerResponse, err error) {
	p := "/api/v5/market/ticker"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// GetIndexTickers
// Retrieve index tickers.
//
// https://www.okx.com/docs-v5/en/#rest-api-market-data-get-index-tickers
func (c *Market) GetIndexTickers(req requests.GetIndexTickersRequest) (response responses.IndexTickerResponse, err error) {
	p := "/api/v5/market/ticker"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// GetOrderBook
// Retrieve a instrument is order book.
//
// https://www.okx.com/docs-v5/en/#rest-api-market-data-get-order-book
func (c *Market) GetOrderBook(req requests.GetOrderBookRequest) (response responses.OrderBookResponse, err error) {
	p := "/api/v5/market/books"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// GetCandlesticks
// Retrieve the candlestick charts. This endpoint can retrieve the latest 1,440 data entries. Charts are returned in groups based on the requested bar.
//
// https://www.okx.com/docs-v5/en/#rest-api-market-data-get-candlesticks
func (c *Market) GetCandlesticks(req requests.GetCandlesticksRequest) (response responses.CandleResponse, err error) {
	p := "/api/v5/market/candles"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// GetCandlesticksHistory
// Retrieve history candlestick charts from recent years.
//
// https://www.okx.com/docs-v5/en/#rest-api-market-data-get-candlesticks
func (c *Market) GetCandlesticksHistory(req requests.GetCandlesticksRequest) (response responses.CandleResponse, err error) {
	p := "/api/v5/market/history-candles"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// GetIndexCandlesticks
// Retrieve the candlestick charts of the index. This endpoint can retrieve the latest 1,440 data entries. Charts are returned in groups based on the requested bar.
//
// https://www.okx.com/docs-v5/en/#rest-api-market-data-get-index-candlesticks
func (c *Market) GetIndexCandlesticks(req requests.GetCandlesticksRequest) (response responses.IndexCandleResponse, err error) {
	p := "/api/v5/market/index-candles"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// GetMarkPriceCandlesticks
// Retrieve the candlestick charts of mark price. This endpoint can retrieve the latest 1,440 data entries. Charts are returned in groups based on the requested bar.
//
// https://www.okx.com/docs-v5/en/#rest-api-market-data-get-mark-price-candlesticks
func (c *Market) GetMarkPriceCandlesticks(req requests.GetCandlesticksRequest) (response responses.CandleMarketResponse, err error) {
	p := "/api/v5/market/mark-price-candles"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// GetTrades
// Retrieve the recent transactions of an instrument.
//
// https://www.okx.com/docs-v5/en/#rest-api-market-data-get-trades
func (c *Market) GetTrades(req requests.GetTradesRequest) (response responses.TradeResponse, err error) {
	p := "/api/v5/market/trades"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// Get24HTotalVolume
// The 24-hour trading volume is calculated on a rolling basis, using USD as the pricing unit.
//
// https://www.okx.com/docs-v5/en/#rest-api-market-data-get-24h-total-volume
func (c *Market) Get24HTotalVolume() (response responses.TotalVolume24HResponse, err error) {
	p := "/api/v5/market/platform-24-volume"
	res, err := c.client.Do(http.MethodGet, p, false)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// GetIndexComponents
// Get the index component information data on the market
//
// https://www.okx.com/docs-v5/en/#rest-api-market-data-get-index-components
func (c *Market) GetIndexComponents(req requests.GetIndexComponentsRequest) (response responses.IndexComponentResponse, err error) {
	p := "/api/v5/market/index-components"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}
