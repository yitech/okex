package ws

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/yitech/okex"
	"github.com/yitech/okex/events"
	"github.com/yitech/okex/events/public"
	requests "github.com/yitech/okex/requests/ws/public"
)

// Public
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels
type Public struct {
	*ClientWs
	iCh    chan *public.Instruments
	tCh    chan *public.Tickers
	oiCh   chan *public.OpenInterest
	cCh    chan *public.Candlesticks
	trCh   chan *public.Trades
	edepCh chan *public.EstimatedDeliveryExercisePrice
	mpCh   chan *public.MarkPrice
	mpcCh  chan *public.MarkPriceCandlesticks
	plCh   chan *public.PriceLimit
	obCh   chan *public.OrderBook
	osCh   chan *public.OPTIONSummary
	frCh   chan *public.FundingRate
	icCh   chan *public.IndexCandlesticks
	itCh   chan *public.IndexTickers
}

// NewPublic returns a pointer to a fresh Public
func NewPublic(c *ClientWs) *Public {
	return &Public{ClientWs: c}
}

// Instruments
// The full instrument list will be pushed for the first time after subscription. Subsequently, the instruments will be pushed if there's any change to the instrument’s state (such as delivery of FUTURES, exercise of OPTION, listing of new contracts / trading pairs, trading suspension, etc.).
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-instruments-channel
func (c *Public) Instruments(req requests.Instruments, ch ...chan *public.Instruments) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.iCh = ch[0]
	}
	return c.Subscribe(false, []okex.ChannelName{"instruments"}, m)
}

// UInstruments
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-instruments-channel
func (c *Public) UInstruments(req requests.Instruments, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.iCh = nil
	}
	return c.Unsubscribe(false, []okex.ChannelName{"instruments"}, m)
}

// Tickers
// Retrieve the last traded price, bid price, ask price and 24-hour trading volume of instruments. Data will be pushed every 100 ms.
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-tickers-channel
func (c *Public) Tickers(req requests.Tickers, ch ...chan *public.Tickers) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.tCh = ch[0]
	}
	return c.Subscribe(false, []okex.ChannelName{"tickers"}, m)
}

// UTickers
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-tickers-channel
func (c *Public) UTickers(req requests.Tickers, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.tCh = nil
	}
	return c.Unsubscribe(false, []okex.ChannelName{"tickers"}, m)
}

// OpenInterest
// Retrieve the open interest. Data will be pushed every 3 seconds.
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-open-interest-channel
func (c *Public) OpenInterest(req requests.OpenInterest, ch ...chan *public.OpenInterest) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.oiCh = ch[0]
	}
	return c.Subscribe(false, []okex.ChannelName{"open-interest"}, m)
}

// UOpenInterest
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-open-interest-channel
func (c *Public) UOpenInterest(req requests.OpenInterest, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.oiCh = nil
	}
	return c.Unsubscribe(false, []okex.ChannelName{"open-interest"}, m)
}

// Candlesticks
// Retrieve the open interest. Data will be pushed every 3 seconds.
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-candlesticks-channel
func (c *Public) Candlesticks(req requests.Candlesticks, ch ...chan *public.Candlesticks) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.cCh = ch[0]
	}
	return c.Subscribe(false, []okex.ChannelName{}, m)
}

// UCandlesticks
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-candlesticks-channel
func (c *Public) UCandlesticks(req requests.Candlesticks, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.cCh = nil
	}
	return c.Unsubscribe(false, []okex.ChannelName{}, m)
}

// Trades
// Retrieve the recent trades data. Data will be pushed whenever there is a trade.
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-trades-channel
func (c *Public) Trades(req requests.Trades, ch ...chan *public.Trades) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.trCh = ch[0]
	}
	return c.Subscribe(false, []okex.ChannelName{"trades"}, m)
}

// UTrades
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-trades-channel
func (c *Public) UTrades(req requests.Trades, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.trCh = nil
	}
	return c.Unsubscribe(false, []okex.ChannelName{"trades"}, m)
}

