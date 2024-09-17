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

func Actions(c *cache.Cache, conn *websocket.Conn, data *Socket, token string) {
	switch data.Event[7:] {
	case "subscribe":
		_subscribe(conn, data)

		channelSubscribers := make(map[string][]*websocket.Conn)
		channelSubscribers[data.Channel] = append(channelSubscribers[data.Channel], conn)

		c.Set(token, channelSubscribers, time.Hour)

		break
	case "unsubscribe":
		stack, _ := c.Get(token)
		_unsubscribe(conn, data)

		if stack != nil {
			connections, ok := stack.(map[string][]*websocket.Conn)
			if ok {
				connectionsList, exists := connections[data.Channel]
				if exists {
					for i, connection := range connectionsList {
						if connection == conn {
							// Remove the connection from the slice
							connectionsList = append(connectionsList[:i], connectionsList[i+1:]...)
							break
						}
					}

					// Update the connections list in the map
					connections[data.Channel] = connectionsList

					// If the connections list is empty, remove the channel from the stack
					if len(connectionsList) == 0 {
						delete(connections, data.Channel)
					}

					// Update the cache with the modified stack
					c.Set(token, connections, time.Hour)
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
