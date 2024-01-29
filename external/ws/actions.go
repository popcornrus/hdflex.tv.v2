package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/patrickmn/go-cache"
	"time"
)

type Socket struct {
	Channel string      `json:"channel"`
	Event   string      `json:"event"`
	Data    interface{} `json:"data"`
}

func Actions(c *cache.Cache, conn *websocket.Conn, data *Socket) {
	switch data.Event[7:] {
	case "subscribe":
		channelSubscribers, _ := c.Get(data.Channel)
		_subscribe(conn, data)

		switch channelSubscribers.(type) {
		case []*websocket.Conn:
			channelSubscribers = append(channelSubscribers.([]*websocket.Conn), conn)
			c.Set(data.Channel, channelSubscribers, time.Hour)
		default:
			c.Set(data.Channel, []*websocket.Conn{conn}, time.Hour)
		}

		break
	case "unsubscribe":
		channelSubscribers, _ := c.Get(data.Channel)
		_unsubscribe(conn, data)

		if channelSubscribers != nil {
			for i, subscriber := range channelSubscribers.([]*websocket.Conn) {
				if subscriber == conn {
					channelSubscribers = append(channelSubscribers.([]*websocket.Conn)[:i], channelSubscribers.([]*websocket.Conn)[i+1:]...)

					if len(channelSubscribers.([]*websocket.Conn)) == 0 {
						c.Delete(data.Channel)
						break
					}

					c.Set(data.Channel, channelSubscribers, time.Hour)
					break
				}
			}
		}

		break
	default:
		return
	}
}

func _subscribe(conn *websocket.Conn, data *Socket) {
	bytes, err := json.Marshal(Socket{
		Channel: data.Channel,
		Event:   "action:subscribed",
		Data:    nil,
	})
	if err != nil {
		return
	}

	err = conn.WriteMessage(websocket.TextMessage, bytes)
	if err != nil {
		return
	}
}

func _unsubscribe(conn *websocket.Conn, data *Socket) {
	bytes, err := json.Marshal(Socket{
		Channel: data.Channel,
		Event:   "action:unsubscribed",
		Data:    nil,
	})
	if err != nil {
		return
	}

	err = conn.WriteMessage(websocket.TextMessage, bytes)
	if err != nil {
		return
	}
}
