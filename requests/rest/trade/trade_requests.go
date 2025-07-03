package trade

import "github.com/yitech/okex"

type (
	PlaceOrderRequest struct {
		InstID  string         `json:"instId"`
		TdMode  okex.TradeMode `json:"tdMode"`
		Side    okex.OrderSide `json:"side"`
		OrdType okex.OrderType `json:"ordType"`
		Sz      string         `json:"sz"`
		Px      string         `json:"px,omitempty"`
		Tag     string         `json:"tag,omitempty"`
		ClOrdId string         `json:"clOrdId,omitempty"`
	}

	CancelOrderRequest struct {
		InstID  string `json:"instId"`
		OrdId   string `json:"ordId,omitempty"`
		ClOrdId string `json:"clOrdId,omitempty"`
	}

	OrderListRequest struct {
		InstID  string          `json:"instId,omitempty"`
		OrdType okex.OrderType  `json:"ordType,omitempty"`
		State   okex.OrderState `json:"state,omitempty"`
		After   int64           `json:"after,omitempty,string"`
		Before  int64           `json:"before,omitempty,string"`
		Limit   int64           `json:"limit,omitempty,string"`
	}

	OrderDetailsRequest struct {
		InstID  string `json:"instId"`
		OrdId   string `json:"ordId,omitempty"`
		ClOrdId string `json:"clOrdId,omitempty"`
	}

	TransactionDetailsRequest struct {
		InstID string `json:"instId,omitempty"`
		OrdId  string `json:"ordId,omitempty"`
		After  int64  `json:"after,omitempty,string"`
		Before int64  `json:"before,omitempty,string"`
		Limit  int64  `json:"limit,omitempty,string"`
	}

	ClosePositionRequest struct {
		InstID  string            `json:"instId"`
		MgnMode okex.TradeMode    `json:"mgnMode"`
		PosSide okex.PositionSide `json:"posSide,omitempty"`
	}

	PlaceAlgoOrderRequest struct {
		InstID      string             `json:"instId"`
		TdMode      okex.TradeMode     `json:"tdMode"`
		Side        okex.OrderSide     `json:"side"`
		OrdType     okex.AlgoOrderType `json:"ordType"`
		Sz          string             `json:"sz"`
		Tag         string             `json:"tag,omitempty"`
		ClOrdId     string             `json:"clOrdId,omitempty"`
		SlTriggerPx string             `json:"slTriggerPx,omitempty"`
		SlOrdPx     string             `json:"slOrdPx,omitempty"`
		TpTriggerPx string             `json:"tpTriggerPx,omitempty"`
		TpOrdPx     string             `json:"tpOrdPx,omitempty"`
	}

	CancelAlgoOrderRequest struct {
		InstID  string `json:"instId"`
		AlgoId  string `json:"algoId,omitempty"`
		ClOrdId string `json:"clOrdId,omitempty"`
	}

	AlgoOrderListRequest struct {
		InstID  string             `json:"instId,omitempty"`
		OrdType okex.AlgoOrderType `json:"ordType,omitempty"`
		State   okex.OrderState    `json:"state,omitempty"`
		After   int64              `json:"after,omitempty,string"`
		Before  int64              `json:"before,omitempty,string"`
		Limit   int64              `json:"limit,omitempty,string"`
	}
)
