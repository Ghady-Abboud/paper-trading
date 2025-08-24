package server

import (
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

func HandleAlpacaWs() {

	log.SetFlags(0)
	interrupt := make(chan os.Signal, 1)
	done := make(chan struct{})

	signal.Notify(interrupt, os.Interrupt)

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
		defer close(done)
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", msg)

		}
	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		}
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
