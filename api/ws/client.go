package ws

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/yitech/okex"
	"github.com/yitech/okex/events"
)

// ClientWs is the websocket api client
//
// https://www.okx.com/docs-v5/en/#websocket-api
type ClientWs struct {
	Cancel              context.CancelFunc
	DoneChan            chan interface{}
	StructuredEventChan chan interface{}
	RawEventChan        chan *events.Basic
	ErrChan             chan *events.Error
	SubscribeChan       chan *events.Subscribe
	UnsubscribeCh       chan *events.Unsubscribe
	LoginChan           chan *events.Login
	SuccessChan         chan *events.Success
	sendChan            map[bool]chan []byte
	url                 map[bool]okex.BaseURL
	conn                map[bool]*websocket.Conn
	dialer              *websocket.Dialer
	apiKey              string
	secretKey           []byte
	passphrase          string
	lastTransmit        map[bool]*time.Time
	mu                  map[bool]*sync.RWMutex
	AuthRequested       *time.Time
	Authorized          bool
	Private             *Private
	Public              *Public
	Trade               *Trade
	ctx                 context.Context
}

const (
	redialTick = 2 * time.Second
	writeWait  = 3 * time.Second
	pongWait   = 30 * time.Second
	PingPeriod = (pongWait * 8) / 10
)

// NewClient returns a pointer to a fresh ClientWs
func NewClient(ctx context.Context, apiKey, secretKey, passphrase string, url map[bool]okex.BaseURL) *ClientWs {
	ctx, cancel := context.WithCancel(ctx)
	c := &ClientWs{
		apiKey:       apiKey,
		secretKey:    []byte(secretKey),
		passphrase:   passphrase,
		ctx:          ctx,
		Cancel:       cancel,
		url:          url,
		sendChan:     map[bool]chan []byte{true: make(chan []byte, 3), false: make(chan []byte, 3)},
		DoneChan:     make(chan interface{}),
		conn:         make(map[bool]*websocket.Conn),
		dialer:       websocket.DefaultDialer,
		lastTransmit: make(map[bool]*time.Time),
		mu:           map[bool]*sync.RWMutex{true: {}, false: {}},
	}
	c.Private = NewPrivate(c)
	c.Public = NewPublic(c)
	c.Trade = NewTrade(c)
	return c
}

// Connect into the server
//
// https://www.okx.com/docs-v5/en/#websocket-api-connect
func (c *ClientWs) Connect(p bool) error {
	if c.conn[p] != nil {
		return nil
	}
	err := c.dial(p)
	if err == nil {
		return nil
	}
	ticker := time.NewTicker(redialTick)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			err = c.dial(p)
			if err == nil {
				return nil
			}
		case <-c.ctx.Done():
			return c.handleCancel("connect")
		}
	}
}

// Login
//
// https://www.okx.com/docs-v5/en/#websocket-api-login
func (c *ClientWs) Login() error {
	if c.Authorized {
		return nil
	}
	if c.AuthRequested != nil && time.Since(*c.AuthRequested).Seconds() < 30 {
		return nil
	}
	now := time.Now()
	c.AuthRequested = &now
	method := http.MethodGet
	path := "/users/self/verify"
	ts, sign := c.sign(method, path)
	args := []map[string]string{
		{
			"apiKey":     c.apiKey,
			"passphrase": c.passphrase,
			"timestamp":  ts,
			"sign":       sign,
		},
	}
	return c.Send(true, okex.LoginOperation, args)
}

// Subscribe
// Users can choose to subscribe to one or more channels, and the total length of multiple channels cannot exceed 4096 bytes.
//
// https://www.okx.com/docs-v5/en/#websocket-api-subscribe
func (c *ClientWs) Subscribe(p bool, ch []okex.ChannelName, args ...map[string]string) error {
	chCount := max(len(ch), 1)
	tmpArgs := make([]map[string]string, chCount*len(args))

	n := 0
	for i := 0; i < chCount; i++ {
		for _, arg := range args {
			tmpArgs[n] = make(map[string]string)
			for k, v := range arg {
				tmpArgs[n][k] = v
			}
			if len(ch) > 0 {
				tmpArgs[n]["channel"] = string(ch[i])
			}
			n++
		}
	}

	return c.Send(p, okex.SubscribeOperation, tmpArgs)
}