// EstimatedDeliveryExercisePrice
// Retrieve the estimated delivery/exercise price of FUTURES contracts and OPTION.
//
// Only the estimated delivery/exercise price will be pushed an hour before delivery/exercise, and will be pushed if there is any price change.
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-estimated-delivery-exercise-price-channel
func (c *Public) EstimatedDeliveryExercisePrice(req requests.EstimatedDeliveryExercisePrice, ch ...chan *public.EstimatedDeliveryExercisePrice) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.edepCh = ch[0]
	}
	return c.Subscribe(false, []okex.ChannelName{"estimated-price"}, m)
}

// UEstimatedDeliveryExercisePrice
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-estimated-delivery-exercise-price-channel
func (c *Public) UEstimatedDeliveryExercisePrice(req requests.EstimatedDeliveryExercisePrice, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.edepCh = nil
	}
	return c.Unsubscribe(false, []okex.ChannelName{"estimated-price"}, m)
}

// MarkPrice
// Retrieve the mark price. Data will be pushed every 200 ms when the mark price changes, and will be pushed every 10 seconds when the mark price does not change.
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-mark-price-channel
func (c *Public) MarkPrice(req requests.MarkPrice, ch ...chan *public.MarkPrice) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.mpCh = ch[0]
	}
	return c.Subscribe(false, []okex.ChannelName{"mark-price"}, m)
}

// UMarkPrice
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-mark-price-channel
func (c *Public) UMarkPrice(req requests.MarkPrice, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.mpCh = nil
	}
	return c.Unsubscribe(false, []okex.ChannelName{"mark-price"}, m)
}

// MarkPriceCandlesticks
// Retrieve the candlesticks data of the mark price. Data will be pushed every 500 ms.
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-mark-price-candlesticks-channel
func (c *Public) MarkPriceCandlesticks(req requests.MarkPriceCandlesticks, ch ...chan *public.MarkPriceCandlesticks) error {
	m := okex.S2M(req)
	m["channel"] = "mark-price-" + m["channel"]
	if len(ch) > 0 {
		c.mpcCh = ch[0]
	}
	return c.Subscribe(false, []okex.ChannelName{}, m)
}

// UMarkPriceCandlesticks
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-mark-price-candlesticks-channel
func (c *Public) UMarkPriceCandlesticks(req requests.MarkPriceCandlesticks, rCh ...bool) error {
	m := okex.S2M(req)
	m["channel"] = "mark-price-" + m["channel"]
	if len(rCh) > 0 && rCh[0] {
		c.mpcCh = nil
	}
	return c.Unsubscribe(false, []okex.ChannelName{}, m)
}

// PriceLimit
// Retrieve the maximum buy price and minimum sell price of the instrument. Data will be pushed every 5 seconds when there are changes in limits, and will not be pushed when there is no changes on limit.
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-price-limit-channel
func (c *Public) PriceLimit(req requests.PriceLimit, ch ...chan *public.PriceLimit) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.plCh = ch[0]
	}
	return c.Subscribe(false, []okex.ChannelName{"price-limit"}, m)
}

// UPriceLimit
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-price-limit-channel
func (c *Public) UPriceLimit(req requests.PriceLimit, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.plCh = nil
	}
	return c.Unsubscribe(false, []okex.ChannelName{"price-limit"}, m)
}

// OrderBook
// Retrieve order book data for multiple instruments.
//
// Use books for 400 depth levels, book5 for 5 depth levels, books50-l2-tbt tick-by-tick 50 depth levels, and books-l2-tbt for tick-by-tick 400 depth levels.
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-order-book-channel
func (c *Public) OrderBook(reqs []requests.OrderBook, ch ...chan *public.OrderBook) error {
	if len(ch) > 0 {
		c.obCh = ch[0]
	}
	var subscriptions []map[string]string
	for _, req := range reqs {
		m := okex.S2M(req)
		subscriptions = append(subscriptions, m)
	}
	return c.Subscribe(false, []okex.ChannelName{}, subscriptions...)
}

