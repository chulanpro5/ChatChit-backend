package entity

import "gorm.io/gorm"

type Friendship struct {
	gorm.Model
	UserId   uint `json:"userId"`
	FriendId uint `json:"friendId"`
	User     User `gorm:"foreignKey:UserId" json:"user"`
	Friend   User `gorm:"foreignKey:FriendId" json:"friend"`
}
