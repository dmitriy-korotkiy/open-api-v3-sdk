package okex

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/darkfoxs96/golimiter/limiter"
	"github.com/shopspring/decimal"
)

type StateStatus string

const (
	StateStatusLive    StateStatus = "live"
	StateStatusSuspend StateStatus = "suspend"
	StateStatusPreopen StateStatus = "preopen"
)

//easyjson:json
type SpotInstrument struct {
	BaseCurrency  string          `json:"baseCcy"`
	QuoteCurrency string          `json:"quoteCcy"` // Main coin
	Pair          string          `json:"instId"`
	MinSize       decimal.Decimal `json:"minSz"`
	SizeIncrement decimal.Decimal `json:"lotSz"`
	TickSize      decimal.Decimal `json:"tickSz"`
	State         StateStatus     `json:"state"`
}

//easyjson:json
type SpotInstrumentsList []*SpotInstrument

//easyjson:json
type SpotInstrumentsListWrapper struct {
	Data SpotInstrumentsList `json:"data"`
}

type BoolNum string

const (
	BoolNumTrue  = "1"
	BoolNumFalse = "0"
)

func (b BoolNum) IsTrue() bool {
	return b == BoolNumTrue
}

//easyjson:json
type AccountCurrency struct {
	CanDeposit    bool            `json:"canDep"`
	CanWithdraw   bool            `json:"canWd"`
	CanInternal   bool            `json:"canInternal"`
	Currency      string          `json:"ccy"`
	FullName      string          `json:"name"` // TODO: where?
	MinWithdrawal decimal.Decimal `json:"minWd"`
	MaxFee        decimal.Decimal `json:"maxFee"`
	MinFee        decimal.Decimal `json:"minFee"`
}

//easyjson:json
type AccountCurrenciesList []*AccountCurrency

//easyjson:json
type AccountCurrenciesListWrapper struct {
	Data AccountCurrenciesList `json:"data"`
}

//easyjson:json
type SpotInstrumentTicker struct {
	BestAsk        decimal.Decimal `json:"askPx"`
	BestBid        decimal.Decimal `json:"bidPx"`
	BestAskSize    decimal.Decimal `json:"askSz"`
	BestBidSize    decimal.Decimal `json:"bidSz"`
	BaseVolume24h  decimal.Decimal `json:"vol24h"`
	QuoteVolume24h decimal.Decimal `json:"volCcy24h"` // In main coin (like total)
	High24h        decimal.Decimal `json:"high24h"`
	Low24h         decimal.Decimal `json:"low24h"`
	Open24h        decimal.Decimal `json:"open24h"`
	Last           decimal.Decimal `json:"last"`
	LastQty        decimal.Decimal `json:"lastSz"`
	Timestamp      string          `json:"ts"` // milliseconds
	Pair           string          `json:"instId"`
	Type           string          `json:"instType"`
}

//easyjson:json
type SpotInstrumentsTickerList []*SpotInstrumentTicker

//easyjson:json
type SpotInstrumentsTickerListWrapper struct {
	Data SpotInstrumentsTickerList `json:"data"`
}

//easyjson:json
type SpotAccountBalance struct {
	Details []struct {
		Frozen    decimal.Decimal `json:"frozenBal"`
		Currency  string          `json:"ccy"`
		EqUSD     decimal.Decimal `json:"eqUsd"`
		Balance   decimal.Decimal `json:"cashBal"`  // Remaining balance
		Available decimal.Decimal `json:"availBal"` // Available amount
	} `json:"details"`
}

//easyjson:json
type SpotAccountBalancesList []*SpotAccountBalance

//easyjson:json
type SpotAccountBalancesListWrapper struct {
	Data SpotAccountBalancesList `json:"data"`
}

//easyjson:json
type SpotOrderResponse struct {
	OrderID       string `json:"ordId"`
	ClientOrderID string `json:"clOrdId"`
	SCode         string `json:"sCode"`
	SMsg          string `json:"sMsg"`
	Tag           string `json:"tag"`
}

