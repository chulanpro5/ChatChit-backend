package room

import (
	"fmt"
	"test-chat/pkg/common"
	"test-chat/pkg/entity"
	"test-chat/pkg/util"
)

type Service struct {
	common *common.Common
}

func NewRoomService(common *common.Common) *Service {
	return &Service{common: common}
}

func (s *Service) CreateRoom(userId string, dto CreateRoomRequest) (entity.Room, error) {
	room := entity.Room{
		Name: dto.Name,
	}

	err := s.common.Database.DB.Create(&room).Error
	if err != nil {
		return room, err
	}

	// Add room creator as member
	err = s.AddMember(userId, fmt.Sprint(room.ID))
	if err != nil {
		return room, err
	}

	return room, nil
}

func (s *Service) GetRoom(userId string, roomId string) (entity.Room, error) {
	var room entity.Room

	err := s.common.Database.DB.
		Table("rooms").
		Joins("JOIN room_members ON room_members.room_id = rooms.id").
		Where("room_members.user_id = ? AND room_members.room_id = ?", userId, roomId).
		First(&room).Error

	if err != nil {
		return room, err
	}

	return room, nil
}

func (s *Service) GetRooms(userId string) ([]entity.Room, error) {
	var rooms []entity.Room

	err := s.common.Database.DB.
		Table("rooms").
		Joins("JOIN room_members ON room_members.room_id = rooms.id").
		Where("room_members.user_id = ?", userId).
		Find(&rooms).Error

	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (s *Service) GetMembers(roomId string) ([]entity.User, error) {
	var members []entity.User

	err := s.common.Database.DB.
		Table("users").
		Joins("JOIN room_members ON room_members.user_id = users.id").
		Where("room_members.room_id = ?", roomId).
		Find(&members).Error

	if err != nil {
		return nil, err
	}

	return members, nil
}

func (s *Service) AddMember(userIdString string, roomIdString string) error {
	userId, err := util.StrToUInt(userIdString)
	if err != nil {
		return err
	}

	roomId, err := util.StrToUInt(roomIdString)
	if err != nil {
		return err
	}

	roomMember := entity.RoomMember{
		RoomId: roomId,
		UserId: userId,
	}

	err = s.common.Database.DB.Create(&roomMember).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) RemoveMember(userIdString string, roomIdString string) error {
	userId, err := util.StrToUInt(userIdString)
	if err != nil {
		return err
	}

	roomId, err := util.StrToUInt(roomIdString)
	if err != nil {
		return err
	}

	err = s.common.Database.DB.
		Where("user_id = ? AND room_id = ?", userId, roomId).
		Delete(&entity.RoomMember{}).Error
	if err != nil {
		return err
	}

	return nil
}
