package account

import (
	models "github.com/yitech/okex/models/account"
	"github.com/yitech/okex/responses"
)

type (
	GetBalanceResponse struct {
		responses.Basic
		Balances []*models.Balance `json:"data,omitempty"`
	}
	GetPositionsResponse struct {
		responses.Basic
		Positions []*models.Position `json:"data"`
	}
	GetAccountAndPositionRiskResponse struct {
		responses.Basic
		PositionAndAccountRisks []*models.PositionAndAccountRisk `json:"data"`
	}
	GetBillsResponse struct {
		responses.Basic
		Bills []*models.Bill `json:"data"`
	}
	GetConfigResponse struct {
		responses.Basic
		Configs []*models.Config `json:"data"`
	}
	SetPositionModeResponse struct {
		responses.Basic
		PositionModes []*models.PositionMode `json:"data"`
	}
	LeverageResponse struct {
		responses.Basic
		Leverages []*models.Leverage `json:"data"`
	}
	GetMaxBuySellAmountResponse struct {
		responses.Basic
		MaxBuySellAmounts []*models.MaxBuySellAmount `json:"data"`
	}
	GetMaxAvailableTradeAmountResponse struct {
		responses.Basic
		MaxAvailableTradeAmounts []*models.MaxAvailableTradeAmount `json:"data"`
	}
	IncreaseDecreaseMarginResponse struct {
		responses.Basic
		MarginBalanceAmounts []*models.MarginBalanceAmount `json:"data"`
	}
	GetMaxLoanResponse struct {
		responses.Basic
		Loans []*models.Loan `json:"data"`
	}
	GetFeeRatesResponse struct {
		responses.Basic
		Fees []*models.Fee `json:"data"`
	}
	GetInterestAccruedResponse struct {
		responses.Basic
		InterestAccrues []*models.InterestAccrued `json:"data"`
	}
	GetInterestRatesResponse struct {
		responses.Basic
		Interests []*models.InterestRate `json:"data"`
	}
	SetGreeksResponse struct {
		responses.Basic
		Greeks []*models.Greek `json:"data"`
	}
	GetMaxWithdrawalsResponse struct {
		responses.Basic
		MaxWithdrawals []*models.MaxWithdrawal `json:"data"`
	}
)
