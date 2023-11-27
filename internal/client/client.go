package client

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"test-chat/internal/common"
	"test-chat/internal/entity"

	"github.com/gorilla/websocket"
)

type Client struct {
	Common  *common.Common
	Conn    *websocket.Conn
	Message chan *entity.Message
	Id      string `json:"id"`
}

func (c *Client) writeMessage() {
	defer func() {
		c.Conn.Close()
	}()
	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) readMessage(hub *Hub) {
	defer func() {
		hub.RemoveClient(c.Id)
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				zap.L().Error(fmt.Sprintf("Unexpected close error: %v", err))
			}
			break
		}

		// parse message to JSON
		var msg entity.Message
		err = json.Unmarshal(m, &msg)
		if err != nil {
			zap.L().Error(fmt.Sprintf("Error parsing message: %s", err.Error()))
		}

		c.SendMessage(&msg)
	}
}

func (c *Client) SendMessage(message *entity.Message) {
	err := c.Common.MongoDb.InsertMessage(message)
	if err != nil {
		zap.L().Error(fmt.Sprintf("Error inserting message: %s", err.Error()))
	}

	zap.L().Debug(fmt.Sprintf("Message inserted: %s", message))

	// parse message to JSON
	messageJson, err := json.Marshal(message)
	if err != nil {
		zap.L().Error(fmt.Sprintf("Error parsing message: %s", err.Error()))
	}
	err = c.Common.Redis.Publish(context.Background(), "message", string(messageJson))
	if err != nil {
		zap.L().Error(fmt.Sprintf("Error publishing message: %s", err.Error()))
	}

}