//easyjson:json
type SpotNewOrderResponse struct {
	Code string               `json:"code"` // good = 1
	Msg  string               `json:"msg"`
	Data []*SpotOrderResponse `json:"data"`
}

//easyjson:json
type SpotCancelOrderResponseWrapper struct {
	Code string               `json:"code"` // good = 1
	Data []*SpotOrderResponse `json:"data"`
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
	OrderTypeLimit           OrderType = "limit"
	OrderTypeMarket          OrderType = "market"
	OrderTypeFOK             OrderType = "fok" // TODO: why without errors if not full filled?
	OrderTypeIOC             OrderType = "ioc"
	OrderTypePostOnly        OrderType = "post_only"
	OrderTypeOptimalLimitIOC OrderType = "optimal_limit_ioc"
)

type OrderStrategy string

const (
	OrderStrategyNormal   OrderStrategy = "0"
	OrderStrategyPostOnly OrderStrategy = "1"
	OrderStrategyFOK      OrderStrategy = "2"
	OrderStrategyIOC      OrderStrategy = "3"
)

type OrderStatus string

const (
	OrderStatusCanceled        OrderStatus = "canceled"
	OrderStatusOpen            OrderStatus = "live"
	OrderStatusPartiallyFilled OrderStatus = "partially_filled"
	OrderStatusFullyFilled     OrderStatus = "filled"
)

//easyjson:json
type Order struct {
	OrderID        string          `json:"ordId"`
	ClientOrderID  string          `json:"clOrdId"`
	Timestamp      string          `json:"cTime"` // Creation time, Unix timestamp format in milliseconds, e.g. 1597026383085
	Price          decimal.Decimal `json:"px"`
	PriceAvgStr    string          `json:"avgPx"` // Average filled price. If none is filled, it will return 0.
	PriceAvg       decimal.Decimal // Average filled price. If none is filled, it will return 0.
	Quantity       decimal.Decimal `json:"sz"`
	QuantityFilled decimal.Decimal `json:"accFillSz"`
	Total          decimal.Decimal
	ProfitAndLoss  decimal.Decimal `json:"pnl"`
	Fee            decimal.Decimal `json:"fee"`
	FeeCurrency    string          `json:"feeCcy"`
	Pair           string          `json:"instId"`
	Tag            string          `json:"tag"`
	PosSide        string          `json:"posSide"` // net, long, ...
	Type           OrderType       `json:"ordType"`
	Side           OrderSide       `json:"side"`  // 'buy' or 'sell'
	Status         OrderStatus     `json:"state"` // Order Status: -2 = Failed -1 = Canceled 0 = Open 1 = Partially Filled 2 = Fully Filled 3 = Submitting 4 = Canceling
}

//easyjson:json
type OrdersList []*Order

//easyjson:json
type OrdersListWrapper struct {
	Data OrdersList `json:"data"`
}

//easyjson:json
type TradeFee struct {
	Maker     decimal.Decimal `json:"maker"` // can be less than 0
	Taker     decimal.Decimal `json:"taker"` // can be less than 0
	Timestamp string          `json:"ts"`    // milliseconds
}

//easyjson:json
type TradeFeeListWrapper struct {
	Data []*TradeFee `json:"data"`
}

//easyjson:json
type WithdrawalFee struct {
	Currency      string          `json:"ccy"`
	MinWithdrawal decimal.Decimal `json:"minWd"`
	MaxFee        decimal.Decimal `json:"maxFee"`
	MinFee        decimal.Decimal `json:"minFee"`
}

//easyjson:json
type WithdrawalFeesList []*WithdrawalFee

//easyjson:json
type WithdrawalFeesListWrapper struct {
	Data WithdrawalFeesList `json:"data"`
}

