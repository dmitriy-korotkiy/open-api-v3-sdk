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

type BaseOp struct {
	Op   string   `json:"op"`
	Args []string `json:"args"`
}

func subscribeOp(sts []*SubscriptionTopic) (op *BaseOp, err error) {

	strArgs := []string{}

	for i := 0; i < len(sts); i++ {
		channel, err := sts[i].ToString()
		if err != nil {
			return nil, err
		}
		strArgs = append(strArgs, channel)
	}

	b := BaseOp{
		Op:   "subscribe",
		Args: strArgs,
	}
	return &b, nil
}

func unsubscribeOp(sts []*SubscriptionTopic) (op *BaseOp, err error) {

	strArgs := []string{}

	for i := 0; i < len(sts); i++ {
		channel, err := sts[i].ToString()
		if err != nil {
			return nil, err
		}
		strArgs = append(strArgs, channel)
	}

	b := BaseOp{
		Op:   CHNL_EVENT_UNSUBSCRIBE,
		Args: strArgs,
	}
	return &b, nil
}

func loginOp(apiKey string, passphrase string, timestamp string, sign string) (op *BaseOp, err error) {
	b := BaseOp{
		Op:   "login",
		Args: []string{apiKey, passphrase, timestamp, sign},
	}
	return &b, nil
}

type SubscriptionTopic struct {
	channel string
	filter  string `default:""`
}

func (st *SubscriptionTopic) ToString() (topic string, err error) {
	if len(st.channel) == 0 {
		return "", ERR_WS_SUBSCRIOTION_PARAMS
	}

	if len(st.filter) > 0 {
		return st.channel + ":" + st.filter, nil
	} else {
		return st.channel, nil
	}
}

type WSEventResponse struct {
	Event   string `json:"event"`
	Success string `json:success`
	Channel string `json:"channel"`
}

func (r *WSEventResponse) Valid() bool {
	return (len(r.Event) > 0 && len(r.Channel) > 0) || r.Event == "login"
}

type WSTableResponse struct {
	Table  string          `json:"table"`
	Action string          `json:"action",default:""`
	Data   json.RawMessage `json:"data"`
}

func (r *WSTableResponse) Valid() bool {
	return (len(r.Table) > 0 || len(r.Action) > 0) && len(r.Data) > 0
}

type WSDepthItem struct {
	InstrumentId string               `json:"instrument_id"`
	Asks         [][3]decimal.Decimal `json:"asks"`
	Bids         [][3]decimal.Decimal `json:"bids"`
	Timestamp    string               `json:"timestamp"`
	Checksum     int32                `json:"checksum"`
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

type WSDepthTableResponse struct {
	Table  string         `json:"table"`
	Action string         `json:"action",default:""`
	Data   []*WSDepthItem `json:"data"`
}

func (r *WSDepthTableResponse) Valid() bool {
	return (len(r.Table) > 0 || len(r.Action) > 0) && strings.Contains(r.Table, "depth") && len(r.Data) > 0
}

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
	if d.Table != r.Table {
		return fmt.Errorf("Loading WSDepthTableResponse failed becoz of "+
			"WSTableResponse(%s) not matched with WSHotDepths(%s)", r.Table, d.Table)
	}

	if !r.Valid() {
		return errors.New("WSDepthTableResponse's format error.")
	}

	switch r.Action {
	case "partial":
		d.Table = r.Table
		for i := 0; i < len(r.Data); i++ {
			crc32BaseBuffer, expectCrc32 := calCrc32(&r.Data[i].Asks, &r.Data[i].Bids)
			if expectCrc32 == r.Data[i].Checksum || true { // TODO fix!
				d.DepthMap[r.Data[i].InstrumentId] = r.Data[i]
			} else {
				return fmt.Errorf("Checksum's not correct. LocalString: %s, LocalCrc32: %d, RemoteCrc32: %d",
					crc32BaseBuffer.String(), expectCrc32, r.Data[i].Checksum)
			}
		}

	case "update":
		for i := 0; i < len(r.Data); i++ {
			newDI := r.Data[i]
			oldDI := d.DepthMap[newDI.InstrumentId]
			if oldDI != nil {
				if err := oldDI.update(newDI); err != nil {
					return err
				}
			} else {
				d.DepthMap[newDI.InstrumentId] = newDI
			}
		}

	default:
		break
	}

	return nil
}

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
		if strings.Contains(tr.Table, "depth") {
			dtr := WSDepthTableResponse{
				Table:  tr.Table,
				Action: tr.Action,
			}

			err = JsonBytes2Struct(tr.Data, &dtr.Data)
			if err == nil && dtr.Valid() {
				return &dtr, nil
			}
		} else if tr.Table == "spot/account" {
			atr := UserSpotAccountWS{
				Table:  WSEventTable(tr.Table),
			}

			err = JsonBytes2Struct(tr.Data, &atr.Data)
			if err == nil {
				return &atr, nil
			}
		} else if tr.Table == "spot/order" {
			atr := UserOrdersWS{
				Table:  WSEventTable(tr.Table),
			}

			err = JsonBytes2Struct(tr.Data, &atr.Data)
			if err == nil {
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
