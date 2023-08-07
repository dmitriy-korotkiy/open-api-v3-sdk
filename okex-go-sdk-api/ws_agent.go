package okex

/*
 OKEX websocket API agent
 @author Lingting Fu
 @date 2018-12-27
 @version 1.0.0
*/

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/darkfoxs96/golimiter/limiter"
	"github.com/gorilla/websocket"
)

type ChanelType string

const (
	AccountChanel ChanelType = "account"
	OrderChanel   ChanelType = "orders"
	BooksChanel   ChanelType = "books"
	Books50Chanel ChanelType = "books50-l2-tbt"
	Books5Chanel  ChanelType = "books5"
)

type OKWSAgent struct {
	baseUrl string
	config  *Config
	conn    *websocket.Conn

	wsEvtCh  chan interface{}
	wsErrCh  chan interface{}
	wsTbCh   chan interface{}
	stopCh   chan interface{}
	errCh    chan error
	signalCh chan os.Signal

	subMap         map[string][]ReceivedDataCallback
	activeChannels map[string]bool
	hotDepthsMap   map[string]*WSHotDepths

	processMut sync.RWMutex
	wsMut      sync.Mutex
}

func (a *OKWSAgent) Start(config *Config) error {
	a.baseUrl = config.WSEndpoint + ""
	if config.IsPrint {
		log.Printf("Connecting to %s", a.baseUrl)
	}

	c, _, err := websocket.DefaultDialer.Dial(a.baseUrl, nil)
	oldCh := a.subMap

	a.config = config
	if err != nil {
		if config.IsPrint {
			log.Fatalf("dial:%+v", err)
		}
		return err
	}

	if a.config.IsPrint {
		log.Printf("Connected to %s", a.baseUrl)
	}
	a.conn = c
	a.config = config

	a.wsEvtCh = make(chan interface{})
	a.wsErrCh = make(chan interface{})
	a.wsTbCh = make(chan interface{})
	a.errCh = make(chan error)
	a.stopCh = make(chan interface{}, 16)
	a.signalCh = make(chan os.Signal)
	a.activeChannels = make(map[string]bool)
	a.subMap = make(map[string][]ReceivedDataCallback)
	a.hotDepthsMap = make(map[string]*WSHotDepths)

	signal.Notify(a.signalCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go a.work()
	go a.receive()
	go a.finalize()

	for chName, fns := range oldCh {
		var channel string
		var filter string

		chFil := strings.Split(chName, ":")
		if len(chFil) > 0 {
			channel = chFil[0]
		}
		if len(chFil) > 1 {
			filter = chFil[1]
		}

		for _, fn := range fns {
			var filters []string
			if filter == "" {
				filters = append(filters, "")
			} else {
				filters = strings.Split(filter, ";")
			}

			if err = a.Subscribe(ChanelType(channel), filters, fn); err != nil {
				return err
			}
		}
	}

	return nil
}

var wsLimiter = limiter.Limiter{Limit: 40, PeriodMillisecond: 1_000}

func (a *OKWSAgent) Subscribe(channel ChanelType, filters []string, cb ReceivedDataCallback) error {
	a.processMut.Lock()
	defer a.processMut.Unlock()

	sts := make([]*SubscriptionTopic, 0, len(filters))
	for _, filter := range filters {
		sts = append(sts, &SubscriptionTopic{string(channel), filter})
	}
	bo, err := subscribeOp(sts)
	if err != nil {
		return err
	}

	msg, err := Struct2JsonString(bo)
	if a.config.IsPrint {
		log.Printf("Send Msg: %s", msg)
	}

	ch := make(chan interface{}, 2)

	wsLimiter.Wait(func() {
		a.wsMut.Lock()
		defer a.wsMut.Unlock()

		err = a.conn.WriteMessage(websocket.TextMessage, []byte(msg))
		ch <- nil
	})

	<-ch
	if err != nil {
		return err
	}

	chName := getFullTopic(string(channel), filters)
	cbs := a.subMap[chName]
	if cbs == nil {
		cbs = []ReceivedDataCallback{}
		a.activeChannels[chName] = false
	}

	if cb != nil {
		cbs = append(cbs, cb)
		a.subMap[chName] = cbs
	}

	return nil
}

func (a *OKWSAgent) UnSubscribe(channel ChanelType, filters []string) error {
	a.processMut.Lock()
	defer a.processMut.Unlock()

	sts := make([]*SubscriptionTopic, 0, len(filters))
	for _, filter := range filters {
		sts = append(sts, &SubscriptionTopic{string(channel), filter})
	}
	bo, err := unsubscribeOp(sts)
	if err != nil {
		return err
	}

	msg, err := Struct2JsonString(bo)
	if a.config.IsPrint {
		log.Printf("Send Msg: %s", msg)
	}

	ch := make(chan interface{}, 2)

	wsLimiter.Wait(func() {
		a.wsMut.Lock()
		defer a.wsMut.Unlock()

		err = a.conn.WriteMessage(websocket.TextMessage, []byte(msg))
		ch <- nil
	})

	<-ch
	if err != nil {
		return err
	}

	chName := getFullTopic(string(channel), filters)
	a.subMap[chName] = nil
	a.activeChannels[chName] = false

	return nil
}

func getFullTopic(channel string, filters []string) string {
	if len(filters) == 0 {
		return channel
	}

	return channel + ":" + strings.Join(filters, ";")
}

func (a *OKWSAgent) Login(apiKey, passphrase string) error {
	a.wsMut.Lock()
	defer a.wsMut.Unlock()

	timestamp := EpochTime()

	preHash := PreHashString(timestamp, GET, "/users/self/verify", "")
	if sign, err := HmacSha256Base64Signer(preHash, a.config.SecretKey); err != nil {
		return err
	} else {
		op, err := loginOp(apiKey, passphrase, timestamp, sign)
		data, err := Struct2JsonString(op)
		log.Printf("Send Msg: %s", data)
		err = a.conn.WriteMessage(websocket.TextMessage, []byte(data))
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 100)
	}
	return nil
}

