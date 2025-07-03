package market

import (
	"github.com/yitech/okex/models/market"
	"github.com/yitech/okex/responses"
)

type (
	TickerResponse struct {
		responses.Basic
		Tickers []*market.Ticker `json:"data,omitempty"`
	}
	IndexTickerResponse struct {
		responses.Basic
		IndexTickers []*market.IndexTicker `json:"data,omitempty"`
	}
	OrderBookResponse struct {
		responses.Basic
		OrderBooks []*market.OrderBook `json:"data,omitempty"`
	}
	CandleResponse struct {
		responses.Basic
		Candles []*market.Candle `json:"data,omitempty"`
	}
	IndexCandleResponse struct {
		responses.Basic
		Candles []*market.IndexCandle `json:"data,omitempty"`
	}
	CandleMarketResponse struct {
		responses.Basic
		Candles []*market.IndexCandle `json:"data,omitempty"`
	}
	TradeResponse struct {
		responses.Basic
		Trades []*market.Trade `json:"data,omitempty"`
	}
	TotalVolume24HResponse struct {
		responses.Basic
		TotalVolume24Hs []*market.TotalVolume24H `json:"data,omitempty"`
	}
	IndexComponentResponse struct {
		responses.Basic
		IndexComponents *market.IndexComponent `json:"data,omitempty"`
	}
)
