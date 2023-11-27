package entity

type User struct {
	ID       uint    `gorm:"primaryKey"`
	Username string  `gorm:"unique"`
	Rooms    []*Room `gorm:"many2many:user_rooms;"`
}
