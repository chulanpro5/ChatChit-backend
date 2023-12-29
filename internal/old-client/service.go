package old_client

import (
	"fmt"
	"test-chat/internal/room"
	"test-chat/pkg/common"
	"test-chat/pkg/entity"
	"test-chat/pkg/util"
)

type ClientService struct {
	common      *common.Common
	roomService *room.RoomService
	hub         *Hub
}

func NewClientService(common *common.Common) *ClientService {
	return &ClientService{
		common:      common,
		roomService: room.NewRoomService(common),
		hub:         GetHubInstance(),
	}
}

func (c *ClientService) BroadcastMessage(message *entity.Message) error {
	// Find the old-client with the given id
	fmt.Println("BroadcastToRoom")
	fmt.Println(message)

	roomID, _ := util.StrToUInt(message.RoomID)
	fmt.Println(roomID)
	receivers, _ := c.roomService.GetUsersByRoomID(roomID)

	fmt.Println(receivers)

	for _, receiver := range receivers {
		fmt.Println(receiver)
		receiverID, _ := util.UIntToStr(receiver.ID)
		// Get the old-client from the hub
		receiverClient, err := c.hub.GetClient(receiverID)
		fmt.Println(receiverClient)

		if err != nil {
			return err
		}
		fmt.Println("Sending message to old-client")
		receiverClient.Message <- message
	}

	return nil
}
