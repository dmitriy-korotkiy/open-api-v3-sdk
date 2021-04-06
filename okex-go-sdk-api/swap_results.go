package okex

/*
 OKEX api result definition
 @author Lingting Fu
 @date 2018-12-27
 @version 1.0.0
*/

//easyjson:json
type SwapPositionHolding struct {
	LiquidationPrice string `json:"liquidation_price"`
	Position         string `json:"position"`
	AvailPosition    string `json:"avail_position"`
	AvgCost          string `json:"avg_cost"`
	SettlementPrice  string `json:"settlement_price"`
	InstrumentId     string `json:"instrument_id"`
	Leverage         string `json:"leverage"`
	RealizedPnl      string `json:"realized_pnl"`
	Side             string `json:"side"`
	Timestamp        string `json:"timestamp"`
	Margin           string `json:"margin";default:""`
}

//easyjson:json
type SwapPosition struct {
	BizWarmTips
	MarginMode string                `json:"margin_mode"`
	Holding    []SwapPositionHolding `json:"holding"`
}

//easyjson:json
type SwapPositionList []SwapPosition

//easyjson:json
type SwapAccountInfo struct {
	InstrumentId      string `json:"instrument_id"`
	Timestamp         string `json:"timestamp"`
	MarginFrozen      string `json:"margin_frozen"`
	TotalAvailBalance string `json:"total_avail_balance"`
	MarginRatio       string `json:"margin_ratio"`
	RealizedPnl       string `json:"realized_pnl"`
	UnrealizedPnl     string `json:"unrealized_pnl"`
	FixedBalance      string `json:"fixed_balance"`
	Equity            string `json:"equity"`
	Margin            string `json:"margin"`
	MarginMode        string `json:"margin_mode"`
}

//easyjson:json
type SwapAccounts struct {
	BizWarmTips
	Info []SwapAccountInfo `json:"info"`
}

//easyjson:json
type SwapAccount struct {
	Info SwapAccountInfo `json:"info"`
}

//easyjson:json
type BaseSwapOrderResult struct {
	OrderId      string `json:"order_id"`
	ClientOid    string `json:"client_oid"`
	ErrorMessage string `json:"error_message"`
	ErrorCode    string `json:"error_code"`
	Result       string `json:"result"`
}

//easyjson:json
type SwapOrderResult struct {
	BaseSwapOrderResult
	BizWarmTips
}

//easyjson:json
type SwapOrdersResult struct {
	BizWarmTips
	OrderInfo []BaseSwapOrderResult `json:"order_info"`
}

//easyjson:json
type SwapCancelOrderResult struct {
	ErrorMessage string `json:"error_message"`
	ErrorCode    string `json:"error_code"`
	OrderId      string `json:"order_id"`
	Result       string `json:"result"`
}

//easyjson:json
type SwapBatchCancelOrderResult struct {
	BizWarmTips
	InstrumentId string   `json:"instrument_id"`
	Ids          []string `json:"ids"`
	Result       string   `json:"result"`
}

//easyjson:json
type BaseOrderInfo struct {
	InstrumentId string `json:"instrument_id"`
	Status       string `json:"status"`
	OrderId      string `json:"order_id"`
	Timestamp    string `json:"timestamp"`
	Price        string `json:"price"`
	PriceAvg     string `json:"price_avg"`
	Size         string `json:"size"`
	Fee          string `json:"fee"`
	FilledQty    string `json:"filled_qty"`
	ContractVal  string `json:"contract_val"`
	Type         string `json:"type"`
}

//easyjson:json
type SwapOrdersInfo struct {
	BizWarmTips
	OrderInfo []BaseOrderInfo `json:"order_info"`
}

//easyjson:json
type BaseFillInfo struct {
	InstrumentId string `json:"instrument_id"`
	OrderQty     string `json:"order_qty"`
	TradeId      string `json:"trade_id"`
	Fee          string `json:"fee"`
	OrderId      string `json:"order_id"`
	Timestamp    string `json:"timestamp"`
	Price        string `json:"price"`
	Side         string `json:"side"`
	ExecType     string `json:"exec_type"`
}

//easyjson:json
type SwapFillsInfo []BaseFillInfo

//easyjson:json
type SwapAccountsSetting struct {
	BizWarmTips
	InstrumentId  string `json:"instrument_id"`
	LongLeverage  string `json:"long_leverage"`
	ShortLeverage string `json:"short_leverage"`
	MarginMode    string `json:"margin_mode"`
}

