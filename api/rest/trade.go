package rest

import (
	"encoding/json"
	"net/http"

	"github.com/yitech/okex"
	requestsTrade "github.com/yitech/okex/requests/rest/trade"
	responsesTrade "github.com/yitech/okex/responses/trade"
)

// Trade
//
// https://www.okx.com/docs-v5/en/#rest-api-trade
type Trade struct {
	client *ClientRest
}

// NewTrade returns a pointer to a fresh Trade
func NewTrade(c *ClientRest) *Trade {
	return &Trade{c}
}

// PlaceOrder
// You can place an order only if you have sufficient funds.
//
// https://www.okx.com/docs-v5/en/#rest-api-trade-get-positions
func (c *Trade) PlaceOrder(req []requestsTrade.PlaceOrderRequest) (response responsesTrade.PlaceOrderResponse, err error) {
	p := "/api/v5/trade/order"
	var tmp interface{}
	tmp = req[0]
	if len(req) > 1 {
		tmp = req
		p = "/api/trade/batch-orders"
	}
	m := okex.S2M(tmp)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// PlaceMultipleOrders
// Cancel an incomplete order.
//
// https://www.okx.com/docs-v5/en/#rest-api-trade-place-multiple-orders
func (c *Trade) PlaceMultipleOrders(req []requestsTrade.PlaceOrderRequest) (response responsesTrade.PlaceOrderResponse, err error) {
	p := "/api/v5/trade/batch-order"
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

// CancelOrder
// Cancel an incomplete order.
//
// https://www.okx.com/docs-v5/en/#rest-api-trade-cancel-order
func (c *Trade) CancelOrder(req requestsTrade.CancelOrderRequest) (response responsesTrade.CancelOrderResponse, err error) {
	p := "/api/v5/trade/cancel-order"
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

// CancelBatchOrders
// Cancel incomplete orders in batches. Maximum 20 orders can be canceled at a time. Request parameters should be passed in the form of an array.
//
// https://www.okx.com/docs-v5/en/#rest-api-trade-cancel-multiple-orders
func (c *Trade) CancelBatchOrders(req []requestsTrade.CancelOrderRequest) (response responsesTrade.CancelOrderResponse, err error) {
	p := "/api/v5/trade/cancel-batch-orders"

	// Marshal the slice directly to JSON for batch endpoints
	j, err := json.Marshal(req)
	if err != nil {
		return
	}

	res, err := c.client.DoRawBody(http.MethodPost, p, true, j)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// AmendOrder
// Amend an incomplete order.
//
// https://www.okx.com/docs-v5/en/#rest-api-trade-amend-order
//
// Amend incomplete orders in batches. Maximum 20 orders can be amended at a time. Request parameters should be passed in the form of an array.
//
// https://www.okx.com/docs-v5/en/#rest-api-trade-amend-multiple-orders
func (c *Trade) AmendOrder(req []requestsTrade.OrderListRequest) (response responsesTrade.AmendOrderResponse, err error) {
	p := "/api/v5/trade/amend-order"
	var tmp interface{}
	tmp = req[0]
	if len(req) > 1 {
		tmp = req
		p = "/api/trade/amend-batch-orders"
	}
	m := okex.S2M(tmp)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// ClosePosition
// Close all positions of an instrument via a market order.
//
// https://www.okx.com/docs-v5/en/#rest-api-trade-close-positions
func (c *Trade) ClosePosition(req requestsTrade.ClosePositionRequest) (response responsesTrade.ClosePositionResponse, err error) {
	p := "/api/v5/trade/close-position"
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

// GetOrderDetail
// Retrieve order details.
//
// https://www.okx.com/docs-v5/en/#rest-api-trade-get-order-details
func (c *Trade) GetOrderDetail(req requestsTrade.OrderDetailsRequest) (response responsesTrade.OrderListResponse, err error) {
	p := "/api/v5/trade/order"
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

// GetOrderList
// Retrieve all incomplete orders under the current account.
//
// https://www.okx.com/docs-v5/en/#rest-api-trade-get-order-list
func (c *Trade) GetOrderList(req requestsTrade.OrderListRequest) (response responsesTrade.OrderListResponse, err error) {
	p := "/api/v5/trade/orders-pending"
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

// GetOrderHistory
// Retrieve the completed order data for the last 7 days, and the incomplete orders that have been cancelled are only reserved for 2 hours.
//
// https://www.okx.com/docs-v5/en/#rest-api-trade-get-order-history-last-7-days
//
// Retrieve the completed order data of the last 3 months, and the incomplete orders that have been canceled are only reserved for 2 hours.
// https://www.okx.com/docs-v5/en/#rest-api-trade-get-order-history-last-3-months
func (c *Trade) GetOrderHistory(req requestsTrade.OrderListRequest, arch bool) (response responsesTrade.OrderListResponse, err error) {
	p := "/api/v5/trade/orders-history"
	if arch {
		p = "/api/trade/orders-history-archive"
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

// GetTransactionDetails
// Retrieve recently-filled transaction details in the last 3 day.
//
// https://www.okx.com/docs-v5/en/#rest-api-trade-get-order-history-last-7-days
//
// Retrieve recently-filled transaction details in the last 3 months.
//
// https://www.okx.com/docs-v5/en/#rest-api-trade-get-transaction-details-last-3-months
func (c *Trade) GetTransactionDetails(req requestsTrade.TransactionDetailsRequest, arch bool) (response responsesTrade.TransactionDetailResponse, err error) {
	p := "/api/v5/trade/fills"
	if arch {
		p = "/api/trade/fills-history"
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

// PlaceAlgoOrder
// The algo order includes trigger order, oco order, conditional order,iceberg order and twap order.
//
// `iceberg` order and `twap` order just supported on demo trading
//
// https://www.okx.com/docs-v5/en/#rest-api-trade-place-algo-order
func (c *Trade) PlaceAlgoOrder(req requestsTrade.PlaceAlgoOrderRequest) (response responsesTrade.PlaceAlgoOrderResponse, err error) {
	p := "/api/v5/trade/order-algo"
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

// CancelAlgoOrder
// Cancel unfilled algo orders(trigger order, oco order, conditional order). A maximum of 10 orders can be canceled at a time. Request parameters should be passed in the form of an array.
//
// https://www.okx.com/docs-v5/en/#rest-api-trade-cancel-algo-order
func (c *Trade) CancelAlgoOrder(req requestsTrade.CancelAlgoOrderRequest) (response responsesTrade.CancelAlgoOrderResponse, err error) {
	p := "/api/v5/trade/cancel-algos"
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

// CancelAdvanceAlgoOrder
// Cancel unfilled algo orders(iceberg order and twap order). A maximum of 10 orders can be canceled at a time. Request parameters should be passed in the form of an array.
//
// # Only released on demo trading
//
// https://www.okx.com/docs-v5/en/#rest-api-trade-cancel-advance-algo-order
func (c *Trade) CancelAdvanceAlgoOrder(req requestsTrade.CancelAlgoOrderRequest) (response responsesTrade.CancelAlgoOrderResponse, err error) {
	p := "/api/v5/trade/cancel-advance-algos"
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

// GetAlgoOrderList
// Retrieve a list of untriggered Algo orders under the current account.
//
// https://www.okx.com/docs-v5/en/#rest-api-trade-get-algo-order-list
//
// Retrieve a list of untriggered Algo orders under the current account in the last 3 months.
//
// https://www.okx.com/docs-v5/en/#rest-api-trade-get-algo-order-history
func (c *Trade) GetAlgoOrderList(req requestsTrade.AlgoOrderListRequest, arch bool) (response responsesTrade.AlgoOrderListResponse, err error) {
	p := "/api/v5/trade/orders-algo-pending"
	if arch {
		p = "/api/trade/orders-algo-history"
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
