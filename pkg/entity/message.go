package entity

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Content   string         `json:"content"`
	RoomId    uint           `json:"roomId"`
	SenderId  uint           `json:"senderId"`
	Sender    User           `gorm:"foreignKey:SenderId" json:"sender"`
	Metadata  string         `gorm:"unique" json:"metadata"`
}
