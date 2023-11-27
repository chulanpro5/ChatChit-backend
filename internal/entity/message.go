package entity

import "time"

type Message struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	Content   string    `bson:"content" json:"content"`
	RoomID    string    `bson:"roomId" json:"roomId"`
	SenderID  string    `bson:"senderId" json:"senderId"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}
