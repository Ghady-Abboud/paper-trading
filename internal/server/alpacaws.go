package server

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

type authMessage struct {
	Action string `json:"action"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

type subChannelMessage struct {
	Action string   `json:"action"`
	Quotes []string `json:"quotes"`
	Bars   []string `json:"bars"`
}

func HandleAlpacaWs(ctx context.Context) {
	var msg interface{}

	log.SetFlags(0)
	u, err := url.Parse(ALPACA_MARKET_WEBSOCKET_URL)
	if err != nil {
		log.Fatal("Error parsing websocket URL:", err)
	}
	log.Printf("Connecting to %s", u.String())

	c, resp, err := websocket.Dial(ctx, u.String(), nil)

	if err != nil {
		log.Printf("Handshake failed with status %d", resp.StatusCode)
		log.Fatal("dial:", err)
	}
	defer c.CloseNow()

	err = authenticateAlpacaWs(ctx, c)
	if err != nil {
		log.Println("Error authenticating", err)
		return
	}

	err = subscribeChannel(ctx, c)
	if err != nil {
		log.Println("Error subscribing to channel", err)
		return
	}

	fmt.Println("we got here")
	for {
		err = wsjson.Read(ctx, c, &msg)
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", msg)
	}
}

func authenticateAlpacaWs(ctx context.Context, c *websocket.Conn) error {
	msg := authMessage{
		Action: "auth",
		Key:    ALPACA_API_KEY,
		Secret: ALPACA_SECRET_KEY,
	}
	return wsjson.Write(ctx, c, msg)
}

func subscribeChannel(ctx context.Context, c *websocket.Conn) error {
	msg := subChannelMessage{
		Action: "subscribe",
		Quotes: []string{"META"},
		Bars:   []string{"*"},
	}
	return wsjson.Write(ctx, c, msg)
}
