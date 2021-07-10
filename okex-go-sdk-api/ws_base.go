package okex

/*
 OKEX websocket api wrapper
 @author Lingting Fu
 @date 2018-12-27
 @version 1.0.0
*/

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"hash/crc32"
	"strings"

	"github.com/shopspring/decimal"
)

//easyjson:json
type BaseOp struct {
	Op string `json:"op"`
}

//easyjson:json
type BaseOpLoginArgs struct {
	ApiKey     string `json:"apiKey"`
	Passphrase string `json:"passphrase"`
	Timestamp  string `json:"timestamp"`
	Sign       string `json:"sign"`
}

//easyjson:json
type BaseOpLogin struct {
	BaseOp
	Args []BaseOpLoginArgs `json:"args"`
}

//easyjson:json
type BaseOpSubscriptionArgs struct {
	Channel  string `json:"channel"`
	InstId   string `json:"instId,omitempty"`
	Ccy      string `json:"ccy,omitempty"`
	InstType string `json:"instType,omitempty"`
}

//easyjson:json
type BaseOpSubscription struct {
	BaseOp
	Args []BaseOpSubscriptionArgs `json:"args"`
}

func subscribeOp(sts []*SubscriptionTopic) (op *BaseOpSubscription, err error) {
	args := make([]BaseOpSubscriptionArgs, 0, len(sts))
	for _, st := range sts {
		arg := BaseOpSubscriptionArgs{
			Channel: st.channel,
		}

		switch st.channel {
		case string(AccountChanel):
			arg.Ccy = st.filter
		case string(OrderChanel):
			arg.InstType = st.filter
		default:
			arg.InstId = st.filter
		}

		args = append(args, arg)
	}

	b := BaseOpSubscription{
		BaseOp: BaseOp{
			Op: CHNL_EVENT_SUBSCRIBE,
		},
		Args: args,
	}
	return &b, nil
}

func unsubscribeOp(sts []*SubscriptionTopic) (op *BaseOpSubscription, err error) {
	args := make([]BaseOpSubscriptionArgs, 0, len(sts))
	for _, st := range sts {
		arg := BaseOpSubscriptionArgs{
			Channel: st.channel,
		}

		switch st.channel {
		case string(AccountChanel):
			arg.Ccy = st.filter
		case string(OrderChanel):
			arg.InstType = st.filter
		default:
			arg.InstId = st.filter
		}

		args = append(args, arg)
	}

	b := BaseOpSubscription{
		BaseOp: BaseOp{
			Op: CHNL_EVENT_UNSUBSCRIBE,
		},
		Args: args,
	}
	return &b, nil
}

func loginOp(apiKey string, passphrase string, timestamp string, sign string) (op *BaseOpLogin, err error) {
	b := BaseOpLogin{
		BaseOp: BaseOp{
			Op: "login",
		},
		Args: []BaseOpLoginArgs{{
			ApiKey:     apiKey,
			Passphrase: passphrase,
			Timestamp:  timestamp,
			Sign:       sign,
		}},
	}
	return &b, nil
}

//easyjson:json
type SubscriptionTopic struct {
	channel string
	filter  string `default:""`
}

//easyjson:json
type WSEventResponse struct {
	Event   string `json:"event"`
	Success string `json:success`
	Channel string `json:"channel"`
}

func (r *WSEventResponse) Valid() bool {
	return (len(r.Event) > 0 && len(r.Channel) > 0) || r.Event == "login"
}

//easyjson:json
type WSTableResponse struct {
	Arg    *BaseOpSubscriptionArgs `json:"arg"`
	Action string                  `json:"action",default:""`
	Data   json.RawMessage         `json:"data"`
}

func (r *WSTableResponse) Valid() bool {
	return (r.Arg != nil || len(r.Action) > 0) && len(r.Data) > 0
}

//easyjson:json
type WSDepthItem struct {
	Asks      [][3]decimal.Decimal `json:"asks"`
	Bids      [][3]decimal.Decimal `json:"bids"`
	Timestamp string               `json:"ts"`
	Checksum  int32                `json:"checksum"`
}