//easyjson:json
type OrderUpdateWS struct {
	*Order
	// if need, should add string fields
	//LastFillPrice   decimal.Decimal `json:"fillPx"`   // Latest Filled Price. '0' will be returned if the data is empty
	//LastFillQty     decimal.Decimal `json:"fillSz"`   // Latest Filled Volume. '0' will be returned if the data is empty.
	//LastFillTradeID decimal.Decimal `json:"tradeId"`  // Trade id. '0' will be returned if the data is empty
	//LastFillTime    string          `json:"fillTime"` // 1597026383085
}

type WSEventTable string

const (
	WSEventTableSpotOrder   WSEventTable = "spot/order"
	WSEventTableSpotAccount WSEventTable = "spot/account"
	WSEventTableSpotDepth   WSEventTable = "spot/depth"
)

type WSEventAction string

const (
	WSEventActionPartial WSEventAction = "partial"
	WSEventActionUpdate  WSEventAction = "update"
)

//easyjson:json
type WSEvent struct {
	Table  WSEventTable    `json:"table"`
	Action WSEventAction   `json:"action"`
	Data   json.RawMessage `json:"data"`
}

//easyjson:json
type UserOrdersWS struct {
	Arg  *BaseOpSubscriptionArgs `json:"arg"`
	Data []*OrderUpdateWS        `json:"data"`
}

//easyjson:json
type UserSpotAccountWS struct {
	Arg  *BaseOpSubscriptionArgs `json:"arg"`
	Data SpotAccountBalancesList `json:"data"`
}

//easyjson:json
type Depth400PartialWS struct {
	Table  WSEventTable    `json:"table"`  // WSEventTableSpotDepth
	Action WSEventAction   `json:"action"` // WSEventActionPartial
	Data   []*Depth400Data `json:"data"`
}

//easyjson:json
type Depth400UpdateWS struct {
	Table  WSEventTable    `json:"table"`  // WSEventTableSpotDepth
	Action WSEventAction   `json:"action"` // WSEventActionUpdate
	Data   []*Depth400Data `json:"data"`
}

//easyjson:json
type Depth400Data struct {
	Pair      string               `json:"instrument_id"`
	Asks      [][3]decimal.Decimal `json:"asks"` // bids and asks value example: In ["411.8","10","8"], 411.8 is price depth, 10 is the amount at the price, 8 is the number of orders at the price.
	Bids      [][3]decimal.Decimal `json:"bids"` // bids and asks value example: In ["411.8","10","8"], 411.8 is price depth, 10 is the amount at the price, 8 is the number of orders at the price.
	Timestamp time.Time            `json:"timestamp"`
	Checksum  int64                `json:"checksum"`
}

var ( // LIMIT
	spotAccountListLimit  = &limiter.Limiter{Limit: 10, PeriodMillisecond: 1_200}
	spotTickersLimit      = &limiter.Limiter{Limit: 10, PeriodMillisecond: 1_200}
	spotInstrumentsLimit  = &limiter.Limiter{Limit: 10, PeriodMillisecond: 1_200}
	accountCurrencyLimit  = &limiter.Limiter{Limit: 3, PeriodMillisecond: 1_200}
	spotPlaceOrderLimit   = &limiter.Limiter{Limit: 50, PeriodMillisecond: 1_200}
	spotCancelOrderLimit  = &limiter.Limiter{Limit: 50, PeriodMillisecond: 1_200}
	spotOrderListLimit    = &limiter.Limiter{Limit: 10, PeriodMillisecond: 1_200}
	spotOrderDetailsLimit = &limiter.Limiter{Limit: 10, PeriodMillisecond: 1_200}
	spotTradeFeeLimit     = &limiter.Limiter{Limit: 1, PeriodMillisecond: 11_000}
	withdrawalFeeLimit    = &limiter.Limiter{Limit: 5, PeriodMillisecond: 1_200}
)

func (client *Client) LimitedGetSpotAccounts() (SpotAccountBalancesList, error) {
	r := SpotAccountBalancesListWrapper{}
	ch := make(chan interface{}, 2)
	var err error

	spotAccountListLimit.Wait(func() {
		_, err = client.Request(GET, SPOT_ACCOUNTS, nil, &r)
		ch <- nil
	})

	<-ch
	return r.Data, err
}

