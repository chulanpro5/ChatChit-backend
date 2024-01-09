package user

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

func (s *Service) UpdateLanguage(userId string, languageIdString string) error {
	user, err := s.GetUser(userId)
	if err != nil {
		return err
	}

	languageId, err := util.StrToUInt(languageIdString)
	if err != nil {
		return err
	}

	user.LanguageId = languageId

	if err := s.common.Database.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (s *Service) GetLanguage(userId string) (*entity.Language, error) {
	var language entity.Language
	err := s.common.Database.DB.
		Table("languages").
		Joins("JOIN users ON users.language_id = languages.id").
		Where("users.id = ?", userId).
		First(&language).Error

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, nil
	}

	return &language, nil
}