// Unsubscribe into channel(s)
//
// https://www.okx.com/docs-v5/en/#websocket-api-unsubscribe
func (c *ClientWs) Unsubscribe(p bool, ch []okex.ChannelName, args map[string]string) error {
	tmpArgs := make([]map[string]string, len(ch))
	for i, name := range ch {
		tmpArgs[i] = make(map[string]string)
		tmpArgs[i]["channel"] = string(name)
		for k, v := range args {
			tmpArgs[i][k] = v
		}
	}
	return c.Send(p, okex.UnsubscribeOperation, tmpArgs)
}

// Send message through either connections
func (c *ClientWs) Send(p bool, op okex.Operation, args []map[string]string, extras ...map[string]string) error {
	if op != okex.LoginOperation {
		err := c.Connect(p)
		if err == nil {
			if p {
				err = c.WaitForAuthorization()
				if err != nil {
					return err
				}
			}
		} else {
			return err
		}
	}

	data := map[string]interface{}{
		"op":   op,
		"args": args,
	}
	for _, extra := range extras {
		for k, v := range extra {
			data[k] = v
		}
	}
	j, err := json.Marshal(data)
	if err != nil {
		return err
	}
	c.sendChan[p] <- j
	return nil
}

// SetChannels to receive certain events on separate channel
func (c *ClientWs) SetChannels(errCh chan *events.Error, subCh chan *events.Subscribe, unSub chan *events.Unsubscribe, lCh chan *events.Login, sCh chan *events.Success) {
	c.ErrChan = errCh
	c.SubscribeChan = subCh
	c.UnsubscribeCh = unSub
	c.LoginChan = lCh
	c.SuccessChan = sCh
}

// SetDialer sets a custom dialer for the WebSocket connection.
func (c *ClientWs) SetDialer(dialer *websocket.Dialer) {
	c.dialer = dialer
}

func (c *ClientWs) SetEventChannels(structuredEventCh chan interface{}, rawEventCh chan *events.Basic) {
	c.StructuredEventChan = structuredEventCh
	c.RawEventChan = rawEventCh
}

// WaitForAuthorization waits for the auth response and try to log in if it was needed
func (c *ClientWs) WaitForAuthorization() error {
	if c.Authorized {
		return nil
	}
	if err := c.Login(); err != nil {
		return err
	}
	ticker := time.NewTicker(time.Millisecond * 300)
	defer ticker.Stop()
	for range ticker.C {
		if c.Authorized {
			return nil
		}
	}
	return nil
}

func (c *ClientWs) dial(p bool) error {
	c.mu[p].Lock()
	conn, res, err := c.dialer.Dial(string(c.url[p]), nil)
	if err != nil {
		var statusCode int
		if res != nil {
			statusCode = res.StatusCode
		}
		return fmt.Errorf("error %d: %w", statusCode, err)
	}
	c.conn[p] = conn
	c.mu[p].Unlock()

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("error closing body: %v\n", err)
		}
	}(res.Body)
	go func() {
		err := c.receiver(p)
		if err != nil {
			fmt.Printf("receiver error: %v\n", err)
		}
	}()
	go func() {
		err := c.sender(p)
		if err != nil {
			fmt.Printf("sender error: %v\n", err)
		}
	}()

	return nil
}

