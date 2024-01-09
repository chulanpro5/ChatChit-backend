package entity

import (
	"gorm.io/gorm"
	"time"
)

type Friendship struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserId    uint           `json:"userId"`
	FriendId  uint           `json:"friendId"`
	User      User           `gorm:"foreignKey:UserId" json:"user"`
	Friend    User           `gorm:"foreignKey:FriendId" json:"friend"`
}
