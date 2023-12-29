package entity

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Content  string `json:"content"`
	RoomId   uint   `json:"roomId"`
	SenderId uint   `json:"senderId"`
	User     User   `gorm:"foreignKey:SenderId,references:ID" json:"user"`
	Metadata string `gorm:"unique" json:"metadata"`
}
