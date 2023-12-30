package entity

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Content  string `json:"content"`
	RoomId   uint   `json:"roomId"`
	SenderId uint   `json:"senderId"`
	User     User   `gorm:"foreignKey:SenderId" json:"user"`
	Metadata string `gorm:"unique" json:"metadata"`
}

type MessageResponse struct {
	ID       uint         `json:"id"`
	Content  string       `json:"content"`
	RoomId   uint         `json:"roomId"`
	SenderId uint         `json:"senderId"`
	User     UserResponse `json:"user"`
	Metadata string       `json:"metadata"`
}
