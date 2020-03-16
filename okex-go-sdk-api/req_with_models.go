package okex

import (
	"strconv"
	"time"

	"github.com/shopspring/decimal"
	"github.com/darkfoxs96/golimiter/limiter"
)

type SpotInstrument struct {
	BaseCurrency  string          `json:"base_currency"`
	QuoteCurrency string          `json:"quote_currency"` // Main coin
	Pair          string          `json:"instrument_id"`
	MinSize       decimal.Decimal `json:"min_size"`
	SizeIncrement decimal.Decimal `json:"size_increment"`
	TickSize      decimal.Decimal `json:"tick_size"`
}

type SpotInstrumentsList []*SpotInstrument

type BoolNum string
const (
	BoolNumTrue  = "1"
	BoolNumFalse = "0"
)
func (b BoolNum) IsTrue() bool {
	return b == BoolNumTrue
}

type AccountCurrency struct {
	CanDeposit    BoolNum         `json:"can_deposit"`  // "1" == true
	CanWithdraw   BoolNum         `json:"can_withdraw"` // "1" == true
	Currency      string          `json:"currency"`
	FullName      string          `json:"name"`
	MinWithdrawal decimal.Decimal `json:"min_withdrawal"`
}

type AccountCurrenciesList []*AccountCurrency

type SpotInstrumentTicker struct {
	Ask            decimal.Decimal `json:"ask"`
	Bid            decimal.Decimal `json:"bid"`
	BestAsk        decimal.Decimal `json:"best_ask"`
	BestBid        decimal.Decimal `json:"best_bid"`
	BestAskSize    decimal.Decimal `json:"best_ask_size"`
	BestBidSize    decimal.Decimal `json:"best_bid_size"`
	BaseVolume24h  decimal.Decimal `json:"base_volume_24h"`
	QuoteVolume24h decimal.Decimal `json:"quote_volume_24h"` // In main coin (like total)
	High24h        decimal.Decimal `json:"high_24h"`
	Low24h         decimal.Decimal `json:"low_24h"`
	Open24h        decimal.Decimal `json:"open_24h"`
	Last           decimal.Decimal `json:"last"`
	LastQty        decimal.Decimal `json:"last_qty"`
	Timestamp      time.Time       `json:"timestamp"`
	Pair           string          `json:"instrument_id"`
	ProductID      string          `json:"product_id"` // TODO instrument_id == product_id ?
}

type SpotInstrumentsTickerList []*SpotInstrumentTicker

type SpotAccountBalance struct {
	Frozen    decimal.Decimal `json:"frozen"`
	Hold      decimal.Decimal `json:"hold"` // Amount on hold (not available)
	AccountID string          `json:"id"`
	Currency  string          `json:"currency"`
	Balance   decimal.Decimal `json:"balance"` // Remaining balance
	Available decimal.Decimal `json:"available"` // Available amount
	Holds     decimal.Decimal `json:"holds"` // TODO ?
}

type SpotAccountBalancesList []*SpotAccountBalance

type SpotOrderResponse struct {
	OrderID       string `json:"order_id"` // TODO may use like a number?
	ClientOrderID string `json:"client_oid"`
	ErrorMessage  string `json:"error_message"`
	ErrorCode     string `json:"error_code"`
	Result        bool   `json:"result"`
}

type SpotNewOrderResponse struct {
	*SpotOrderResponse
}

type SpotCancelOrderResponse struct {
	*SpotOrderResponse
}

type OrderSide string
const (
	OrderSideBuy  OrderSide = "buy"
	OrderSideSell OrderSide = "sell"
)

func (o OrderSide) IsBuy() bool {
	return o == OrderSideBuy
}

type OrderType string
const (
	OrderTypeLimit  OrderType = "limit"
	OrderTypeMarket OrderType = "market"
)

type OrderStrategy string
const (
	OrderStrategyNormal   OrderStrategy = "0"
	OrderStrategyPostOnly OrderStrategy = "1"
	OrderStrategyFOK 	  OrderStrategy = "2"
	OrderStrategyIOC 	  OrderStrategy = "3"
)

