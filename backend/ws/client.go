package ws

import (
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/service"
)

type Client struct {
	ID   uint
	Conn *websocket.Conn
	Send chan (people.Notification)
}

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

func (c *Client) pongHandler(string) error {
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	return nil
}

func (c *Client) readPump(h *Hub) {
	const (
		message = "message"
	)
	defer func() {
		h.unregister <- c
	}()
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(c.pongHandler)
	for {
		_, data, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
		msgType := gjson.Get(string(data), "type").String()
		switch msgType {
		case message:
			if err := h.cs.ReadMessage(c.ID, data); err != nil {
				c.Conn.WriteJSON(err)
				continue
			}
		default:
			c.Conn.WriteJSON(service.NewError(people.ValidationError, "Invalid message type"))
			continue
		}
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, nil)
			}
			if err := c.Conn.WriteJSON(msg); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Serve(h *Hub) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		conn, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
		if err != nil {
			return err
		}
		userID, _ := people.FromContext(ctx.Request().Context(), people.UserIDKey)
		c := &Client{ID: userID, Conn: conn, Send: make(chan people.Notification, 32)}
		h.register <- c
		go c.readPump(h)
		go c.writePump()
		return nil
	}
}
