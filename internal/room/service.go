package room

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
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

func (s *Service) CreateRoom(userId string, dto CreateRoomRequest, roomType string) (entity.Room, error) {
	room := entity.Room{
		Name:        dto.Name,
		Type:        roomType,
		ImageBase64: util.ImageToBase64(util.GenerateAvatar(dto.Name)),
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

func (s *Service) GetRoom(userId string, roomId string) (*entity.Room, error) {
	var room entity.Room

	err := s.common.Database.DB.
		Table("rooms").
		Joins("JOIN room_members ON room_members.room_id = rooms.id").
		Where("room_members.user_id = ? AND room_members.room_id = ? AND room_members.deleted_at IS NULL", userId, roomId).
		First(&room).Error

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, nil
	}

	return &room, nil
}

func (s *Service) GetGroups(userId string) ([]entity.Room, error) {
	var rooms []entity.Room

	err := s.common.Database.DB.
		Table("rooms").
		Joins("JOIN room_members ON room_members.room_id = rooms.id").
		Where("room_members.user_id = ? AND type = 'group' AND room_members.deleted_at IS NULL", userId).
		Find(&rooms).Error

	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (s *Service) GetFriendChats(userId string) ([]entity.Room, error) {
	var rooms []entity.Room

	err := s.common.Database.DB.
		Table("rooms").
		Joins("JOIN room_members ON room_members.room_id = rooms.id").
		Where("room_members.user_id = ? AND rooms.type = 'private'", userId).
		Find(&rooms).Error

	if err != nil {
		return nil, err
	}

	// For each room, assign friend name to room name
	for i, room := range rooms {
		member, err := s.GetFriendChatMember(userId, fmt.Sprint(room.ID))
		if err != nil {
			return nil, err
		}
		if member != nil {
			rooms[i].Name = member.Name
		}
	}

	return rooms, nil
}

func (s *Service) GetFriendChatMember(userId string, roomId string) (*entity.User, error) {
	var member entity.User

	err := s.common.Database.DB.
		Table("users").
		Joins("JOIN room_members ON room_members.user_id = users.id").
		Where("room_members.room_id = ? AND room_members.user_id != ? AND room_members.deleted_at IS NULL", roomId, userId).
		First(&member).Error

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, nil
	}

	return &member, nil
}

func (s *Service) GetMembers(roomId string) ([]entity.User, error) {
	var members []entity.User

	err := s.common.Database.DB.
		Table("users").
		Joins("JOIN room_members ON room_members.user_id = users.id").
		Where("room_members.room_id = ? AND room_members.deleted_at IS NULL", roomId).
		Find(&members).Error

	if err != nil {
		return nil, err
	}

	return members, nil
}

func (s *Service) GetMember(roomId string, memberId string) (*entity.User, error) {
	var member entity.User

	err := s.common.Database.DB.
		Table("users").
		Joins("JOIN room_members ON room_members.user_id = users.id").
		Where("room_members.room_id = ? AND room_members.user_id = ? AND room_members.deleted_at IS NULL", roomId, memberId).
		First(&member).Error

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, nil
	}

	return &member, nil
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
