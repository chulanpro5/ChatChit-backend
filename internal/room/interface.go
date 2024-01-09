package room

import "test-chat/pkg/entity"

type WithLastMessage struct {
	ID          uint            `json:"id"`
	Name        string          `json:"name"`
	Type        string          `json:"type"`
	ImageBase64 string          `json:"avatar"`
	LastMessage *entity.Message `json:"lastMessage"`
}
