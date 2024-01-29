package ws

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"go.uber.org/fx"
	"log/slog"
	"os"
	"os/signal"
)

type WebSocketClient struct {
	log     *slog.Logger
	Channel chan Socket
	conn    *websocket.Conn
}

func NewWebSocketClient(log *slog.Logger) *WebSocketClient {
	return &WebSocketClient{
		log:     log,
		Channel: make(chan Socket),
	}
}

func RunWebSocketClient(
	lc fx.Lifecycle,
	log *slog.Logger,
	w *WebSocketClient,
) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				interrupt := make(chan os.Signal, 1)
				signal.Notify(interrupt, os.Interrupt)

				u := "ws://websocket:5051/ws/echo"

				// Create a WebSocket dialer
				dialer := websocket.DefaultDialer
				conn, _, err := dialer.Dial(u, nil)
				if err != nil {
					w.log.Error("dial:", err)
				}

				w.conn = conn
				defer conn.Close()

				go w.Send()

				for {
					select {
					case <-interrupt:
						w.log.Info("interrupt")
						// To cleanly close a connection, a client should send a close
						// frame and wait for the server to close the connection.
						err := w.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
						if err != nil {
							w.log.Error("write close:", err)
							return
						}

						return
					}
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Error("Server stopped")
			return nil
		},
	})
}

func (w *WebSocketClient) Send() {
	for {
		select {
		case message := <-w.Channel:
			err := w.conn.WriteJSON(message)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func SendEvent(w *WebSocketClient, message Socket) {
	go func() {
		w.Channel <- message
	}()
}
