package account

import "github.com/yitech/okex"

type (
	GetBalanceRequest struct {
		Ccy []string `json:"ccy,omitempty"`
	}
	GetPositionsRequest struct {
		InstID   []string            `json:"instId,omitempty"`
		PosID    []string            `json:"posId,omitempty"`
		InstType okex.InstrumentType `json:"instType,omitempty"`
	}
	GetAccountAndPositionRiskRequest struct {
		InstType okex.InstrumentType `json:"instType,omitempty"`
	}
	GetBillsRequest struct {
		Ccy      string              `json:"ccy,omitempty"`
		After    int64               `json:"after,omitempty,string"`
		Before   int64               `json:"before,omitempty,string"`
		Limit    int64               `json:"limit,omitempty,string"`
		InstType okex.InstrumentType `json:"instType,omitempty"`
		MgnMode  okex.MarginMode     `json:"mgnMode,omitempty"`
		CtType   okex.ContractType   `json:"ctType,omitempty"`
		Type     okex.BillType       `json:"type,omitempty,string"`
		SubType  okex.BillSubType    `json:"subType,omitempty,string"`
	}
	SetPositionModeRequest struct {
		PositionMode okex.PositionType `json:"positionMode"`
	}
	SetLeverageRequest struct {
		Lever   int64             `json:"lever,string"`
		InstID  string            `json:"instId,omitempty"`
		Ccy     string            `json:"ccy,omitempty"`
		MgnMode okex.MarginMode   `json:"mgnMode"`
		PosSide okex.PositionSide `json:"posSide,omitempty"`
	}
	GetMaxBuySellAmountRequest struct {
		Ccy    string         `json:"ccy,omitempty"`
		Px     float64        `json:"px,string,omitempty"`
		InstID []string       `json:"instId"`
		TdMode okex.TradeMode `json:"tdMode"`
	}
	GetMaxAvailableTradeAmountRequest struct {
		Ccy        string         `json:"ccy,omitempty"`
		InstID     string         `json:"instId"`
		ReduceOnly bool           `json:"reduceOnly,omitempty"`
		TdMode     okex.TradeMode `json:"tdMode"`
	}
	IncreaseDecreaseMarginRequest struct {
		InstID     string            `json:"instId"`
		Amt        float64           `json:"amt,string"`
		PosSide    okex.PositionSide `json:"posSide"`
		ActionType okex.CountAction  `json:"actionType"`
	}
	GetLeverageRequest struct {
		InstID  []string        `json:"instId"`
		MgnMode okex.MarginMode `json:"mgnMode"`
	}
	GetMaxLoanRequest struct {
		InstID  string          `json:"instId"`
		MgnCcy  string          `json:"mgnCcy,omitempty"`
		MgnMode okex.MarginMode `json:"mgnMode"`
	}
	GetFeeRatesRequest struct {
		InstID   string              `json:"instId,omitempty"`
		Uly      string              `json:"uly,omitempty"`
		Category okex.FeeCategory    `json:"category,omitempty,string"`
		InstType okex.InstrumentType `json:"instType"`
	}
	GetInterestAccruedRequest struct {
		InstID  string          `json:"instId,omitempty"`
		Ccy     string          `json:"ccy,omitempty"`
		After   int64           `json:"after,omitempty,string"`
		Before  int64           `json:"before,omitempty,string"`
		Limit   int64           `json:"limit,omitempty,string"`
		MgnMode okex.MarginMode `json:"mgnMode,omitempty"`
	}
	SetGreeksRequest struct {
		GreeksType okex.GreekType `json:"greeksType"`
	}
)
