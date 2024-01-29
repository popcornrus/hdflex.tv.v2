package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/patrickmn/go-cache"
	"go-chat/external/ws"
	"log/slog"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // Allow all connections
}

var c *cache.Cache

func init() {
	c = cache.New(time.Hour*24, cache.NoExpiration)
}

var (
	conn *websocket.Conn
	err  error
)

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			slog.Error("Error closing connection: ", err)
		}
	}(conn)

	for {
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			slog.Error("Error reading message:", err)
			break // Exit the loop if there's a read error
		}

		if messageType == websocket.TextMessage {
			var s *ws.Socket

			err = json.Unmarshal(p, &s)
			if err != nil {
				slog.Error("Error unmarshalling socket: ", err)
				continue
			}

			if s.Channel == "" {
				slog.Error("Channel is required")
				continue
			}

			if s.Event == "" {
				slog.Error("Event is required")
				continue
			}

			if s.Event[0:6] == "action" {
				ws.Actions(c, conn, s)
				continue
			}

			if s.Event[0:4] == "send" {
				sendEvent(s)
				continue
			}
		}
	}
}

func sendEvent(data *ws.Socket) {
	data.Event = data.Event[5:]

	subscribedUsers, found := c.Get(data.Channel)
	if !found {
		slog.Error("Channel not found: ", data.Channel)
		return
	}

	fmt.Println(data)

	for _, user := range subscribedUsers.([]*websocket.Conn) {
		err := user.WriteJSON(data)
		if err != nil {
			continue
		}
	}
}

func main() {
	http.HandleFunc("/ws/echo", handleConnection)

	err := http.ListenAndServe(":5051", nil)
	if err != nil {
		return
	}
}
