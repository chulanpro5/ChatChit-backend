package message

type GetMessagesFilter struct {
	RoomId string `json:"roomId" validate:"required"`
	Page   uint   `json:"page" validate:"number,gte=1"`
	Limit  uint   `json:"limit" validate:"number,gte=1,lte=100"`
}
