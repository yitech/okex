package market

import "github.com/yitech/okex"

type (
	GetTickersRequest struct {
		Uly      string              `json:"uly,omitempty"`
		InstType okex.InstrumentType `json:"instType"`
	}
	GetTickerRequest struct {
		InstID string `json:"instId"`
	}
	GetIndexTickersRequest struct {
		InstID   string `json:"instId,omitempty"`
		QuoteCcy string `json:"quoteCcy,omitempty"`
	}
	GetOrderBookRequest struct {
		InstID string `json:"instId"`
		Sz     int    `json:"sz,omitempty,string"`
	}
	GetCandlesticksRequest struct {
		InstID string       `json:"instId"`
		After  int64        `json:"after,omitempty,string"`
		Before int64        `json:"before,omitempty,string"`
		Limit  int64        `json:"limit,omitempty,string"`
		Bar    okex.BarSize `json:"bar,omitempty"`
	}
	GetTradesRequest struct {
		InstID string `json:"instId"`
		Limit  int64  `json:"limit,omitempty,string"`
	}
	GetIndexComponentsRequest struct {
		Index string `json:"index"`
	}
)
