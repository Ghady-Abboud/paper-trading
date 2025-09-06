package server

import (
	"github.com/coder/websocket"
)

type WsClientList map[*WsClient]bool

type WsClient struct {
	connection *websocket.Conn
	manager    *Manager
}

func NewClient(conn *websocket.Conn, manager *Manager) *WsClient {
	return &WsClient{
		connection: conn,
		manager:    manager,
	}
}
