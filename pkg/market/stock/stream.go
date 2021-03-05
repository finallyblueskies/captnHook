package stock

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bareish/captnHook/pkg/services"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const (
	MaxConnectionAttempts = 5 			// max retries
	TradeUpdates   = "trade_updates" 	// alpaca trade updates string
	AccountUpdates = "account_updates"  // alpaca account updates string
)

// Errors
const(
	InvalidAlpacaStreamErr = "invalid Alpaca stream (%s)"
	MaxRetriesExceededErr = "could not open Alpaca Alpaca stream (max retries exceeded)"
	FailedAuthErr = "failed to authorize connection"
)

// Stream represents the websocket connection
type Stream struct {
	sync.Mutex
	sync.Once
	conn                  *websocket.Conn			// websocket connection
	authenticated, closed atomic.Value
	handlers              sync.Map
	quoteHandlers         sync.Map
	streamTradeHandler	     chan StreamTrade
	streamTradeUpdateHandler chan TradeUpdate
	streamAggHandler 	     chan StreamAgg
	streamAccountHandler     chan AccountActivity
	base                  string
	id   				  string
	secret				  string
}

// NewRestClient creates a new Stream
func NewStreamClient(cs services.ConfigService) *Stream {
	// get alpaca config variables
	alpacaConfig := cs.Get().Alpaca
	stream := &Stream{
		authenticated: atomic.Value{},
		handlers:      sync.Map{},
		quoteHandlers: sync.Map{},
		base:          alpacaConfig.WebSocketURL,
		id: 		   alpacaConfig.ClientID,
		secret: 	   alpacaConfig.ClientSecret,
	}

	stream.authenticated.Store(false)
	stream.closed.Store(false)
	return stream
}

// StreamQuote ...
func (s *Stream) StreamQuote(ticker string, handler chan StreamQuote) error {
	chanKey:= "Q."+ ticker
	s.quoteHandlers.Store(chanKey, handler)
	if err := s.sub(chanKey); err != nil {
		s.quoteHandlers.Delete(chanKey)
		log.Println("StreamQuote failed on ticker: " + ticker)
		log.Println(err)
		return err
	}

	return nil
}

// StreamTradeUpdates...
func (s *Stream) StreamTradeUpdates(ticker string, handler chan TradeUpdate) error {
	s.streamTradeUpdateHandler = handler
	err := s.sub(TradeUpdates)
	if err != nil {
		return err
	}

	return nil
}

// AccountUpdates ...
func (s *Stream) AccountUpdates(ticker string, handler chan AccountActivity) error {
	s.streamAccountHandler = handler
	err := s.sub(AccountUpdates)
	if err != nil {
		return err
	}

	return nil
}


// todo could be called by multiple threads consider adding mutex
// Subscribe to the specified Alpaca channel.
func (s *Stream) Subscribe(channel string, handler chan interface{}) (err error) {
	switch {
	case channel == TradeUpdates:
		fallthrough
	case channel == AccountUpdates:
		fallthrough
	case strings.HasPrefix(channel, "Q."):
		fallthrough
	case strings.HasPrefix(channel, "T."):
		fallthrough
	case strings.HasPrefix(channel, "AM."):
	default:
		err = fmt.Errorf(InvalidAlpacaStreamErr, channel)
		return
	}
	s.handlers.Store(channel, handler)
	if err = s.sub(channel); err != nil {
		s.handlers.Delete(channel)
		return
	}
	return
}

// Unsubscribe removes the specified Alpaca
func (s *Stream) Unsubscribe(channel string) (err error){
	// stop listening
	err = s.unsub(channel)
	if err != nil {
		return err
	}
	// remove
	s.handlers.Delete(channel)
	return nil
}

// Start
func (s *Stream) Start() (err error) {
	// todo: if markets are closed exit early
	// connect
	if s.conn == nil {
		s.conn, err = s.openSocket()
		if err != nil {
			log.Print(s.base, ": Web Socket Connection Error: ", err)
			return err
		}
	}
	// auth
	if err = s.auth(); err != nil {
		log.Print(s.base, ": Web Socket Auth Error:", err)
		return err
	}
	for {
		msg := ServerMsg{}

		if err := s.conn.ReadJSON(&msg); err == nil {
			msgBytes, _ := json.Marshal(msg.Data)
			if strings.HasPrefix(msg.Stream, "Q.") {
				var quote StreamQuote
				err = json.Unmarshal(msgBytes, &quote)
				channelHandler := s.findStockHandler(msg.Stream)
				channelHandler <- quote
			}else if strings.HasPrefix(msg.Stream, "T.") {
				var trade StreamTrade
				err = json.Unmarshal(msgBytes, &trade)
				if err != nil {
					return err
				}
				s.streamTradeHandler <- trade
			}else if strings.HasPrefix(msg.Stream, "AM.") {
				var agg StreamAgg
				err = json.Unmarshal(msgBytes, &agg)
				if err != nil {
					return err
				}
				s.streamAggHandler <- agg
			}else if msg.Stream == TradeUpdates {
				var tradeupdate TradeUpdate
				err = json.Unmarshal(msgBytes, &tradeupdate)
				if err != nil {
					return err
				}
				s.streamTradeUpdateHandler <- tradeupdate
			}else if msg.Stream == AccountUpdates {

				// Todo there isn't an AccountUpdate struct...

				log.Print(string(msgBytes))
			}
		} else {
			if websocket.IsCloseError(err) {
				// if this was a graceful closure, don't reconnect
				if s.closed.Load().(bool) {
					return err
				}
			} else {
				return fmt.Errorf("alpaca alpacastream read error (%v)", err)
			}
			err := s.reconnect()
			if err != nil {
				return err
			}
		}
	}
}

