package ws

import (
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/service/message"
)

type Hub struct {
	clients      map[uint]*Client
	register     chan *Client
	unregister   chan *Client
	notification chan people.Notification
	cs           message.Service
}

func NewHub(notification chan people.Notification, cs message.Service) *Hub {
	return &Hub{
		clients:      make(map[uint]*Client),
		register:     make(chan *Client),
		unregister:   make(chan *Client),
		notification: notification,
		cs:           cs,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.register:
			h.clients[c.ID] = c
		case c := <-h.unregister:
			if _, ok := h.clients[c.ID]; ok {
				c.Conn.Close()
				close(c.Send)
				delete(h.clients, c.ID)
			}
		case notif := <-h.notification:
			c, ok := h.clients[notif.To]
			if ok {
				c.Send <- notif
			}
		}
	}
}
