package user

import (
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

func (s *Service) GetUser(userId string) (entity.User, error) {
	var user entity.User
	result := s.common.Database.Where("id = ?", userId).First(&user)
	return user, result.Error
}
