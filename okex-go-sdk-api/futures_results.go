package okex

/*
 OKEX futures contract api response results
 @author Tony Tian
 @date 2018-03-17
 @version 1.0.0
*/

//easyjson:json
type ServerTime struct {
	Iso   string `json:"iso"`
	Epoch string `json:"epoch"`
}

//easyjson:json
type ExchangeRate struct {
	InstrumentId string  `json:"instrument_id"`
	Rate         float64 `json:"rate,string"`
	Timestamp    string  `json:"timestamp"`
}

//easyjson:json
type BizWarmTips struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Msg     string `json:"msg"`
}

//easyjson:json
type ResultReq struct {
	Result bool `json:"result"`
}

//easyjson:json
type PageResult struct {
	From  int
	To    int
	Limit int
}

//easyjson:json
type FuturesPosition struct {
	BizWarmTips
	Result        bool `json:"result"`
	MarginMode    string
	CrossPosition []FuturesCrossPositionHolding
	FixedPosition []FuturesFixedPositionHolding
}

//easyjson:json
type FuturesCrossPosition struct {
	Result        bool                          `json:"result"`
	MarginMode    string                        `json:"margin_mode"`
	CrossPosition []FuturesCrossPositionHolding `json:"holding"`
}

//easyjson:json
type FuturesFixedPosition struct {
	Result        bool                          `json:"result"`
	MarginMode    string                        `json:"margin_mode"`
	FixedPosition []FuturesFixedPositionHolding `json:"holding"`
}

//easyjson:json
type FuturesCrossPositionHolding struct {
	FuturesPositionBase
	LiquidationPrice float64 `json:"liquidation_price,string"`
	Leverage         float64 `json:"leverage,string"`
}

//easyjson:json
type FuturesFixedPositionHolding struct {
	FuturesPositionBase
	LongMargin      float64 `json:"long_margin,string"`
	LongLiquiPrice  float64 `json:"long_liqui_price,string"`
	LongPnlRatio    float64 `json:"long_pnl_ratio,string"`
	LongLeverage    float64 `json:"long_leverage,string"`
	ShortMargin     float64 `json:"short_margin,string"`
	ShortLiquiPrice float64 `json:"short_liqui_price,string"`
	ShortPnlRatio   float64 `json:"short_pnl_ratio,string"`
	ShortLeverage   float64 `json:"short_leverage,string"`
}

//easyjson:json
type FuturesPositionBase struct {
	LongQty              float64 `json:"long_qty,string"`
	LongAvailQty         float64 `json:"long_avail_qty,string"`
	LongAvgCost          float64 `json:"long_avg_cost,string"`
	LongSettlementPrice  float64 `json:"long_settlement_price,string"`
	RealizedPnl          float64 `json:"realized_pnl,string"`
	ShortQty             float64 `json:"short_qty,string"`
	ShortAvailQty        float64 `json:"short_avail_qty,string"`
	ShortAvgCost         float64 `json:"short_avg_cost,string"`
	ShortSettlementPrice float64 `json:"short_settlement_price,string"`
	InstrumentId         string  `json:"instrument_id"`
	CreatedAt            string  `json:"created_at"`
	UpdatedAt            string  `json:"updated_at"`
}

//easyjson:json
type FuturesAccount struct {
	BizWarmTips
	Result       bool `json:"result"`
	MarginMode   string
	CrossAccount map[string]FuturesCrossAccount
	FixedAccount map[string]FuturesFixedAccount
}

//easyjson:json
type FuturesMarkdown struct {
	BizWarmTips
	InstrumentId string `json:"instrument_id"`
	Timestamp    string `json:"timestamp"`
	MarkPrice    string `json:"mark_price"`
}

//easyjson:json
type FuturesFixedAccountInfo struct {
	Result bool                           `json:"result"`
	Info   map[string]FuturesFixedAccount `json:"info"`
}

//easyjson:json
type FuturesCrossAccountInfo struct {
	Result bool                           `json:"result"`
	Info   map[string]FuturesCrossAccount `json:"info"`
}

//easyjson:json
type FuturesFixedAccount struct {
	MarginMode        string                         `json:"margin_mode"`
	Equity            float64                        `json:"equity,string"`
	TotalAvailBalance float64                        `json:"total_avail_balance,string"`
	Contracts         []FuturesFixedAccountContracts `json:"contracts"`
}

