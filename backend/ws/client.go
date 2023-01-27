package ws

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/service"
)

type Client struct {
	ID   uint
	Conn *websocket.Conn
	Send chan (people.Message)
}

func (c *Client) readPump(h *Hub) {
	const (
		message = "message"
	)
	defer func() {
		h.unregister <- c
	}()
	for {
		_, data, err := c.Conn.ReadMessage()
		if err != nil {
			continue
		}
		msgType := gjson.Get(string(data), "type").String()
		switch msgType {
		case message:
			if err := h.cs.ReadMessage(c.ID, data); err != nil {
				c.Conn.WriteJSON(err)
				continue
			}
		default:
			{
				c.Conn.WriteJSON(service.NewError(people.ValidationError, "Invalid message type"))
				continue
			}
		}
	}
}

func (c *Client) writePump() {

}

func Serve(h *Hub) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		conn, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
		if err != nil {
			return err
		}
		userID, _ := people.FromContext(ctx.Request().Context(), people.UserIDKey)
		c := &Client{ID: userID, Conn: conn, Send: make(chan people.Message, 32)}
		h.register <- c
		go c.readPump(h)
		go c.writePump()
		return nil
	}
}
