package websocket

import (
	"errors"
	"net/http"

	"github.com/gorilla/websocket"
)

// Listener websocket listener function
type Listener func(message string)

// Websocket interface
type Websocket interface {
	RegisterListener(listener Listener)
	Open(w http.ResponseWriter, r *http.Request) error
	Broadcast(message string) error
}

type ws struct {
	upgrader  websocket.Upgrader
	conn      *websocket.Conn
	listeners []Listener
}

// NewWebsocket initialize new websocket instance
func NewWebsocket() Websocket {
	wsupgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	return &ws{
		upgrader:  wsupgrader,
		listeners: []Listener{},
	}
}

func (h *ws) RegisterListener(listener Listener) {
	if listener == nil {
		return
	}

	h.listeners = append(h.listeners, listener)
}

func (h *ws) Open(w http.ResponseWriter, r *http.Request) error {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	h.conn = conn
	for {
		t, msg, err := h.conn.ReadMessage()
		if err != nil {
			break
		}

		// skip non text message
		if t != websocket.TextMessage {
			continue
		}

		// send message to all listener
		for _, listener := range h.listeners {
			listener(string(msg))
		}
	}

	return nil
}

func (h *ws) Broadcast(message string) error {
	if h.conn == nil {
		return errors.New("connection is closed")
	}

	return h.conn.WriteMessage(websocket.TextMessage, []byte(message))
}