//easyjson:json
type FuturesFixedAccountContracts struct {
	AvailableQty      float64 `json:"available_qty,string"`
	FixedBalance      float64 `json:"fixed_balance,string"`
	InstrumentId      string  `json:"instrument_id"`
	MarginFixed       float64 `json:"margin_fixed,string"`
	MarginForUnfilled float64 `json:"margin_for_unfilled,string"`
	MarginFrozen      float64 `json:"margin_frozen,string"`
	RealizedPnl       float64 `json:"realized_pnl,string"`
	UnrealizedPnl     float64 `json:"unrealizedPnl,string"`
}

//easyjson:json
type FuturesCrossAccount struct {
	Equity            float64 `json:"equity,string"`
	Margin            float64 `json:"margin,string"`
	MarginMode        string  `json:"margin_mode"`
	MarginRatio       float64 `json:"margin_ratio,string"`
	RealizedPnl       float64 `json:"realized_pnl,string"`
	UnrealizedPnl     float64 `json:"unrealized_pnl,string"`
	TotalAvailBalance float64 `json:"total_avail_balance,string"`
}

//easyjson:json
type FuturesCurrencyAccount struct {
	BizWarmTips
	Result       bool `json:"result"`
	MarginMode   string
	CrossAccount FuturesCrossAccount
	FixedAccount FuturesFixedAccount
}

//easyjson:json
type FuturesCurrencyLedger struct {
	LedgerId  int64                        `json:"ledger_id,string"`
	Amount    float64                      `json:"amount,string"`
	Balance   float64                      `json:"balance,string"`
	Currency  string                       `json:"currency"`
	Type      string                       `json:"type"`
	Timestamp string                       `json:"timestamp"`
	Details   FuturesCurrencyLedgerDetails `json:"details"`
}

//easyjson:json
type FuturesCurrencyLedgerDetails struct {
	OrderId      int64  `json:"order_id"`
	InstrumentId string `json:"instrument_id"`
}

//easyjson:json
type FuturesAccountsHolds struct {
	InstrumentId string  `json:"instrument_id"`
	Amount       float64 `json:"amount,string"`
	Timestamp    string  `json:"timestamp"`
}

//easyjson:json
type FuturesNewOrderResult struct {
	BizWarmTips
	Result    bool   `json:"result"`
	ClientOid string `json:"client_oid"`
	OrderId   int64  `json:"order_id,string"`
}

//easyjson:json
type FuturesBatchNewOrderResult struct {
	Result    bool        `json:"result"`
	OrderInfo []OrderInfo `json:"order_info"`
}

