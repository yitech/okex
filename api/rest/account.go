package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/yitech/okex"
	requests "github.com/yitech/okex/requests/rest/account"
	responses "github.com/yitech/okex/responses/account"
)

// Account
//
// https://www.okx.com/docs-v5/en/#rest-api-account
type Account struct {
	client *ClientRest
}

// NewAccount returns a pointer to a fresh Account
func NewAccount(c *ClientRest) *Account {
	return &Account{c}
}

// GetBalance
// Retrieve a list of assets (with non-zero balance), remaining balance, and available amount in the account.
//
// https://www.okx.com/docs-v5/en/#rest-api-account-get-balance
func (c *Account) GetBalance(req requests.GetBalanceRequest) (response responses.GetBalanceResponse, err error) {
	p := "/api/v5/account/balance"
	m := okex.S2M(req)
	if len(req.Ccy) > 0 {
		m["ccy"] = strings.Join(req.Ccy, ",")
	}
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetPositions
// Retrieve information on your positions. When the account is in net mode, net positions will be displayed, and when the account is in long/short mode, long or short positions will be displayed.
//
// https://www.okx.com/docs-v5/en/#rest-api-account-get-positions
func (c *Account) GetPositions(req requests.GetPositionsRequest) (response responses.GetPositionsResponse, err error) {
	p := "/api/v5/account/positions"
	m := okex.S2M(req)
	if len(req.InstID) > 0 {
		m["instId"] = strings.Join(req.InstID, ",")
	}
	if len(req.PosID) > 0 {
		m["posId"] = strings.Join(req.PosID, ",")
	}
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetAccountAndPositionRisk
// Get account and position risk
//
// https://www.okx.com/docs-v5/en/#rest-api-account-get-account-and-position-risk
func (c *Account) GetAccountAndPositionRisk(req requests.GetAccountAndPositionRiskRequest) (response responses.GetAccountAndPositionRiskResponse, err error) {
	p := "/api/v5/account/positions"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetBills
// Retrieve the bills of the account. The bill refers to all transaction records that result in changing the balance of an account. Pagination is supported, and the response is sorted with the most recent first. This endpoint can retrieve data from the last 7 days.
//
// https://www.okx.com/docs-v5/en/#rest-api-account-get-bills-details-last-7-days
//
// Retrieve the account’s bills. The bill refers to all transaction records that result in changing the balance of an account. Pagination is supported, and the response is sorted with most recent first. This endpoint can retrieve data from the last 3 months.
//
// https://www.okx.com/docs-v5/en/#rest-api-account-get-bills-details-last-3-months
func (c *Account) GetBills(req requests.GetBillsRequest, arc bool) (response responses.GetBillsResponse, err error) {
	p := "/api/v5/account/bills"
	if arc {
		p = "/api/account/bills-archive"
	}
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetConfig
// Retrieve current account configuration.
//
// https://www.okx.com/docs-v5/en/#rest-api-account-get-account-configuration
func (c *Account) GetConfig() (response responses.GetConfigResponse, err error) {
	p := "/api/v5/account/config"
	res, err := c.client.Do(http.MethodGet, p, true)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// SetPositionMode
// FUTURES and SWAP support both long/short mode and net mode. In net mode, users can only have positions in one direction; In long/short mode, users can hold positions in long and short directions.
//
// https://www.okx.com/docs-v5/en/#rest-api-account-set-position-mode
func (c *Account) SetPositionMode(req requests.SetPositionModeRequest) (response responses.SetPositionModeResponse, err error) {
	p := "/api/v5/account/set-position-mode"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// SetLeverage
// The following are the setting leverage cases for an instrument:
//
// Set leverage for isolated MARGIN at pairs level.
//
// Set leverage for cross MARGIN in Single-currency margin at pairs level.
//
// Set leverage for cross MARGIN in Multi-currency margin at currency level.
//
// Set leverage for cross/isolated FUTURES/SWAP at underlying/contract level.
// https://www.okx.com/docs-v5/en/#rest-api-account-set-leverage
func (c *Account) SetLeverage(req requests.SetLeverageRequest) (response responses.LeverageResponse, err error) {
	p := "/api/v5/account/set-leverage"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetMaxBuySellAmount
//
// https://www.okx.com/docs-v5/en/#rest-api-account-get-maximum-buy-sell-amount-or-open-amount
func (c *Account) GetMaxBuySellAmount(req requests.GetMaxBuySellAmountRequest) (response responses.GetMaxBuySellAmountResponse, err error) {
	p := "/api/v5/account/max-size"
	m := okex.S2M(req)
	if len(req.InstID) > 0 {
		m["instId"] = strings.Join(req.InstID, ",")
	}
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetMaxAvailableTradeAmount
//
// https://www.okx.com/docs-v5/en/#rest-api-account-get-maximum-available-tradable-amount
func (c *Account) GetMaxAvailableTradeAmount(req requests.GetMaxAvailableTradeAmountRequest) (response responses.GetMaxAvailableTradeAmountResponse, err error) {
	p := "/api/v5/account/max-avail-size"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// IncreaseDecreaseMargin
// Increase or decrease the margin of the isolated position.
//
// https://www.okx.com/docs-v5/en/#rest-api-account-increase-decrease-margin
func (c *Account) IncreaseDecreaseMargin(req requests.IncreaseDecreaseMarginRequest) (response responses.IncreaseDecreaseMarginResponse, err error) {
	p := "/api/v5/account/position/margin-balance"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetLeverage
//
// https://www.okx.com/docs-v5/en/#rest-api-account-get-leverage
func (c *Account) GetLeverage(req requests.GetLeverageRequest) (response responses.LeverageResponse, err error) {
	p := "/api/v5/account/leverage-info"
	m := okex.S2M(req)
	if len(req.InstID) > 0 {
		m["instId"] = strings.Join(req.InstID, ",")
	}
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetMaxLoan
//
// https://www.okx.com/docs-v5/en/#rest-api-account-get-the-maximum-loan-of-instrument
func (c *Account) GetMaxLoan(req requests.GetMaxLoanRequest) (response responses.GetMaxLoanResponse, err error) {
	p := "/api/v5/account/max-loan"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetFeeRates
//
// https://www.okx.com/docs-v5/en/#rest-api-account-get-fee-rates
func (c *Account) GetFeeRates(req requests.GetFeeRatesRequest) (response responses.GetFeeRatesResponse, err error) {
	p := "/api/v5/account/trade-fee"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetInterestAccrued
//
// https://www.okx.com/docs-v5/en/#rest-api-account-get-interest-accrued
func (c *Account) GetInterestAccrued(req requests.GetInterestAccruedRequest) (response responses.GetInterestAccruedResponse, err error) {
	p := "/api/v5/account/interest-accrued"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetInterestRates
// Get the user's current leveraged currency borrowing interest rate
//
// https://www.okx.com/docs-v5/en/#rest-api-account-get-interest-rate
func (c *Account) GetInterestRates(req requests.GetInterestAccruedRequest) (response responses.GetInterestAccruedResponse, err error) {
	p := "/api/v5/account/interest-rate"
	m := okex.S2M(req)
	if req.Ccy != "" {
		m["ccy"] = req.Ccy
	}
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// SetGreeks
// Set the display type of Greeks.
//
// https://www.okx.com/docs-v5/en/#rest-api-account-set-greeks-m-bs
func (c *Account) SetGreeks(req requests.SetGreeksRequest) (response responses.SetGreeksResponse, err error) {
	p := "/api/v5/account/set-greeks"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetMaxWithdrawals
//
// https://www.okx.com/docs-v5/en/#rest-api-account-get-maximum-withdrawals
func (c *Account) GetMaxWithdrawals(req requests.GetBalanceRequest) (response responses.GetMaxWithdrawalsResponse, err error) {
	p := "/api/v5/account/max-withdrawal"
	m := okex.S2M(req)
	if len(req.Ccy) > 0 {
		m["ccy"] = strings.Join(req.Ccy, ",")
	}
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}