// UOrderBook
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-order-book-channel
func (c *Public) UOrderBook(req requests.OrderBook, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.obCh = nil
	}
	return c.Unsubscribe(false, []okex.ChannelName{okex.ChannelName(req.Channel)}, m)
}

// OPTIONSummary
// Retrieve detailed pricing information of all OPTION contracts. Data will be pushed at once.
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-option-summary-channel
func (c *Public) OPTIONSummary(req requests.OPTIONSummary, ch ...chan *public.OPTIONSummary) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.osCh = ch[0]
	}
	return c.Subscribe(false, []okex.ChannelName{"opt-summary"}, m)
}

// UOPTIONSummary
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-option-summary-channel
func (c *Public) UOPTIONSummary(req requests.OPTIONSummary, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.osCh = nil
	}
	return c.Unsubscribe(false, []okex.ChannelName{"opt-summary"}, m)
}

// FundingRate
// Retrieve funding rate. Data will be pushed every minute.
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-funding-rate-channel
func (c *Public) FundingRate(req requests.FundingRate, ch ...chan *public.FundingRate) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.frCh = ch[0]
	}
	return c.Subscribe(false, []okex.ChannelName{"funding-rate"}, m)
}

// UFundingRate
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-funding-rate-channel
func (c *Public) UFundingRate(req requests.FundingRate, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.frCh = nil
	}
	return c.Unsubscribe(false, []okex.ChannelName{"funding-rate"}, m)
}

// IndexCandlesticks
// Retrieve the candlesticks data of the index. Data will be pushed every 500 ms.
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-index-candlesticks-channel
func (c *Public) IndexCandlesticks(req requests.IndexCandlesticks, ch ...chan *public.IndexCandlesticks) error {
	m := okex.S2M(req)
	m["channel"] = req.Channel
	if len(ch) > 0 {
		c.icCh = ch[0]
	}
	return c.Subscribe(false, []okex.ChannelName{}, m)
}

// UIndexCandlesticks
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-index-candlesticks-channel
func (c *Public) UIndexCandlesticks(req requests.IndexCandlesticks, rCh ...bool) error {
	m := okex.S2M(req)
	m["channel"] = req.Channel
	if len(rCh) > 0 && rCh[0] {
		c.icCh = nil
	}
	return c.Unsubscribe(false, []okex.ChannelName{}, m)
}

// IndexTickers
// Retrieve index tickers data
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-index-tickers-channel
func (c *Public) IndexTickers(req requests.IndexTickers, ch ...chan *public.IndexTickers) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.itCh = ch[0]
	}
	return c.Subscribe(false, []okex.ChannelName{"index-tickers"}, m)
}

// UIndexTickers
//
// https://www.okx.com/docs-v5/en/#websocket-api-public-channels-index-tickers-channel
func (c *Public) UIndexTickers(req requests.IndexTickers, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.itCh = nil
	}
	return c.Unsubscribe(false, []okex.ChannelName{"index-tickers"}, m)
}

