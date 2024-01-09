package friend

import "test-chat/internal/user"

type FindUserByEmailResponse struct {
	IsFriendAdded bool           `json:"isFriendAdded"`
	User          *user.Response `json:"user"`
}