func (client *Client) LimitedGetSpotInstrumentsTicker() (SpotInstrumentsTickerList, error) {
	r := SpotInstrumentsTickerListWrapper{}
	ch := make(chan interface{}, 2)
	var err error

	options := NewParams()
	options["instType"] = "SPOT"
	uri := BuildParams(SPOT_INSTRUMENTS_TICKER, options)

	spotTickersLimit.Wait(func() {
		_, err = client.Request(GET, uri, nil, &r)
		ch <- nil
	})

	<-ch
	return r.Data, err
}

func (client *Client) LimitedGetSpotInstrumentTicker(instID string) (SpotInstrumentsTickerList, error) {
	r := SpotInstrumentsTickerListWrapper{}
	ch := make(chan interface{}, 2)
	var err error

	options := NewParams()
	options["instId"] = instID
	uri := BuildParams(SPOT_INSTRUMENT_TICKER, options)

	spotTickersLimit.Wait(func() {
		_, err = client.Request(GET, uri, nil, &r)
		ch <- nil
	})

	<-ch
	return r.Data, err
}

func (client *Client) LimitedGetSpotInstruments() (SpotInstrumentsList, error) {
	r := SpotInstrumentsListWrapper{}
	ch := make(chan interface{}, 2)
	var err error

	options := NewParams()
	options["instType"] = "SPOT"
	uri := BuildParams(SPOT_INSTRUMENTS, options)

	spotInstrumentsLimit.Wait(func() {
		_, err = client.Request(GET, uri, nil, &r)
		ch <- nil
	})

	<-ch
	return r.Data, err
}

func (client *Client) LimitedGetAccountCurrencies() (AccountCurrenciesList, error) {
	r := AccountCurrenciesListWrapper{}
	ch := make(chan interface{}, 2)
	var err error

	accountCurrencyLimit.Wait(func() {
		_, err = client.Request(GET, ACCOUNT_CURRENCIES, nil, &r)
		ch <- nil
	})

	<-ch
	return r.Data, err
}

func (client *Client) LimitedGetAllSpotOpenedOrders(status OrderStatus, pair string, isOpened bool) (list OrdersList, err error) {
	var data OrdersList
	limit := 100
	limitStr := strconv.Itoa(limit)
	afterID, beforeID := "", ""

	for {
		data, err = client.LimitedGetSpotOpenedOrders(status, pair, afterID, beforeID, limitStr, isOpened)
		if err != nil {
			return
		}
		list = append(list, data...)

		if len(data) < limit {
			break
		}

		afterID = data[len(data)-1].OrderID
	}

	return
}

func (client *Client) LimitedGetSpotOpenedOrders(status OrderStatus, pair string, afterID string, beforeID string, limit string, isOpened bool) (OrdersList, error) {
	r := OrdersListWrapper{}
	ch := make(chan interface{}, 2)
	var err error

	options := NewParams()
	options["instType"] = "SPOT"
	options["instId"] = pair
	options["state"] = string(status)
	options["after"] = afterID
	options["before"] = beforeID
	options["limit"] = limit
	url := SPOT_ORDERS_OPENED
	if !isOpened { // TODO: closed for the last 3 months
		url = SPOT_ORDERS_CLOSED
	}
	uri := BuildParams(url, options)

	spotOrderListLimit.Wait(func() {
		_, err = client.Request(GET, uri, nil, &r)
		ch <- nil
	})

	<-ch

	for _, order := range r.Data {
		err = calcOrderFields(order)
		if err != nil {
			return nil, err
		}
	}
	return r.Data, err
}

