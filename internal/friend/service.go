package friend

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"test-chat/internal/user"
	"test-chat/pkg/common"
	"test-chat/pkg/entity"
	"test-chat/pkg/util"
)

type Service struct {
	common      *common.Common
	userService *user.Service
}

func NewFriendService(common *common.Common) *Service {
	return &Service{
		common:      common,
		userService: user.NewUserService(common),
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
			return false, result.Error
		}
	}

	return true, nil
}

func (s *Service) GetFriends(userId string) ([]entity.User, error) {
	var friends []entity.User
	err := s.common.Database.DB.
		Table("users").
		Joins("JOIN friendships ON friendships.friend_id = users.id").
		Where("friendships.user_id = ? AND users.deleted_at IS NULL AND friendships.deleted_at IS NULL", userId).
		Find(&friends).Error

	if err != nil {
		return nil, err
	}

	var friendResponses []entity.User
	for _, friend := range friends {
		friendResponses = append(friendResponses, entity.User{
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

func (s *Service) FindFriendByEmail(userId string, dto *FindUserByEmailRequest) (*entity.User, bool, error) {
	friend, err := s.userService.FindUserByEmail(dto.Email)
	if err != nil {
		return friend, false, err
	}

	if friend == nil {
		return friend, false, nil
	}

	isAdded, err := s.GetFriend(userId, fmt.Sprint(friend.ID))
	if err != nil {
		return friend, false, err
	}

	return friend, isAdded, nil
}
