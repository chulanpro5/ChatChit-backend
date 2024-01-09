package client

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"go.uber.org/zap"
	"test-chat/pkg/common"
	"test-chat/pkg/entity"
	"test-chat/pkg/util"
)

type Client struct {
	Common  *common.Common
	Conn    *websocket.Conn
	Message chan *entity.Message
	Id      string `json:"id"`
}

func (h *Handler) writeMessage(c *Client) {
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

func (h *Handler) readMessage(c *Client, conn *websocket.Conn) {
	defer func() {
		h.clientService.hub.RemoveClient(c.Id)
		c.Conn.Close()
	}()

	for {
		_, m, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				zap.L().Info(fmt.Sprintf("Unexpected close error: %v", err))
			}
			break
		}

		// process message
		var msgRequest SendMessageRequest
		err = json.Unmarshal(m, &msgRequest)
		if err != nil {
			zap.L().Info(fmt.Sprintf("Error parsing message: %s", err.Error()))
		}

		// Get sender info
		senderId, _ := util.StrToUInt(c.Conn.Params("id"))
		sender, err := h.userService.GetUser(fmt.Sprint(senderId))
		if err != nil {
			zap.L().Info(fmt.Sprintf("Error getting user: %s", err.Error()))
		}

		msg := &entity.Message{
			RoomId:   msgRequest.RoomId,
			Content:  msgRequest.Content,
			SenderId: senderId,
			Sender: entity.User{
				ID:    sender.ID,
				Name:  sender.Name,
				Email: sender.Email,
			},
			Metadata: msgRequest.Metadata,
		}

		h.clientService.SendMessage(msg)
	}
}
