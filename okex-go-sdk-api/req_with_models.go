package okex

import (
	"time"

	"github.com/shopspring/decimal"
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

type AccountCurrency struct {
	CanDeposit    int8            `json:"can_deposit"` // 1 == true
	CanWithdraw   int8            `json:"can_withdraw"` // 1 == true
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
	AccountID int64           `json:"id"`
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
