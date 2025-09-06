package server

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/coder/websocket"
)

type Manager struct {
	clients WsClientList
	sync.RWMutex
}

func (m *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {
	log.Println("New Connection")

	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	client := NewClient(conn, m)
	m.addClient(client)

	go client.ReadMessages()
}

func NewManager() *Manager {
	return &Manager{
		clients: make(WsClientList),
	}
}

func (m *Manager) addClient(client *WsClient) {
	m.Lock()
	defer m.Unlock()

	m.clients[client] = true
}
func (m *Manager) removeClient(client *WsClient) {
	m.Lock()
	defer m.Unlock()

	if m.clients[client] {
		client.connection.CloseNow()
		delete(m.clients, client)
	}
}

func (c *WsClient) ReadMessages() {

	defer func() {
		c.connection.CloseNow()
		c.manager.removeClient(c)
	}()

	ctx := context.Background()

	for {
		messageType, payload, err := c.connection.Read(ctx)
		if err != nil {
			break
		}

		log.Println(messageType)
		log.Println(string(payload))
	}
}
