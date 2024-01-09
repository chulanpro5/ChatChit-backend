package broadcaster

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"test-chat/internal/client"
	"test-chat/pkg/common"
	"test-chat/pkg/entity"
	"test-chat/pkg/redis"
)

type Service struct {
	common        *common.Common
	redis         *redis.Client
	clientService *client.Service
}

func NewBroadcasterService(common *common.Common) *Service {
	redisClient, err := redis.NewRedisClient(common.Config)
	if err != nil {
		panic(err)
	}
	return &Service{
		common:        common,
		redis:         redisClient,
		clientService: client.NewClientService(common),
	}
}

func (b *Service) CleanUp() {
	b.redis.CleanUp()
}

func (b *Service) Run() {
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

		// Broadcast message to room
		err = b.clientService.BroadcastMessage(&message)
		if err != nil {
			zap.L().Error(fmt.Sprintf("Error broadcasting message: %s", err.Error()))
		}
	}
}