func mergeDepths(oldDepths [][3]decimal.Decimal, newDepths [][3]decimal.Decimal) (*[][3]decimal.Decimal, error) {

	mergedDepths := [][3]decimal.Decimal{}
	oldIdx, newIdx := 0, 0

	for oldIdx < len(oldDepths) && newIdx < len(newDepths) {

		oldItem := oldDepths[oldIdx]
		newItem := newDepths[newIdx]

		oldPrice := oldItem[0]
		newPrice := newItem[0]

		if oldPrice.Equal(newPrice) {
			newNum := newItem[1]

			if newNum.IsPositive() {
				mergedDepths = append(mergedDepths, newItem)
			}

			oldIdx++
			newIdx++
		} else if oldPrice.GreaterThan(newPrice) {
			mergedDepths = append(mergedDepths, newItem)
			newIdx++
		} else if oldPrice.LessThan(newPrice) {
			mergedDepths = append(mergedDepths, oldItem)
			oldIdx++
		}
	}

	for ; oldIdx < len(oldDepths); oldIdx++ {
		mergedDepths = append(mergedDepths, oldDepths[oldIdx])
	}

	for ; newIdx < len(newDepths); newIdx++ {
		mergedDepths = append(mergedDepths, newDepths[newIdx])
	}

	return &mergedDepths, nil
}

func (di *WSDepthItem) update(newDI *WSDepthItem) error {
	newAskDepths, err1 := mergeDepths(di.Asks, newDI.Asks)
	if err1 != nil {
		return err1
	}

	newBidDepths, err2 := mergeDepths(di.Bids, newDI.Bids)
	if err2 != nil {
		return err2
	}

	crc32BaseBuffer, expectCrc32 := calCrc32(newAskDepths, newBidDepths)

	if expectCrc32 != newDI.Checksum && false { // TODO fix!
		return fmt.Errorf("Checksum's not correct. LocalString: %s, LocalCrc32: %d, RemoteCrc32: %d",
			crc32BaseBuffer.String(), expectCrc32, newDI.Checksum)
	} else {
		di.Checksum = newDI.Checksum
		di.Bids = *newBidDepths
		di.Asks = *newAskDepths
		di.Timestamp = newDI.Timestamp
	}

	return nil
}

func calCrc32(askDepths *[][3]decimal.Decimal, bidDepths *[][3]decimal.Decimal) (bytes.Buffer, int32) {
	crc32BaseBuffer := bytes.Buffer{}
	crcAskDepth, crcBidDepth := 25, 25
	if len(*askDepths) < 25 {
		crcAskDepth = len(*askDepths)
	}
	if len(*bidDepths) < 25 {
		crcBidDepth = len(*bidDepths)
	}
	if crcAskDepth == crcBidDepth {
		for i := 0; i < crcAskDepth; i++ {
			if crc32BaseBuffer.Len() > 0 {
				crc32BaseBuffer.WriteString(":")
			}
			crc32BaseBuffer.WriteString(
				fmt.Sprintf("%v:%v:%v:%v",
					(*bidDepths)[i][0], (*bidDepths)[i][1],
					(*askDepths)[i][0], (*askDepths)[i][1]))
		}
	} else {
		for i := 0; i < crcBidDepth; i++ {
			if crc32BaseBuffer.Len() > 0 {
				crc32BaseBuffer.WriteString(":")
			}
			crc32BaseBuffer.WriteString(
				fmt.Sprintf("%v:%v", (*bidDepths)[i][0], (*bidDepths)[i][1]))
		}

		for i := 0; i < crcAskDepth; i++ {
			if crc32BaseBuffer.Len() > 0 {
				crc32BaseBuffer.WriteString(":")
			}
			crc32BaseBuffer.WriteString(
				fmt.Sprintf("%v:%v", (*askDepths)[i][0], (*askDepths)[i][1]))
		}
	}
	expectCrc32 := int32(crc32.ChecksumIEEE(crc32BaseBuffer.Bytes()))
	return crc32BaseBuffer, expectCrc32
}

//easyjson:json
type WSDepthTableResponse struct {
	Arg    *BaseOpSubscriptionArgs `json:"arg"`
	Action string                  `json:"action",default:""`
	Data   WSDepthItemList         `json:"data"`
}

//easyjson:json
type WSDepthItemList []*WSDepthItem

func (r *WSDepthTableResponse) Valid() bool {
	return (r.Arg != nil || len(r.Action) > 0) && strings.Contains(r.Arg.Channel, string(BooksChanel)) && len(r.Data) > 0
}

