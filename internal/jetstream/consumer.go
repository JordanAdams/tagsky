package jetstream

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/url"

	"github.com/gorilla/websocket"
)

type Consumer struct {
	Handler *Handler
	host    string
	dialer  websocket.Dialer
	conn    *websocket.Conn
}

func NewConsumer(host string) (*Consumer, error) {
	_, err := url.Parse(host)
	if err != nil {
		return nil, fmt.Errorf("invalid consumer host (%s): %w", host, err)
	}

	return &Consumer{
		Handler: NewHandler(),
		host:    host,
		dialer:  *websocket.DefaultDialer,
	}, nil
}

func (c Consumer) readMessage() error {
	msgType, msg, err := c.conn.ReadMessage()
	if err != nil {
		slog.Error(fmt.Errorf("failed to read message: %w", err).Error())
	}

	switch msgType {
	case websocket.TextMessage:
		return c.Handler.Handle(msg)
	default:
		return fmt.Errorf("unsupported message type: %v", msgType)
	}
}

func (c *Consumer) Start(ctx context.Context) error {
	conn, _, err := c.dialer.DialContext(ctx, c.host, nil)
	if err != nil {
		fmt.Println("ERR", err)
		return fmt.Errorf("failed to connect: %w", err)
	}

	c.conn = conn

	done := make(chan bool, 1)
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				err := c.readMessage()
				if err != nil {
					slog.Error(fmt.Errorf("failed to read message: %w", err).Error())
				}
			}
		}
	}()

	<-ctx.Done()
	log.Println("DONE FROM CTX")
	done <- true

	return nil
}
