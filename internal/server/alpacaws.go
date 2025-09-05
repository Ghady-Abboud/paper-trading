package server

import (
	"context"
	"log"
	"net/url"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

type AuthenticationMessage struct {
	Action string `json:"action"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

type SubscriptionMessage struct {
	Action string   `json:"action"`
	Bars   []string `json:"bars"`
}

type Bars struct {
	T              string  `json:"T"`
	Symbol         string  `json:"S"`
	OPrice         float32 `json:"o"`
	HPrice         float32 `json:"h"`
	LPrice         float32 `json:"l"`
	CPrice         float32 `json:"c"`
	Volume         int     `json:"v"`
	VolumeWeighted float32 `json:"vw"`
	NumTrades      int     `json:"n"`
	TimeStamp      string  `json:"t"`
}

func HandleAlpacaWs(ctx context.Context) {
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

	var msgs []Bars

	for {
		err = wsjson.Read(ctx, c, &msgs)
		if err != nil {
			c.CloseNow()
			return
		}
		for _, m := range msgs {
			if m.T == "b" {
				log.Printf("recv: %s c=%v t=%s", m.Symbol, m.CPrice, m.TimeStamp)
			}
		}
	}
}

func authenticateAlpacaWs(ctx context.Context, c *websocket.Conn) error {
	msg := AuthenticationMessage{
		Action: "auth",
		Key:    ALPACA_API_KEY,
		Secret: ALPACA_SECRET_KEY,
	}
	return wsjson.Write(ctx, c, msg)
}

func subscribeChannel(ctx context.Context, c *websocket.Conn) error {
	msg := SubscriptionMessage{
		Action: "subscribe",
		Bars:   []string{"META"},
	}
	return wsjson.Write(ctx, c, msg)
}
