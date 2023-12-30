package friend

type AddFriendRequest struct {
	FriendId uint `json:"friendId" validate:"required"`
}