func (a *OKWSAgent) keepalive() {
	a.ping()
}

func (a *OKWSAgent) Stop() error {
	defer func() {
		if r := recover(); r != nil && a.config.IsPrint {
			log.Printf("Stop End. Recover msg: %+v", a)
		}
	}()

	close(a.stopCh)
	return nil
}

func (a *OKWSAgent) finalize() error {
	defer func() {
		r := recover()

		if r != nil {
			log.Println("Finalize End. Connection to WebSocket is closed.\n", r)
		} else if a.config.IsPrint {
			log.Println("Finalize End. Connection to WebSocket is closed.", r)
		}
	}()

	select {
	case <-a.stopCh:
		if a.conn != nil {
			close(a.errCh)
			close(a.wsTbCh)
			close(a.wsEvtCh)
			close(a.wsErrCh)
			return a.conn.Close()
		}
	}

	return nil
}

func (a *OKWSAgent) ping() {
	a.wsMut.Lock()
	defer a.wsMut.Unlock()

	msg := "ping"
	//log.Printf("Send Msg: %s", msg)
	a.conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

func (a *OKWSAgent) GzipDecode(in []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(in))
	defer reader.Close()

	return ioutil.ReadAll(reader)
}

func (a *OKWSAgent) handleErrResponse(r interface{}) error {
	if r != nil {
		log.Printf("handleErrResponse %+v \n", r)
	}
	return nil
}

func (a *OKWSAgent) handleEventResponse(r interface{}) error {
	if r == nil {
		return nil
	}

	er := r.(*WSEventResponse)

	a.processMut.Lock()
	defer a.processMut.Unlock()
	a.activeChannels[er.Channel] = (er.Event == CHNL_EVENT_SUBSCRIBE)
	return nil
}