//easyjson:json
type WSHotDepths struct {
	Table    string
	DepthMap map[string]*WSDepthItem
}

func NewWSHotDepths(tb string) *WSHotDepths {
	hd := WSHotDepths{}
	hd.Table = tb
	hd.DepthMap = map[string]*WSDepthItem{}
	return &hd
}

func (d *WSHotDepths) loadWSDepthTableResponse(r *WSDepthTableResponse) error {
	if d.Table != r.Arg.Channel {
		return fmt.Errorf("Loading WSDepthTableResponse failed becoz of "+
			"WSTableResponse(%s) not matched with WSHotDepths(%s)", r.Arg.Channel, d.Table)
	}

	if !r.Valid() {
		return errors.New("WSDepthTableResponse's format error.")
	}

	switch r.Action {
	case "partial":
		d.Table = r.Arg.Channel
		for i := 0; i < len(r.Data); i++ {
			crc32BaseBuffer, expectCrc32 := calCrc32(&r.Data[i].Asks, &r.Data[i].Bids)
			if expectCrc32 == r.Data[i].Checksum || true { // TODO fix!
				d.DepthMap[r.Arg.InstId] = r.Data[i]
			} else {
				return fmt.Errorf("Checksum's not correct. LocalString: %s, LocalCrc32: %d, RemoteCrc32: %d",
					crc32BaseBuffer.String(), expectCrc32, r.Data[i].Checksum)
			}
		}

	case "update":
		for i := 0; i < len(r.Data); i++ {
			newDI := r.Data[i]
			oldDI := d.DepthMap[r.Arg.InstId]
			if oldDI != nil {
				if err := oldDI.update(newDI); err != nil {
					return err
				}
			} else {
				d.DepthMap[r.Arg.InstId] = newDI
			}
		}

	default:
		break
	}

	return nil
}

//easyjson:json
type WSErrorResponse struct {
	Event     string `json:"event"`
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

func (r *WSErrorResponse) Valid() bool {
	return len(r.Event) > 0 && len(r.Message) > 0 && r.ErrorCode >= 30000
}

func loadResponse(rspMsg []byte) (interface{}, error) {
	var err error
	//log.Printf("%s", rspMsg)

	if string(rspMsg) == "pong" {
		return string(rspMsg), nil
	}

	tr := WSTableResponse{}
	err = JsonBytes2Struct(rspMsg, &tr)
	if err == nil && tr.Valid() {
		if strings.Contains(tr.Arg.Channel, string(BooksChanel)) {
			dtr := WSDepthTableResponse{
				Arg:    tr.Arg,
				Action: tr.Action,
			}

			err = JsonBytes2Struct(tr.Data, &dtr.Data)
			if err == nil && dtr.Valid() {
				return &dtr, nil
			}
		} else if tr.Arg.Channel == string(AccountChanel) {
			atr := UserSpotAccountWS{
				Arg: tr.Arg,
			}

			err = JsonBytes2Struct(tr.Data, &atr.Data)
			if err == nil {
				return &atr, nil
			}
		} else if tr.Arg.Channel == string(OrderChanel) {
			atr := UserOrdersWS{
				Arg: tr.Arg,
			}

			err = JsonBytes2Struct(tr.Data, &atr.Data)
			if err == nil {
				for _, o := range atr.Data {
					err = calcOrderFields(o.Order)
					if err != nil {
						panic(err)
					}
				}
				return &atr, nil
			}
		}

		return &tr, nil
	}

	evtR := WSEventResponse{}
	err = JsonBytes2Struct(rspMsg, &evtR)
	if err == nil && evtR.Valid() {
		return &evtR, nil
	}

	er := WSErrorResponse{}
	err = JsonBytes2Struct(rspMsg, &er)
	if err == nil && er.Valid() {
		return &er, nil
	}

	return nil, err

}

type ReceivedDataCallback func(interface{}) error

func defaultPrintData(obj interface{}) error {
	switch obj.(type) {
	case string:
		fmt.Println("okex.defaultPrintData:string", obj)
	default:
		msg, err := Struct2JsonString(obj)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		if msg != "null" {
			fmt.Println("okex.defaultPrintData:default", msg)
		}
	}
	return nil
}
