package friend

import (
	"test-chat/pkg/entity"
)

type FindUserByEmailResponse struct {
	IsFriendAdded bool         `json:"isFriendAdded"`
	User          *entity.User `json:"user"`
}