func (a *OKWSAgent) handleTableResponse(r interface{}) error {
	tb := ""
	switch r.(type) {
	case *WSTableResponse:
		tb = r.(*WSTableResponse).Arg.Channel
	case *UserSpotAccountWS:
		v := r.(*UserSpotAccountWS)
		if len(v.Data) > 0 {
			tb = string(v.Arg.Channel) + ":" + v.Arg.Ccy
		} else {
			return fmt.Errorf("handleTableResponse() !(len(UserSpotAccountWS.Data) > 0)")
		}
	case *UserOrdersWS:
		v := r.(*UserOrdersWS)
		if len(v.Data) > 0 {
			tb = v.Arg.Channel + ":" + v.Arg.InstType
		} else {
			return fmt.Errorf("handleTableResponse() !(len(UserOrdersWS.Data) > 0)")
		}
	case *WSDepthTableResponse:
		v := r.(*WSDepthTableResponse)
		if len(v.Data) > 0 {
			tb = v.Arg.Channel + ":" + v.Arg.InstId
		} else {
			return fmt.Errorf("handleTableResponse() !(len(WSDepthTableResponse.Data) > 0)")
		}

		a.processMut.RLock()
		defer a.processMut.RUnlock()

		for key, cbs := range a.subMap {
			if !strings.HasPrefix(key, v.Arg.Channel+":") || !(strings.HasSuffix(key, v.Arg.InstId) || strings.Contains(key, v.Arg.InstId+";")) {
				continue
			}

			for i := 0; i < len(cbs); i++ {
				cb := cbs[i]
				if err := cb(r); err != nil {
					return err
				}
			}
		}

		return nil
	}

	a.processMut.RLock()
	defer a.processMut.RUnlock()

	cbs := a.subMap[tb]
	if cbs != nil {
		for i := 0; i < len(cbs); i++ {
			cb := cbs[i]
			if err := cb(r); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *OKWSAgent) work() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Work End. Recover msg: %+v", r)
			debug.PrintStack()
		}
	}()

	ticker := time.NewTicker(25 * time.Second)
	defer func() {
		_ = a.Stop()
		ticker.Stop()
	}()

	for {
		select {
		case <-ticker.C:
			a.keepalive()
		case errR := <-a.wsErrCh:
			a.handleErrResponse(errR)
		case evtR := <-a.wsEvtCh:
			a.handleEventResponse(evtR)
		case tb := <-a.wsTbCh:
			a.handleTableResponse(tb)
		case <-a.signalCh:
			break
		case err := <-a.errCh:
			DefaultDataCallBack(err)
			break
		case <-a.stopCh:
			return

		}
	}
}

func (a *OKWSAgent) restart() {
	time.Sleep(time.Second * 2)
	err := a.Start(a.config)
	if err != nil {
		log.Println("receive().recover.restart():", err)
		go a.restart()
	}

}

func (a *OKWSAgent) receive() {
	defer func() {
		if r := recover(); r != nil {
			go a.restart()
		}
	}()

	for {
		messageType, message, err := a.conn.ReadMessage()
		if err != nil {
			a.errCh <- err
			break
		}

		txtMsg := message
		switch messageType {
		case websocket.TextMessage:
		case websocket.BinaryMessage:
			txtMsg, err = a.GzipDecode(message)
		}

		rsp, err := loadResponse(txtMsg)
		if rsp != nil {
			if a.config.IsPrint {
				log.Printf("LoadedRep: %+v, err: %+v", rsp, err)
			}
		} else {
			if a.config.IsPrint {
				log.Printf("TextMsg: %s", txtMsg)
			}
		}

		if err != nil {
			break
		}

		switch rsp.(type) {
		case *WSErrorResponse:
			a.wsErrCh <- rsp
		case *WSEventResponse:
			er := rsp.(*WSEventResponse)
			a.wsEvtCh <- er
		case *WSDepthTableResponse:
			//var err error
			dtr := rsp.(*WSDepthTableResponse)
			//hotDepths := a.hotDepthsMap[dtr.Table]
			//if hotDepths == nil {
			//	hotDepths = NewWSHotDepths(dtr.Table)
			//	err = hotDepths.loadWSDepthTableResponse(dtr)
			//	if err == nil {
			//		a.hotDepthsMap[dtr.Table] = hotDepths
			//	}
			//} else {
			//	err = hotDepths.loadWSDepthTableResponse(dtr)
			//}

			//if err == nil {
			a.wsTbCh <- dtr
			//} else {
			//	log.Printf("Failed to loadWSDepthTableResponse, dtr: %+v, err: %+v", dtr, err)
			//}

		case *WSTableResponse:
			tb := rsp.(*WSTableResponse)
			a.wsTbCh <- tb
		case *UserSpotAccountWS:
			tb := rsp.(*UserSpotAccountWS)
			a.wsTbCh <- tb
		case *UserOrdersWS:
			tb := rsp.(*UserOrdersWS)
			a.wsTbCh <- tb
		default:
			//log.Println(rsp)
		}
	}
}