// Close gracefully closes the Alpaca alpacastream.
func (s *Stream) Close() error {
	if s.conn == nil {
		return nil
	}
	if err := s.conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
	); err != nil {
		return err
	}
	// so we know it was gracefully closed
	s.closed.Store(true)

	return s.conn.Close()
}

// openSocket is a private helper function to establish connection to Alpaca ws endpoint
func (s *Stream) openSocket() (*websocket.Conn, error) {
	scheme := "wss"
	ub, _ := url.Parse(s.base)
	if ub.Scheme == "http" {
		scheme = "ws"
	}
	u:= url.URL{Scheme: scheme, Host: ub.Host, Path: "/stream"}
	connectionAttempts := 0
	for connectionAttempts < MaxConnectionAttempts {
		connectionAttempts++
		c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err == nil {
			return c, nil
		}
		if connectionAttempts == MaxConnectionAttempts {
			return nil, err
		}
		time.Sleep(1 * time.Second)
	}

	return nil, errors.New(MaxRetriesExceededErr)

}

// isAuthenticated checks to see if the current alpaca stream is authenticated
func (s *Stream) isAuthenticated() bool {
	return s.authenticated.Load().(bool)
}

func (s *Stream) findStockHandler(stream string) chan StreamQuote {
	if v, ok := s.quoteHandlers.Load(stream); ok {
		return v.(chan StreamQuote)
	}else{
		c := stream[:strings.Index(stream, ".")]
		if v, ok := s.quoteHandlers.Load(c + ".*"); ok {
			return v.(chan StreamQuote)
		}
	}
	return nil
}

// findHandler ...
func (s *Stream) findHandler(stream string) chan interface{} {
	if v, ok := s.handlers.Load(stream); ok {
		return v.(chan interface{})
	}
	if strings.HasPrefix(stream, "Q.") ||
		strings.HasPrefix(stream, "T.") ||
		strings.HasPrefix(stream, "AM.") {
		c := stream[:strings.Index(stream, ".")]
		if v, ok := s.handlers.Load(c + ".*"); ok {
			return v.(chan interface{})
		}
	}
	return nil
}

// auth is a private helper function that authenticates the connection to the Alpaca API
func (s *Stream) auth() (err error) {
	s.Lock()
	defer s.Unlock()
	if s.isAuthenticated() {
		return nil
	}
	authRequest := ClientMsg{
		Action: "authenticate",
		Data: map[string]interface{}{
			"key_id":     s.id,
			"secret_key": s.secret,
		},
	}
	if err = s.conn.WriteJSON(authRequest); err != nil {
		return nil
	}
	msg := ServerMsg{}
	// ensure the auth response comes in a timely manner
	err = s.conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		return err
	}
	defer s.conn.SetReadDeadline(time.Time{})
	if err = s.conn.ReadJSON(&msg); err != nil {
		return
	}
	m := msg.Data.(map[string]interface{})
	if !strings.EqualFold(m["status"].(string), "authorized") {
		return errors.New(FailedAuthErr)
	}

	// change authenticate state to true
	s.authenticated.Store(true)

	return nil
}

// sub is a private function that listens for data on the socket
func (s *Stream) sub(channel string) (err error) {
	subReq := ClientMsg{
		Action: "listen",
		Data: map[string]interface{}{
			"streams": []interface{}{
				channel,
			},
		},
	}
	if err = s.conn.WriteJSON(subReq); err != nil {
		return
	}

	return
}

// unsub is a private function that stops listening on the wire
func (s *Stream) unsub(channel string) (err error) {
	subReq := ClientMsg{
		Action: "unlisten",
		Data: map[string]interface{}{
			"streams": []interface{}{
				channel,
			},
		},
	}
	if err = s.conn.WriteJSON(subReq); err != nil {
		return
	}

	return
}

// reconnect is a private function that tries to establish a new connection if we have failed to open a websocket connection
func (s *Stream) reconnect() error {
	// force authentication field to be false
	s.authenticated.Store(false)
	// open new connection
	conn, err := s.openSocket()
	if err != nil {
		return err
	}
	// set Stream connection to new websocket connection
	s.conn = conn
	if err := s.auth(); err != nil {
		return err
	}
	// todo: is there a better way of doing this
	s.handlers.Range(func(key, value interface{}) bool {
		// there should be no errors if we've previously successfully connected
		s.sub(key.(string))

		return true
	})

	return nil
}


