package ws

import (
	"sync"
)

var WsManager *Hub

type Message struct {
	From   int64       `json:"from"`
	To     int64       `json:"to"`
	Action int         `json:"action"`
	Data   interface{} `json:"data"`
}

type Hub struct {
	clients map[int64]*Client
	lock    sync.RWMutex
}

func InitHub() {
	WsManager = &Hub{
		clients: make(map[int64]*Client),
	}
}

func (h *Hub) RouteMessage() {

}

func (h *Hub) Put(userId int64, client *Client) {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.clients[userId] = client
}

func (h *Hub) Get(userId int64) *Client {
	h.lock.RLock()
	defer h.lock.RUnlock()

	return h.clients[userId]
}

func (h *Hub) Remove(userId int64) {
	h.lock.Lock()
	defer h.lock.Unlock()

	delete(h.clients, userId)
}
