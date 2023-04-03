package ws

import (
	"fmt"
	"github.com/gofiber/websocket/v2"
	"strings"
	"time"
)

type Client struct {
	Conn     *websocket.Conn
	ClientId string `json:"clientId"`
	Username string `json:"username"`
	RoomId   string `json:"roomId"`
	Message  chan *Message
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

func (c *Client) ReadMessage(h *Hub) {
	defer func() {
		h.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			if strings.Contains(err.Error(), "Websocket: close") {
				fmt.Println("Close connection")
			}
			break
		}
		message := Message{
			Message:  string(m),
			ClientId: c.ClientId,
			RoomId:   c.RoomId,
			Username: c.Username,
		}
		h.Broadcast <- &message
	}
}

func (c *Client) WriteMessage() {
	defer func() {
		fmt.Println("Connection was closed")
	}()
	for {
		select {
		case message, ok := <-c.Message:
			if !ok {
				return
			}
			c.Conn.WriteJSON(message)
		}
	}
}
