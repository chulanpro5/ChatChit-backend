package entity

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name string `gorm:"unique" json:"name"`
}

type RoomMember struct {
	gorm.Model
	RoomId uint `json:"roomId"`
	UserId uint `json:"userId"`
}
