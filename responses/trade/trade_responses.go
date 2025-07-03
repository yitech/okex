package trade

import (
	"github.com/yitech/okex/models/trade"
	"github.com/yitech/okex/responses"
)

type (
	PlaceOrderResponse struct {
		responses.Basic
		Orders []*trade.PlaceOrder `json:"data,omitempty"`
	}

	CancelOrderResponse struct {
		responses.Basic
		Orders []*trade.CancelOrder `json:"data,omitempty"`
	}

	AmendOrderResponse struct {
		responses.Basic
		Orders []*trade.AmendOrder `json:"data,omitempty"`
	}

	ClosePositionResponse struct {
		responses.Basic
		Positions []*trade.ClosePosition `json:"data,omitempty"`
	}

	OrderListResponse struct {
		responses.Basic
		Orders []*trade.Order `json:"data,omitempty"`
	}

	TransactionDetailResponse struct {
		responses.Basic
		Transactions []*trade.TransactionDetail `json:"data,omitempty"`
	}

	PlaceAlgoOrderResponse struct {
		responses.Basic
		Orders []*trade.PlaceAlgoOrder `json:"data,omitempty"`
	}

	CancelAlgoOrderResponse struct {
		responses.Basic
		Orders []*trade.CancelAlgoOrder `json:"data,omitempty"`
	}

	AlgoOrderListResponse struct {
		responses.Basic
		Orders []*trade.AlgoOrder `json:"data,omitempty"`
	}
)
