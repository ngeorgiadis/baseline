package wshub

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

// WSHub ...
type WSHub struct {
	clients    map[*WSClient]bool
	Broadcast  chan *WSMessage
	Register   chan *WSClient
	Unregister chan *WSClient
}

// WSClient ...
type WSClient struct {
	Hub    *WSHub
	Conn   *websocket.Conn
	Send   chan []byte
	UserID string
}

// WSMessage ...
type WSMessage struct {
	Message []byte
	Client  *WSClient
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 2 * 60 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 8) / 10
	// pingPeriod = 15 * time.Second

	// Maximum message size allowed from peer.
	maxMessageSize = 1024 * 1024
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// New create a new hub
func New() *WSHub {
	return &WSHub{
		Broadcast:  make(chan *WSMessage),
		Register:   make(chan *WSClient),
		Unregister: make(chan *WSClient),
		clients:    make(map[*WSClient]bool),
	}
}

// Run ...
func (h *WSHub) Run() {

	for {
		select {
		case client := <-h.Register:
			// register client
			h.clients[client] = true
		case client := <-h.Unregister:
			// unregister client
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			fmt.Println(len(h.clients))
			// send message to all clients
			for client := range h.clients {

				// for now broadcast the message
				// only to client connections
				// that have same userID
				// In the future we will
				// edit this section to send messages to
				// users that share tabs

				if client.UserID == message.Client.UserID {
					select {
					case client.Send <- message.Message:
					default:
						close(client.Send)
						delete(h.clients, client)
					}
				}
			}
		}
	}
}

// ReadPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *WSClient) ReadPump() {

	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)

	// c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	// c.Conn.SetPongHandler(func(a string) error {
	// 	fmt.Println("PONG: ", a)
	// 	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	// 	return nil
	// })

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("%v", err.Error())
			}
			break
		}

		// pong message do not broadcast
		if string(message) == "c" {
			continue
		}

		c.Hub.Broadcast <- &WSMessage{
			Message: message,
			Client:  c,
		}
	}
}

// WritePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *WSClient) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				fmt.Printf("%v", err.Error())
				return
			}
			w.Write(message)
			fmt.Println(string(message))

			if err := w.Close(); err != nil {
				fmt.Printf("%v", err.Error())
				return
			}
		case <-ticker.C:

			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))

			// send ping message
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				fmt.Printf("%v", err.Error())
				return
			}
			w.Write([]byte("d"))
			if err := w.Close(); err != nil {
				fmt.Printf("%v", err.Error())
				return
			}
			// if err := c.Conn.WriteMessage(websocket.TextMessage, []byte("beat")); err != nil {
			// 	return
			// }
		}
	}
}
