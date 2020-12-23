package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/ngeorgiadis/baseline/cmd/server/wshub"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  8 * 1024,
	WriteBufferSize: 8 * 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func serveWs(hub *wshub.WSHub) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		userID := r.Header.Get("X-User-ID")

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Printf("%v", err.Error())
			return
		}

		client := &wshub.WSClient{
			Hub:    hub,
			Conn:   conn,
			Send:   make(chan []byte, 8*1024),
			UserID: userID,
		}

		client.Hub.Register <- client

		// Allow collection of memory referenced by the caller by doing all work in
		// new goroutines.
		go client.WritePump()
		go client.ReadPump()

	}
}
