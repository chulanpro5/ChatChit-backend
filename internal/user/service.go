package user

import (
	"errors"
	"gorm.io/gorm"
	"test-chat/pkg/common"
	"test-chat/pkg/entity"
)

type Service struct {
	common *common.Common
}

func NewUserService(common *common.Common) *Service {
	return &Service{
		common: common,
	}
}

func (s *Service) GetUser(userId string) (*entity.User, error) {
	var user entity.User
	result := s.common.Database.Where("id = ?", userId).First(&user)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, nil
	}
	return &user, nil
}

func (s *Service) FindUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	result := s.common.Database.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, nil
	}

	return &user, nil
}