func (client *Client) LimitedGetSpotOrdersByID(pair, orderID, clientOrderID string) (*Order, error) {
	r := &OrdersListWrapper{}
	ch := make(chan interface{}, 2)
	var err error

	options := NewParams()
	options["instId"] = pair
	if orderID != "" {
		options["ordId"] = orderID
	}
	if clientOrderID != "" {
		options["clOrdId"] = clientOrderID
	}
	uri := BuildParams(SPOT_ORDERS_DETAILS, options)

	spotOrderDetailsLimit.Wait(func() {
		_, err = client.Request(GET, uri, nil, r)
		ch <- nil
	})

	<-ch

	order := r.Data[0]
	err = calcOrderFields(order)
	if err != nil {
		return nil, err
	}
	return order, err
}

func calcOrderFields(order *Order) (err error) {
	if order.PriceAvgStr != "" {
		order.PriceAvg, err = decimal.NewFromString(order.PriceAvgStr)
		if err != nil {
			return
		}
	}

	if order.PriceAvg.IsPositive() {
		order.Total = order.QuantityFilled.Mul(order.PriceAvg)
	} else if order.QuantityFilled.IsPositive() && order.Price.IsPositive() {
		order.Total = order.QuantityFilled.Mul(order.Price)
	}

	return
}

func calcTotal(priceAvg, quantityFilled, price decimal.Decimal) decimal.Decimal {
	if priceAvg.IsPositive() {
		return quantityFilled.Mul(priceAvg)
	} else if quantityFilled.IsPositive() && price.IsPositive() {
		return quantityFilled.Mul(price)
	} else {
		return decimal.Decimal{}
	}
}

func (client *Client) LimitedSpotCancelOrders(pair, orderID, clientOrderID string) (*SpotCancelOrderResponseWrapper, error) {
	r := &SpotCancelOrderResponseWrapper{}
	ch := make(chan interface{}, 2)
	var err error

	options := NewParams()
	options["instId"] = pair
	if orderID != "" {
		options["ordId"] = orderID
	}
	if clientOrderID != "" {
		options["clOrdId"] = clientOrderID
	}
	uri := BuildParams(SPOT_CANCEL_ORDERS_BY_ID, options)

	spotCancelOrderLimit.Wait(func() {
		_, err = client.Request(POST, uri, options, r)
		ch <- nil
	})

	<-ch
	return r, err
}

// LimitedPostSpotOrders - clOrdId can be empty
func (client *Client) LimitedPostSpotOrders(side OrderSide, orderType OrderType, pair string, price, quantity decimal.Decimal, clientOrderID string) (*SpotNewOrderResponse, error) {
	r := &SpotNewOrderResponse{}
	ch := make(chan interface{}, 2)
	var err error

	postParams := NewParams()
	postParams["side"] = string(side)
	postParams["instId"] = pair
	postParams["ordType"] = string(orderType)
	postParams["px"] = price.String()
	postParams["sz"] = quantity.String()
	postParams["tdMode"] = "cash"
	if clientOrderID != "" {
		postParams["clOrdId"] = clientOrderID
	}
	// posSide ? long/short

	spotPlaceOrderLimit.Wait(func() {
		_, err = client.Request(POST, PLEACE_SPOT_ORDERS, postParams, &r)
		ch <- nil
	})

	<-ch
	return r, err
}

func (client *Client) LimitedGetSpotTradeFee(category string) (*TradeFee, error) {
	r := &TradeFeeListWrapper{}
	ch := make(chan interface{}, 2)
	var err error

	options := NewParams()
	options["instType"] = "SPOT"
	options["category"] = category
	uri := BuildParams(SPOT_TRADE_FEE, options)

	spotTradeFeeLimit.Wait(func() {
		_, err = client.Request(GET, uri, nil, r)
		ch <- nil
	})

	<-ch
	return r.Data[0], err
}

func (client *Client) LimitedGetWithdrawalFee() (WithdrawalFeesList, error) {
	r := WithdrawalFeesListWrapper{}
	ch := make(chan interface{}, 2)
	var err error

	withdrawalFeeLimit.Wait(func() {
		_, err = client.Request(GET, ACCOUNT_WITHRAWAL_FEE, nil, &r)
		ch <- nil
	})

	<-ch
	return r.Data, err
}