type OrderStatus string
const (
	OrderStatusFailed   	   OrderStatus = "-2"
	OrderStatusCanceled   	   OrderStatus = "-1"
	OrderStatusOpen   		   OrderStatus = "0"
	OrderStatusPartiallyFilled OrderStatus = "1"
	OrderStatusFullyFilled 	   OrderStatus = "2"
	OrderStatusSubmitting 	   OrderStatus = "3"
	OrderStatusCanceling 	   OrderStatus = "4"
)

type Order struct {
	OrderID        string          `json:"order_id"` // TODO may use like a number?
	ClientOrderID  string          `json:"client_oid"`
	Timestamp      time.Time       `json:"timestamp"` // TODO or use 'created_at'?
	Price          decimal.Decimal `json:"price"`
	PriceAvg       decimal.Decimal `json:"price_avg"`
	Quantity       decimal.Decimal `json:"size"`
	QuantityFilled decimal.Decimal `json:"filled_size"`
	Total          decimal.Decimal `json:"filled_notional"` // TODO sure?
	Notional       string          `json:"notional"`        // TODO ? buy (for market orders)
	Pair           string          `json:"instrument_id"`
	Type           OrderType       `json:"type"`       // Order type: limit or market (default: limit)
	Side           OrderSide       `json:"side"`       // 'buy' or 'sell'
	Strategy       OrderStrategy   `json:"order_type"` // Specify 0: Normal order (Unfilled and 0 imply normal limit order) 1: Post only 2: Fill or Kill 3: Immediate Or Cancel
	Status         OrderStatus     `json:"state"`      // Order Status: -2 = Failed -1 = Canceled 0 = Open 1 = Partially Filled 2 = Fully Filled 3 = Submitting 4 = Canceling
}

type OrdersList []*Order

type TradeFee struct {
	Maker     decimal.Decimal `json:"maker"`
	Taker     decimal.Decimal `json:"taker"`
	Timestamp time.Time       `json:"timestamp"`
}

type WithdrawalFee struct {
	Currency string          `json:"currency"`
	MaxFee   decimal.Decimal `json:"max_fee"`
	MinFee   decimal.Decimal `json:"min_fee"`
}

type WithdrawalFeesList []*WithdrawalFee

type OrderUpdateWS struct {
	*Order
	LastFillPrice   decimal.Decimal `json:"last_fill_px"`   // Latest Filled Price. '0' will be returned if the data is empty
	LastFillQty     decimal.Decimal `json:"last_fill_qty"`  // Latest Filled Volume. '0' will be returned if the data is empty.
	LastFillTradeID decimal.Decimal `json:"last_fill_id"`   // Trade id. '0' will be returned if the data is empty
	LastFillTime    time.Time       `json:"last_fill_time"` // Latest Filled Time. The '1970-01-01T00:00:00.000Z' will be returned if the data is empty.
	CreatedAt       time.Time       `json:"created_at"`     // date created order. Now 'Timestamp' this is date event
}

type WSEventTable string
const (
	WSEventTableSpotOrder 	WSEventTable = "spot/order"
	WSEventTableSpotAccount WSEventTable = "spot/account"
	WSEventTableSpotDepth 	WSEventTable = "spot/depth"
)

type WSEventAction string
const (
	WSEventActionPartial WSEventAction = "partial"
	WSEventActionUpdate  WSEventAction = "update"
)

type WSEvent struct {
	Table  WSEventTable  `json:"table"`
	Action WSEventAction `json:"action"`
	Data   []byte        `json:"data"`
}

type UserOrdersWS struct {
	Table WSEventTable     `json:"table"` // WSEventTableSpotOrder
	Data  []*OrderUpdateWS `json:"data"`
}

type UserSpotAccountWS struct {
	Table WSEventTable            `json:"table"` // WSEventTableSpotAccount
	Data  SpotAccountBalancesList `json:"data"`
}

