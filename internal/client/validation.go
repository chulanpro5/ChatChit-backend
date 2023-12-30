package client

type SendMessageRequest struct {
	Content  string `json:"content"`
	RoomId   uint   `json:"roomId"`
	Metadata string `json:"metadata"`
}
