package friend

import (
	"errors"
	"gorm.io/gorm"
	"test-chat/pkg/common"
	"test-chat/pkg/entity"
	"test-chat/pkg/util"
)

type Service struct {
	common *common.Common
}

func NewFriendService(common *common.Common) *Service {
	return &Service{
		common: common,
	}
}

func (s *Service) GetFriend(userId string, friendId string) (bool, error) {
	var friendship entity.Friendship
	result := s.common.Database.DB.
		Where("user_id = ? AND friend_id = ?", userId, friendId).
		First(&friendship)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			return false, nil
		}
	}

	return true, nil
}

func (s *Service) GetFriends(userId string) ([]entity.UserResponse, error) {
	var friends []entity.User
	err := s.common.Database.DB.
		Table("users").
		Joins("JOIN friendships ON friendships.friend_id = users.id").
		Where("friendships.user_id = ?", userId).
		Find(&friends).Error

	if err != nil {
		return nil, err
	}

	var friendResponses []entity.UserResponse
	for _, friend := range friends {
		friendResponses = append(friendResponses, entity.UserResponse{
			ID:    friend.ID,
			Name:  friend.Name,
			Email: friend.Email,
		})
	}

	return friendResponses, nil
}

func (s *Service) AddFriend(userIdString string, friendIdString string) error {
	userId, err := util.StrToUInt(userIdString)
	if err != nil {
		return err
	}

	friendId, err := util.StrToUInt(friendIdString)
	if err != nil {
		return err
	}

	friendship := entity.Friendship{
		UserId:   userId,
		FriendId: friendId,
	}

	friendship2 := entity.Friendship{
		UserId:   friendId,
		FriendId: userId,
	}

	result := s.common.Database.Create(&friendship)
	if result.Error != nil {
		return result.Error
	}

	result = s.common.Database.Create(&friendship2)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Service) RemoveFriend(userIdString string, friendIdString string) error {
	userId, err := util.StrToUInt(userIdString)
	if err != nil {
		return err
	}

	friendId, err := util.StrToUInt(friendIdString)
	if err != nil {
		return err
	}

	result := s.common.Database.Where("user_id = ? AND friend_id = ?", userId, friendId).Delete(&entity.Friendship{})
	if result.Error != nil {
		return result.Error
	}

	result = s.common.Database.Where("user_id = ? AND friend_id = ?", friendId, userId).Delete(&entity.Friendship{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