//easyjson:json
type BaseLedgerInfo struct {
	InstrumentId string `json:"instrument_id"`
	Fee          string `json:"fee"`
	Timestamp    string `json:"timestamp"`
	Amount       string `json:"amount"`
	LedgerId     string `json:"ledger_id"`
	Type         string `json:"type"`
}

//easyjson:json
type SwapAccountsLedgerList []BaseLedgerInfo

//easyjson:json
type BaseInstrumentInfo struct {
	InstrumentId    string `json:"instrument_id"`
	QuoteCurrency   string `json:"quote_currency"`
	TickSize        string `json:"tick_size"`
	ContractVal     string `json:"contract_val"`
	Listing         string `json:"listing"`
	UnderlyingIndex string `json:"underlying_index"`
	Delivery        string `json:"delivery"`
	Coin            string `json:"coin"`
	SizeIncrement   string `json:"size_increment"`
}

//easyjson:json
type SwapInstrumentList []BaseInstrumentInfo

type BaesDepthInfo []interface{}

//easyjson:json
type SwapInstrumentDepth struct {
	BizWarmTips
	Timestamp string          `json:"timestamp"`
	Time      string          `json:"time"`
	Bids      []BaesDepthInfo `json:"bids"`
	Asks      []BaesDepthInfo `json:"asks"`
}

//easyjson:json
type BaseTickerInfo struct {
	InstrumentId string `json:"instrument_id"`
	Last         string `json:"last"`
	Timestamp    string `json:"timestamp"`
	High24h      string `json:"high_24h"`
	Volume24h    string `json:"volume_24h"`
	Low24h       string `json:"low_24h"`
}

//easyjson:json
type SwapTickerList []BaseTickerInfo

//easyjson:json
type BaseTradeInfo struct {
	Timestamp string `json:"timestamp"`
	TradeId   string `json:"trade_id"`
	Side      string `json:"side"`
	Price     string `json:"price"`
	Size      string `json:"size"`
}

//easyjson:json
type SwapTradeList []BaseTradeInfo

type BaseCandleInfo []interface{}
type SwapCandleList []BaseCandleInfo

//easyjson:json
type SwapIndexInfo struct {
	BizWarmTips
	InstrumentId string `json:"instrument_id"`
	Index        string `json:"index"`
	Timestamp    string `json:"timestamp"`
}

//easyjson:json
type SwapRate struct {
	InstrumentId string `json:"instrument_id"`
	Timestamp    string `json:"timestamp"`
	Rate         string `json:"rate"`
}

//easyjson:json
type BaseInstrumentAmount struct {
	BizWarmTips
	InstrumentId string `json:"instrument_id"`
	Timestamp    string `json:"timestamp"`
	Amount       string `json:"amount"`
}

type SwapOpenInterest BaseInstrumentAmount

//easyjson:json
type SwapPriceLimit struct {
	BizWarmTips
	InstrumentId string `json:"instrument_id"`
	Lowest       string `json:"lowest"`
	Highest      string `json:"highest"`
	Timestamp    string `json:"timestamp"`
}

//easyjson:json
type BaseLiquidationInfo struct {
	InstrumentId string `json:"instrument_id"`
	Loss         string `json:"loss"`
	CreatedAt    string `json:"created_at"`
	Type         string `json:"type"`
	Price        string `json:"price"`
	Size         string `json:"size"`
}

//easyjson:json
type SwapLiquidationList []BaseLiquidationInfo

type SwapAccountHolds BaseInstrumentAmount

//easyjson:json
type SwapFundingTime struct {
	BizWarmTips
	InstrumentId string `json:"instrument_id"`
	FundingTime  string `json:"funding_time"`
}

//easyjson:json
type SwapMarkPrice struct {
	BizWarmTips
	InstrumentId string `json:"instrument_id"`
	MarkPrice    string `json:"mark_price"`
	Timestamp    string `json:"timestamp"`
}

//easyjson:json
type BaseHistoricalFundingRate struct {
	InstrumentId string `json:"instrument_id"`
	InterestRate string `json:"interest_rate"`
	FundingRate  string `json:"funding_rate"`
	FundingTime  string `json:"funding_time"`
	RealizedRate string `json:"realized_rate"`
}

//easyjson:json
type SwapHistoricalFundingRateList []BaseHistoricalFundingRate
