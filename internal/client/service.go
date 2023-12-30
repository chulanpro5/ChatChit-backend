package client

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"test-chat/internal/room"
	"test-chat/pkg/common"
	"test-chat/pkg/entity"
)

type Service struct {
	common      *common.Common
	roomService *room.Service
	hub         *Hub
}

func NewClientService(common *common.Common) *Service {
	return &Service{
		common:      common,
		roomService: room.NewRoomService(common),
		hub:         GetHubInstance(),
	}
}

func (c *Service) BroadcastMessage(message *entity.MessageResponse) error {
	// Find the old-client with the given id
	fmt.Println("BroadcastToRoom")
	fmt.Println(message)

	members, err := c.roomService.GetMembers(fmt.Sprint(message.RoomId))
	if err != nil {
		zap.L().Error(fmt.Sprintf("Error getting members: %s", err.Error()))
		return err
	}

	for _, receiver := range members {
		receiverClient, err := c.hub.GetClient(fmt.Sprint(receiver.ID))
		if err != nil {
			return err
		}
		receiverClient.Message <- message

		zap.L().Debug(fmt.Sprintf("Message sent to client: %s", receiverClient.Id))
	}

	return nil
}

func (c *Service) SendMessage(message *entity.MessageResponse) {
	// Check if room exists
	_, err := c.roomService.GetRoom(fmt.Sprint(message.SenderId), fmt.Sprint(message.RoomId))
	if err != nil {
		zap.L().Info(fmt.Sprintf("Room not found: %s", err.Error()))
		return
	}

	// insert message to database
	var msg = entity.Message{
		RoomId:   message.RoomId,
		Content:  message.Content,
		SenderId: message.SenderId,
		Metadata: message.Metadata,
	}
	err = c.common.Database.DB.Create(&msg).Error
	if err != nil {
		zap.L().Error(fmt.Sprintf("Error inserting message: %s", err.Error()))
		return
	}

	zap.L().Debug(fmt.Sprintf("Message inserted: %s", message))

	// parse message to JSON
	messageJson, err := json.Marshal(message)
	if err != nil {
		zap.L().Error(fmt.Sprintf("Error parsing message: %s", err.Error()))
	}

	// publish message to redis
	err = c.common.Redis.Publish(context.Background(), "message", string(messageJson))
	if err != nil {
		zap.L().Error(fmt.Sprintf("Error publishing message: %s", err.Error()))
	}

}
