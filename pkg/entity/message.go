package entity

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Content   string         `json:"content"`
	RoomId    uint           `json:"roomId"`
	SenderId  uint           `json:"senderId"`
	User      User           `gorm:"foreignKey:SenderId" json:"user"`
	Metadata  string         `gorm:"unique" json:"metadata"`
}
