package broadcaster

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"test-chat/internal/old-client"
	"test-chat/pkg/common"
	"test-chat/pkg/entity"
	"test-chat/pkg/redis"
)

type BroadcasterService struct {
	common        *common.Common
	redis         *redis.Client
	clientService *old_client.ClientService
}

func NewBroadcasterService(common *common.Common) *BroadcasterService {
	redisClient, err := redis.NewRedisClient(common.Config)
	if err != nil {
		panic(err)
	}
	return &BroadcasterService{
		common:        common,
		redis:         redisClient,
		clientService: old_client.NewClientService(common),
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