//easyjson:json
type CodeMessage struct {
	ErrorCode    int64  `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

/*
  If OrderId = -1, ErrorCode > 0, error order
*/
//easyjson:json
type OrderInfo struct {
	ClientOid string `json:"client_oid"`
	OrderId   string `json:"order_id"`
	CodeMessage
}

//easyjson:json
type FuturesCancelInstrumentOrderResult struct {
	Result       bool   `json:"result"`
	OrderId      string `json:"order_id"`
	InstrumentId string `json:"instrument_id"`
}

//easyjson:json
type FuturesBatchCancelInstrumentOrdersResult struct {
	Result       bool     `json:"result"`
	OrderIds     []string `json:"order_ids"`
	InstrumentId string   `json:"instrument_id"`
}

//easyjson:json
type FuturesClosePositionResult struct {
	Result            bool                `json:"result"`
	ClosePositionInfo []ClosePositionInfo `json:"close_position_info"`
}

//easyjson:json
type ClosePositionInfo struct {
	InstrumentId string `json:"instrument_id"`
	CodeMessage
}

//easyjson:json
type FuturesGetOrdersResult struct {
	Result bool                    `json:"result"`
	Orders []FuturesGetOrderResult `json:"order_info"`
}

//easyjson:json
type FuturesGetOrderResult struct {
	InstrumentId string  `json:"instrument_id"`
	Size         int64   `json:"size,string"`
	Timestamp    string  `json:"timestamp"`
	FilledQty    float64 `json:"filled_qty,string"`
	Fee          float64 `json:"fee,string"`
	OrderId      int64   `json:"order_id,string"`
	Price        float64 `json:"price,string"`
	PriceAvg     float64 `json:"price_avg,string"`
	Status       int     `json:"status,string"`
	Type         int     `json:"type,string"`
	ContractVal  float64 `json:"contract_val,string"`
	Leverage     float64 `json:"leverage,string"`
}

//easyjson:json
type FuturesFillResult struct {
	TradeId      int64   `json:"trade_id,string"`
	InstrumentId string  `json:"instrument_id"`
	Price        float64 `json:"price,string"`
	OrderQty     float64 `json:"order_qty,string"`
	OrderId      int64   `json:"order_id,string"`
	CreatedAt    string  `json:"created_at"`
	ExecType     string  `json:"exec_type"`
	Fee          float64 `json:"fee,string"`
	Side         string  `json:"side"`
}

//easyjson:json
type FuturesUsersSelfTrailingVolumesResult struct {
	FuturesUsersSelfTrailingVolumeResult []FuturesUsersSelfTrailingVolumeResult
}

//easyjson:json
type FuturesUsersSelfTrailingVolumeResult struct {
	InstrumentId   string  `json:"instrument_id"`
	ExchangeVolume float64 `json:"exchange_volume,string"`
	Volume         float64 `json:"volume,string"`
	RecordedAt     string  `json:"recorded_at"`
}

//easyjson:json
type FuturesInstrumentsResult struct {
	InstrumentId    string  `json:"instrument_id"`
	UnderlyingIndex string  `json:"underlying_index"`
	QuoteCurrency   string  `json:"quote_currency"`
	TickSize        float64 `json:"tick_size,string"`
	ContractVal     float64 `json:"contract_val,string"`
	Listing         string  `json:"listing"`
	Delivery        string  `json:"delivery"`
	TradeIncrement  float64 `json:"trade_increment,string"`
}

//easyjson:json
type FuturesInstrumentCurrenciesResult struct {
	Id      int64   `json:"id,string"`
	Name    string  `json:"name"`
	MinSize float64 `json:"min_size,string"`
}

//easyjson:json
type FuturesInstrumentBookResult struct {
	Asks      [][]string `json:"asks,string"`
	Bids      [][]string `json:"bids,string"`
	Timestamp string     `json:"timestamp"`
}

//easyjson:json
type FuturesInstrumentTickerResult struct {
	InstrumentId string  `json:"instrument_id"`
	BestBid      float64 `json:"best_bid,string"`
	BestAsk      float64 `json:"best_ask,string"`
	High24h      float64 `json:"high_24h,string"`
	Low24h       float64 `json:"low_24h,string"`
	Last         float64 `json:"last,string"`
	Volume24h    float64 `json:"volume_24h,string"`
	Timestamp    string  `json:"timestamp"`
}

//easyjson:json
type FuturesInstrumentTradesResult struct {
	TradeId   string  `json:"trade_id"`
	Side      string  `json:"side"`
	Price     float64 `json:"price,string"`
	Qty       float64 `json:"qty,string"`
	Timestamp string  `json:"timestamp"`
}

//easyjson:json
type FuturesInstrumentIndexResult struct {
	InstrumentId string  `json:"instrument_id"`
	Index        float64 `json:"index,string"`
	Timestamp    string  `json:"timestamp"`
}

//easyjson:json
type FuturesInstrumentEstimatedPriceResult struct {
	InstrumentId    string  `json:"instrument_id"`
	SettlementPrice float64 `json:"settlement_price,string"`
	Timestamp       string  `json:"timestamp"`
}

//easyjson:json
type FuturesInstrumentOpenInterestResult struct {
	InstrumentId string `json:"instrument_id"`
	Amount       int64  `json:"amount,string"`
	Timestamp    string `json:"timestamp"`
}

//easyjson:json
type FuturesInstrumentPriceLimitResult struct {
	InstrumentId string  `json:"instrument_id"`
	Highest      float64 `json:"highest,string"`
	Lowest       float64 `json:"lowest,string"`
	Timestamp    string  `json:"timestamp"`
}

//easyjson:json
type FuturesInstrumentLiquidationListResult struct {
	Page            PageResult
	LiquidationList []FuturesInstrumentLiquidationResult
}

//easyjson:json
type FuturesInstrumentLiquidationResult struct {
	InstrumentId string `json:"instrument_id"`
	Price        string `json:"price"`
	Size         string `json:"size"`
	Loss         string `json:"loss"`
	CreatedAt    string `json:"created_at"`
}
