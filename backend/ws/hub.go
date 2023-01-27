package ws

import (
	"fmt"

	"github.com/gorilla/websocket"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/service/chat"
	"github.com/toxeeec/people/backend/service/notification"
)

type Hub struct {
	clients      map[uint]*Client
	register     chan *Client
	unregister   chan *Client
	notification chan people.Notification
	cs           chat.Service
	ns           notification.Service
}

func NewHub(cs chat.Service, ns notification.Service) *Hub {
	return &Hub{
		clients:      make(map[uint]*Client),
		register:     make(chan *Client),
		unregister:   make(chan *Client),
		notification: make(chan people.Notification),
		cs:           cs,
		ns:           ns,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.register:
			{
				h.clients[c.ID] = c
			}
		case c := <-h.unregister:
			{
				if _, ok := h.clients[c.ID]; ok {
					c.Conn.Close()
					close(c.Send)
					delete(h.clients, c.ID)
				}
			}
		case notif := <-h.notification:
			{
				fmt.Printf("%+v\n", notif)
			}
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
