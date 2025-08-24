package server

import (
	"log"
	"net/url"
	"context"

	"github.com/gorilla/websocket"
)

func HandleAlpacaWs(ctx context.Context) {

	u, err := url.Parse(ALPACA_MARKET_WEBSOCKET_URL)
	if err != nil {
		log.Fatal("Error parsing websocket URL:", err)
	}
	log.Printf("Connecting to %s", u.String())

	c, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		log.Printf("Handshake failed with status %d", resp.StatusCode)
		log.Fatal("dial:", err)
	}

	err = authenticateAlpacaWs(c)
	if err != nil {
		return
	}

	defer c.Close()

	go func() {
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", msg)

		}
	}()

	for range ctx.Done() {
		log.Println("Context cancelled, closing websocket connection")
		c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		return
	}
}

func authenticateAlpacaWs(c *websocket.Conn) error {
	authMsg := `{"action": "auth", "key": "` + ALPACA_API_KEY + `", "secret": "` + ALPACA_SECRET_KEY + `"}`
	err := c.WriteMessage(websocket.TextMessage, []byte(authMsg))
	if err != nil {
		log.Println("Error authenticating", err)
		return err
	}
	return nil
}
