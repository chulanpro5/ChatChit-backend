package room

import (
	"test-chat/internal/common"
	"test-chat/internal/entity"
	"test-chat/internal/util"
)

type RoomService struct {
	common *common.Common
}

func NewRoomService(common *common.Common) *RoomService {
	return &RoomService{common: common}
}

func (rs *RoomService) CreateRoom(name string) (entity.Room, error) {
	room := entity.Room{
		ID:   0,
		Name: name,
	}

	result := rs.common.Database.Create(&room)
	if result.Error != nil {
		return entity.Room{}, result.Error
	}

	return room, nil
}

func (rs *RoomService) GetRoom(id uint) (entity.Room, error) {
	var room entity.Room

	result := rs.common.Database.First(&room, id)
	if result.Error != nil {
		return entity.Room{}, result.Error
	}

	return room, nil
}

func (rs *RoomService) GetRoomsByUserID(userID uint) ([]entity.Room, error) {
	var rooms []entity.Room

	strUserID, _ := util.UIntToStr(userID)
	result := rs.common.Database.Where("id IN (SELECT room_id FROM user_rooms WHERE user_id = ?)", strUserID).Find(&rooms)

	if result.Error != nil {
		return []entity.Room{}, nil
	}

	return rooms, nil
}

func (rs *RoomService) GetUsersByRoomID(roomID uint) ([]entity.User, error) {
	var users []entity.User
	strRoomID, _ := util.UIntToStr(roomID)
	result := rs.common.Database.Where("id IN (SELECT user_id FROM user_rooms WHERE room_id = ?)", strRoomID).Find(&users)
	if result.Error != nil {
		return []entity.User{}, result.Error
	}

	return users, nil
}

func (rs *RoomService) AddUserToRoom(userID uint, roomID uint) error {
	var user entity.User
	var room entity.Room

	rs.common.Database.First(&user, userID)
	rs.common.Database.First(&room, roomID)

	result := rs.common.Database.Model(&room).Association("Members").Append(&user)

	return result
}
