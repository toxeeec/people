package ws

import (
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/service/message"
)

type Hub struct {
	clients      map[uint][]*Client
	register     chan *Client
	unregister   chan *Client
	notification chan people.Notification
	cs           message.Service
}

func NewHub(notification chan people.Notification, cs message.Service) *Hub {
	return &Hub{
		clients:      make(map[uint][]*Client),
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
			clients := h.clients[c.ID]
			clients = append(clients, c)
			h.clients[c.ID] = clients
		case c := <-h.unregister:
			c.Conn.Close()
			close(c.Send)
			clients := h.clients[c.ID]
			for i := range clients {
				if clients[i] == c {
					clients = append(clients[:i], clients[i+1:]...)
					break
				}
			}
			h.clients[c.ID] = clients

		case notif := <-h.notification:
			clients := h.clients[notif.To]
			for _, c := range clients {
				c.Send <- notif
			}
		}
	}
}
