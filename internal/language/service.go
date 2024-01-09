package language

import (
	"errors"
	"gorm.io/gorm"
	"test-chat/pkg/common"
	"test-chat/pkg/entity"
)

type Service struct {
	common *common.Common
}

func NewLanguageService(common *common.Common) *Service {
	return &Service{
		common: common,
	}
}

func (s *Service) GetLanguages() ([]entity.Language, error) {
	var languages []entity.Language
	err := s.common.Database.DB.
		Find(&languages).Error

	if err != nil {
		return nil, err
	}

	return languages, nil
}

func (s *Service) GetLanguage(languageId string) (*entity.Language, error) {
	var language entity.Language
	result := s.common.Database.Where("id = ?", languageId).First(&language)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, nil
	}

	return &language, nil
}
