package entity

import (
	"gorm.io/gorm"
	"time"
)

type Room struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Name        string         `gorm:"unique" json:"name"`
	Type        string         `json:"type"`
	ImageBase64 string         `json:"avatar"`
}

type RoomMember struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	RoomId    uint           `json:"roomId"`
	UserId    uint           `json:"userId"`
}