func (c *ClientWs) sender(p bool) error {
	ticker := time.NewTicker(time.Millisecond * 300)
	defer ticker.Stop()
	for {
		select {
		case data := <-c.sendChan[p]:
			c.mu[p].RLock()
			err := c.conn[p].SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				c.mu[p].RUnlock()
				return err
			}
			w, err := c.conn[p].NextWriter(websocket.TextMessage)
			if err != nil {
				c.mu[p].RUnlock()
				return err
			}
			if _, err = w.Write(data); err != nil {
				c.mu[p].RUnlock()
				return err
			}
			now := time.Now()
			c.lastTransmit[p] = &now
			c.mu[p].RUnlock()
			if err := w.Close(); err != nil {
				return err
			}
		case <-ticker.C:
			c.mu[p].RLock()
			conn := c.conn[p]
			lastTransmit := c.lastTransmit[p]
			c.mu[p].RUnlock()
			if conn != nil && (lastTransmit == nil || (lastTransmit != nil && time.Since(*lastTransmit) > PingPeriod)) {
				go func() {
					c.sendChan[p] <- []byte("ping")
				}()
			}
		case <-c.ctx.Done():
			return c.handleCancel("sender")
		}
	}
}

func (c *ClientWs) receiver(p bool) error {
	for {
		select {
		case <-c.ctx.Done():
			return c.handleCancel("receiver")
		default:
			c.mu[p].RLock()
			err := c.conn[p].SetReadDeadline(time.Now().Add(pongWait))
			if err != nil {
				c.mu[p].RUnlock()
				return err
			}
			mt, data, err := c.conn[p].ReadMessage()
			if err != nil {
				c.mu[p].RUnlock()
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					return c.conn[p].Close()
				}
				return err
			}
			c.mu[p].RUnlock()
			now := time.Now()
			c.mu[p].Lock()
			c.lastTransmit[p] = &now
			c.mu[p].Unlock()
			if mt == websocket.TextMessage && string(data) != "pong" {
				e := &events.Basic{}
				if err := json.Unmarshal(data, &e); err != nil {
					return err
				}
				go func() {
					c.process(data, e)
				}()
			}
		}
	}
}

func (c *ClientWs) sign(method, path string) (string, string) {
	t := time.Now().UTC().Unix()
	ts := fmt.Sprint(t)
	s := ts + method + path
	p := []byte(s)
	h := hmac.New(sha256.New, c.secretKey)
	h.Write(p)
	return ts, base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (c *ClientWs) handleCancel(msg string) error {
	go func() {
		c.DoneChan <- msg
	}()
	return fmt.Errorf("operation cancelled: %s", msg)
}

func (c *ClientWs) process(data []byte, e *events.Basic) bool {
	switch e.Event {
	case "error":
		e := events.Error{}
		_ = json.Unmarshal(data, &e)
		if c.ErrChan != nil {
			c.ErrChan <- &e
		}
		return true
	case "subscribe":
		e := events.Subscribe{}
		_ = json.Unmarshal(data, &e)
		if c.SubscribeChan != nil {
			c.SubscribeChan <- &e
		}
		if c.StructuredEventChan != nil {
			c.StructuredEventChan <- e
		}
		return true
	case "unsubscribe":
		e := events.Unsubscribe{}
		_ = json.Unmarshal(data, &e)
		if c.UnsubscribeCh != nil {
			c.UnsubscribeCh <- &e
		}
		if c.StructuredEventChan != nil {
			c.StructuredEventChan <- e
		}
		return true
	case "login":
		if time.Since(*c.AuthRequested).Seconds() > 30 {
			c.AuthRequested = nil
			_ = c.Login()
			break
		}
		c.Authorized = true
		e := events.Login{}
		_ = json.Unmarshal(data, &e)
		if c.LoginChan != nil {
			c.LoginChan <- &e
		}
		if c.StructuredEventChan != nil {
			c.StructuredEventChan <- e
		}
		return true
	}
	if c.Private.Process(data, e) {
		return true
	}
	if c.Public.Process(data, e) {
		return true
	}
	if e.ID != "" {
		if e.Code != 0 {
			ee := *e
			ee.Event = "error"
			return c.process(data, &ee)
		}
		e := events.Success{}
		_ = json.Unmarshal(data, &e)
		if c.SuccessChan != nil {
			c.SuccessChan <- &e
		}
		if c.StructuredEventChan != nil {
			c.StructuredEventChan <- e
		}
		return true
	}
	c.RawEventChan <- e
	return false
}
