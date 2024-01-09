package friend

type AddFriendRequest struct {
	FriendId uint `json:"friendId" validate:"required"`
}

type FindUserByEmailRequest struct {
	Email string `json:"email"`
}