func (c *Public) Process(data []byte, e *events.Basic) bool {
	if e.Event == "" && e.Arg != nil && e.Data != nil && len(e.Data) > 0 {
		ch, ok := e.Arg.Get("channel")
		if !ok {
			return false
		}
		switch ch {
		case "instruments":
			e := public.Instruments{}
			if err := json.Unmarshal(data, &e); err != nil {
				return false
			}
			if c.iCh != nil {
				c.iCh <- &e
			}
			if c.StructuredEventChan != nil {
				c.StructuredEventChan <- e
			}
			return true
		case "tickers":
			e := public.Tickers{}
			if err := json.Unmarshal(data, &e); err != nil {
				return false
			}
			if c.tCh != nil {
				c.tCh <- &e
			}
			if c.StructuredEventChan != nil {
				c.StructuredEventChan <- e
			}
			return true
		case "open-interest":
			e := public.OpenInterest{}
			if err := json.Unmarshal(data, &e); err != nil {
				return false
			}
			if c.oiCh != nil {
				c.oiCh <- &e
			}
			if c.StructuredEventChan != nil {
				c.StructuredEventChan <- e
			}
			return true
		case "trades":
			e := public.Trades{}
			if err := json.Unmarshal(data, &e); err != nil {
				return false
			}
			if c.trCh != nil {
				c.trCh <- &e
			}
			if c.StructuredEventChan != nil {
				c.StructuredEventChan <- e
			}
			return true
		case "estimated-price":
			e := public.EstimatedDeliveryExercisePrice{}
			if err := json.Unmarshal(data, &e); err != nil {
				return false
			}
			if c.edepCh != nil {
				c.edepCh <- &e
			}
			if c.StructuredEventChan != nil {
				c.StructuredEventChan <- e
			}
			return true
		case "mark-price":
			e := public.MarkPrice{}
			if err := json.Unmarshal(data, &e); err != nil {
				return false
			}
			if c.mpCh != nil {
				c.mpCh <- &e
			}
			if c.StructuredEventChan != nil {
				c.StructuredEventChan <- e
			}
			return true
		case "price-limit":
			e := public.PriceLimit{}
			if err := json.Unmarshal(data, &e); err != nil {
				return false
			}
			if c.plCh != nil {
				c.plCh <- &e
			}
			if c.StructuredEventChan != nil {
				c.StructuredEventChan <- e
			}
			return true
		case "opt-summary":
			e := public.OPTIONSummary{}
			if err := json.Unmarshal(data, &e); err != nil {
				return false
			}
			if c.osCh != nil {
				c.osCh <- &e
			}
			if c.StructuredEventChan != nil {
				c.StructuredEventChan <- e
			}
			return true
		case "funding-rate":
			e := public.FundingRate{}
			if err := json.Unmarshal(data, &e); err != nil {
				return false
			}
			if c.frCh != nil {
				c.frCh <- &e
			}
			if c.StructuredEventChan != nil {
				c.StructuredEventChan <- e
			}
			return true
		case "index-tickers":
			e := public.IndexTickers{}
			if err := json.Unmarshal(data, &e); err != nil {
				return false
			}
			if c.itCh != nil {
				c.itCh <- &e
			}
			if c.StructuredEventChan != nil {
				c.StructuredEventChan <- e
			}
			return true
		default:
			chName := fmt.Sprint(ch)
			if strings.Contains(chName, "mark-price-candle") {
				e := public.MarkPriceCandlesticks{}
				if err := json.Unmarshal(data, &e); err != nil {
					return false
				}
				if c.mpcCh != nil {
					c.mpcCh <- &e
				}
				if c.StructuredEventChan != nil {
					c.StructuredEventChan <- e
				}
				return true
			}
			if strings.Contains(chName, "index-candle") {
				e := public.IndexCandlesticks{}
				if err := json.Unmarshal(data, &e); err != nil {
					return false
				}
				if c.icCh != nil {
					c.icCh <- &e
				}
				if c.StructuredEventChan != nil {
					c.StructuredEventChan <- e
				}
				return true
			}
			if strings.Contains(chName, "candle") {
				e := public.Candlesticks{}
				if err := json.Unmarshal(data, &e); err != nil {
					return false
				}
				if c.cCh != nil {
					c.cCh <- &e
				}
				if c.StructuredEventChan != nil {
					c.StructuredEventChan <- e
				}
				return true
			}
			if strings.Contains(chName, "books") {
				e := public.OrderBook{}
				if err := json.Unmarshal(data, &e); err != nil {
					return false
				}
				if c.obCh != nil {
					c.obCh <- &e
				}
				if c.StructuredEventChan != nil {
					c.StructuredEventChan <- e
				}
				return true
			}
		}
	}
	return false
}
