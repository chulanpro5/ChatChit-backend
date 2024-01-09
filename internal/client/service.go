package client

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"test-chat/internal/room"
	"test-chat/internal/user"
	"test-chat/pkg/common"
	"test-chat/pkg/entity"
)

type Service struct {
	common      *common.Common
	roomService *room.Service
	userService *user.Service
	hub         *Hub
}

func NewClientService(common *common.Common) *Service {
	return &Service{
		common:      common,
		roomService: room.NewRoomService(common),
		userService: user.NewUserService(common),
		hub:         GetHubInstance(),
	}
}

func (c *Service) BroadcastMessage(message *entity.Message) error {
	// Find the client with the given id
	fmt.Println("BroadcastToRoom")
	fmt.Println(message)

	members, err := c.roomService.GetMembers(fmt.Sprint(message.RoomId))
	if err != nil {
		zap.L().Error(fmt.Sprintf("Error getting members: %s", err.Error()))
		return err
	}

	for _, receiver := range members {
		zap.L().Debug(fmt.Sprintf("Sending message to client: %s", fmt.Sprint(receiver.ID)))
		receiverClient, err := c.hub.GetClient(fmt.Sprint(receiver.ID))
		if err != nil {
			zap.L().Debug(fmt.Sprintf("Client not found: %s", err.Error()))
			continue
		}
		receiverClient.Message <- message

		zap.L().Debug(fmt.Sprintf("Message sent to client: %s", receiverClient.Id))
	}

	return nil
}

func (c *Service) SendMessage(message *entity.Message) {
	// Check if room exists
	_, err := c.roomService.GetRoom(fmt.Sprint(message.SenderId), fmt.Sprint(message.RoomId))
	if err != nil {
		zap.L().Info(fmt.Sprintf("Room not found: %s", err.Error()))
		return
	}

	sender, err := c.userService.GetUser(fmt.Sprint(message.SenderId))
	if err != nil {
		zap.L().Info(fmt.Sprintf("Error getting user: %s", err.Error()))
		return
	}

	// insert message to database
	var msg = entity.Message{
		RoomId:   message.RoomId,
		Content:  message.Content,
		SenderId: message.SenderId,
		Metadata: message.Metadata,
		Sender:   *sender,
	}
	err = c.common.Database.DB.Create(&msg).Error
	if err != nil {
		zap.L().Error(fmt.Sprintf("Error inserting message: %s", err.Error()))
		return
	}

	zap.L().Debug(fmt.Sprintf("Message inserted: %s", msg))

	// parse message to JSON
	messageJson, err := json.Marshal(msg)
	if err != nil {
		zap.L().Error(fmt.Sprintf("Error parsing message: %s", err.Error()))
	}

	// publish message to redis
	err = c.common.Redis.Publish(context.Background(), "message", string(messageJson))
	if err != nil {
		zap.L().Error(fmt.Sprintf("Error publishing message: %s", err.Error()))
	}

}