type Depth400PartialWS struct {
	Table  WSEventTable    `json:"table"`  // WSEventTableSpotDepth
	Action WSEventAction   `json:"action"` // WSEventActionPartial
	Data   []*Depth400Data `json:"data"`
}

type Depth400UpdateWS struct {
	Table  WSEventTable    `json:"table"`  // WSEventTableSpotDepth
	Action WSEventAction   `json:"action"` // WSEventActionUpdate
	Data   []*Depth400Data `json:"data"`
}

type Depth400Data struct {
	Pair      string               `json:"instrument_id"`
	Asks      [][3]decimal.Decimal `json:"asks"` // bids and asks value example: In ["411.8","10","8"], 411.8 is price depth, 10 is the amount at the price, 8 is the number of orders at the price.
	Bids      [][3]decimal.Decimal `json:"bids"` // bids and asks value example: In ["411.8","10","8"], 411.8 is price depth, 10 is the amount at the price, 8 is the number of orders at the price.
	Timestamp time.Time            `json:"timestamp"`
	Checksum  int64                `json:"checksum"`
}

var ( // LIMIT
	spotAccountListLimit  = &limiter.Limiter{Limit: 10, PeriodMillisecond: 1_200}
	spotTickersLimit	  = &limiter.Limiter{Limit: 10, PeriodMillisecond: 1_200}
	spotInstrumentsLimit  = &limiter.Limiter{Limit: 10, PeriodMillisecond: 1_200}
	accountCurrencyLimit  = &limiter.Limiter{Limit: 3,  PeriodMillisecond: 1_200}
	spotPlaceOrderLimit	  = &limiter.Limiter{Limit: 50, PeriodMillisecond: 1_200}
	spotCancelOrderLimit  = &limiter.Limiter{Limit: 50, PeriodMillisecond: 1_200}
	spotOrderListLimit	  = &limiter.Limiter{Limit: 10, PeriodMillisecond: 1_200}
	spotOrderDetailsLimit = &limiter.Limiter{Limit: 10, PeriodMillisecond: 1_200}
	spotTradeFeeLimit 	  = &limiter.Limiter{Limit: 1,  PeriodMillisecond: 11_000}
	withdrawalFeeLimit    = &limiter.Limiter{Limit: 5,  PeriodMillisecond: 1_200}
)

func (client *Client) LimitedGetSpotAccounts() (SpotAccountBalancesList, error) {
	r := SpotAccountBalancesList{}
	ch := make(chan interface{}, 2)
	var err error

	spotAccountListLimit.Wait(func() {
		_, err = client.Request(GET, SPOT_ACCOUNTS, nil, &r)
		ch <- nil
	})

	<-ch
	return r, err
}

func (client *Client) LimitedGetSpotInstrumentsTicker() (SpotInstrumentsTickerList, error) {
	r := SpotInstrumentsTickerList{}
	ch := make(chan interface{}, 2)
	var err error

	spotTickersLimit.Wait(func() {
		_, err = client.Request(GET, SPOT_INSTRUMENTS_TICKER, nil, &r)
		ch <- nil
	})

	<-ch
	return r, err
}

func (client *Client) LimitedGetSpotInstruments() (SpotInstrumentsList, error) {
	r := SpotInstrumentsList{}
	ch := make(chan interface{}, 2)
	var err error

	spotInstrumentsLimit.Wait(func() {
		_, err = client.Request(GET, SPOT_INSTRUMENTS, nil, &r)
		ch <- nil
	})

	<-ch
	return r, err
}

func (client *Client) LimitedGetAccountCurrencies() (AccountCurrenciesList, error) {
	r := AccountCurrenciesList{}
	ch := make(chan interface{}, 2)
	var err error

	accountCurrencyLimit.Wait(func() {
		_, err = client.Request(GET, ACCOUNT_CURRENCIES, nil, &r)
		ch <- nil
	})

	<-ch
	return r, err
}

