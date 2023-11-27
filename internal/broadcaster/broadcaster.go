package broadcaster

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"test-chat/internal/client"
	"test-chat/internal/common"
	"test-chat/internal/entity"
	"test-chat/internal/redis"
)

type BroadcasterService struct {
	common        *common.Common
	redis         *redis.Client
	clientService *client.ClientService
}

func NewBroadcasterService(common *common.Common) *BroadcasterService {
	redisClient, err := redis.NewRedisClient(common.Config)
	if err != nil {
		panic(err)
	}
	return &BroadcasterService{
		common:        common,
		redis:         redisClient,
		clientService: client.NewClientService(common),
	}
}

func (b *BroadcasterService) CleanUp() {
	b.redis.CleanUp()
}

func (b *BroadcasterService) Run() {
	pubsub := b.redis.Client.Subscribe(context.Background(), "message")
	defer pubsub.Close()

	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
		// parse message to JSON
		var message entity.Message
		err := json.Unmarshal([]byte(msg.Payload), &message)
		if err != nil {
			zap.L().Error(fmt.Sprintf("Error parsing message: %s", err.Error()))
		}
		b.clientService.BroadcastMessage(&message)
	}
}
