package entity

type Room struct {
	ID      uint    `gorm:"primaryKey:autoIncrement"`
	Name    string  `gorm:"unique"`
	Members []*User `gorm:"many2many:user_rooms;"`
}