func (client *Client) LimitedGetAllSpotOrders(status OrderStatus, pair string) (list OrdersList, err error) {
	var data OrdersList
	limit := 100
	limitStr := strconv.Itoa(limit)
	afterID, beforeID := "", ""

	for {
		data, err = client.LimitedGetSpotOrders(status, pair, afterID, beforeID, limitStr)
		if err != nil {
			return
		}
		list = append(list, data...)

		if len(data) < limit {
			break
		}

		afterID = data[len(data) - 1].OrderID
	}

	return
}

func (client *Client) LimitedGetSpotOrders(status OrderStatus, pair string, afterID string, beforeID string, limit string) (OrdersList, error) {
	r := OrdersList{}
	ch := make(chan interface{}, 2)
	var err error

	options := NewParams()
	options["instrument_id"] = pair
	options["state"] = string(status)
	options["after"] = afterID
	options["before"] = beforeID
	options["limit"] = limit
	uri := BuildParams(SPOT_ORDERS, options)

	spotOrderListLimit.Wait(func() {
		_, err = client.Request(GET, uri, nil, &r)
		ch <- nil
	})

	<-ch
	return r, err
}

func (client *Client) LimitedGetSpotOrdersByID(pair, orderOrClientID string) (*Order, error) {
	r := &Order{}
	ch := make(chan interface{}, 2)
	var err error
	uri := SPOT_ORDERS + "/" + orderOrClientID
	options := NewParams()
	options["instrument_id"] = pair
	uri = BuildParams(uri, options)

	spotOrderDetailsLimit.Wait(func() {
		_, err = client.Request(GET, uri, nil, r)
		ch <- nil
	})

	<-ch
	return r, err
}

func (client *Client) LimitedSpotCancelOrders(pair, orderOrClientID string) (*SpotCancelOrderResponse, error) {
	r := &SpotCancelOrderResponse{}
	ch := make(chan interface{}, 2)
	var err error

	uri := "/api/spot/v3/cancel_orders/" + orderOrClientID
	options := NewParams()
	options["instrument_id"] = pair

	spotCancelOrderLimit.Wait(func() {
		_, err = client.Request(POST, uri, options, r)
		ch <- nil
	})

	<-ch
	return r, err
}

func (client *Client) LimitedPostSpotOrders(side OrderSide, orderType OrderType, strategy OrderStrategy, pair string, price, quantity decimal.Decimal, notional string) (*SpotNewOrderResponse, error) {
	r := &SpotNewOrderResponse{}
	ch := make(chan interface{}, 2)
	var err error

	postParams := NewParams()
	postParams["side"] = string(side)
	postParams["instrument_id"] = pair
	postParams["type"] = string(orderType)
	postParams["order_type"] = string(strategy)

	if orderType == OrderTypeLimit {
		postParams["price"] = price.String()
		postParams["size"] = quantity.String()
	} else {
		postParams["size"] = quantity.String()
		postParams["notional"] = notional
	}

	spotPlaceOrderLimit.Wait(func() {
		_, err = client.Request(POST, SPOT_ORDERS, postParams, &r)
		ch <- nil
	})

	<-ch
	return r, err
}

func (client *Client) LimitedGetSpotTradeFee() (*TradeFee, error) {
	r := &TradeFee{}
	ch := make(chan interface{}, 2)
	var err error

	spotTradeFeeLimit.Wait(func() {
		_, err = client.Request(GET, "/api/spot/v3/trade_fee", nil, r)
		ch <- nil
	})

	<-ch
	return r, err
}

// LimitedGetWithdrawalFee arg currency not required
func (client *Client) LimitedGetWithdrawalFee(currency string) (WithdrawalFeesList, error) {
	r := WithdrawalFeesList{}
	ch := make(chan interface{}, 2)
	var err error

	uri := ACCOUNT_WITHRAWAL_FEE
	if currency != "" {
		params := NewParams()
		params["currency"] = currency
		uri = BuildParams(uri, params)
	}

	withdrawalFeeLimit.Wait(func() {
		_, err = client.Request(GET, uri, nil, &r)
		ch <- nil
	})

	<-ch
	return r, err
}
