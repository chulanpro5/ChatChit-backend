package message

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"test-chat/internal/language"
	"test-chat/pkg/common"
	"test-chat/pkg/entity"
)

type Service struct {
	common *common.Common
}

func NewMessageService(common *common.Common) *Service {
	return &Service{
		common: common,
	}
}

func (s *Service) GetMessages(filter *GetMessagesFilter) ([]entity.Message, uint64, error) {
	query := s.common.Database.Model(&entity.Message{})
	query = query.Where("room_id = ?", filter.RoomId)
	query = query.Order("created_at DESC")

	query = query.Offset(int((filter.Page - 1) * filter.Limit)).Limit(int(filter.Limit))

	// Get total items
	totalItems := int64(0)
	if err := query.Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}

	var messages []entity.Message
	err := query.Preload("Sender").Find(&messages).Error
	if err != nil {
		return nil, 0, err
	}

	return messages, uint64(totalItems), nil
}

func (s *Service) FetchTranslation(messageId string, languageId string) (*entity.MessageTranslation, error) {
	zap.L().Info(fmt.Sprint(messageId, " ", languageId))
	// Check if translation exists
	var messageTranslation entity.MessageTranslation
	err := s.common.Database.DB.
		Table("message_translations").
		Where("message_id = ? AND language_id = ?", messageId, languageId).
		Preload("Message").Preload("Language").Preload("Message.Sender").
		First(&messageTranslation).Error

	if err == nil {
		return &messageTranslation, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Fetch translation
	// Get message
	var message entity.Message
	err = s.common.Database.DB.
		Table("messages").
		Where("id = ?", messageId).
		Preload("Sender").
		First(&message).Error

	if err != nil {
		return nil, err
	}

	// Get language
	var targetLanguage entity.Language
	err = s.common.Database.DB.
		Table("languages").
		Where("id = ?", languageId).
		First(&targetLanguage).Error

	if err != nil {
		return nil, err
	}

	translated, err := language.TranslateText(message.Content, targetLanguage.Code, fmt.Sprint(s.common.Config.Translation.ApiUrl, ":", s.common.Config.Translation.Port))
	if err != nil {
		return nil, err
	}

	// Insert translation
	messageTranslation = entity.MessageTranslation{
		MessageId:         message.ID,
		LanguageId:        targetLanguage.ID,
		TranslatedContent: translated,
		Message:           message,
		Language:          targetLanguage,
	}

	err = s.common.Database.DB.Create(&messageTranslation).Error
	if err != nil {
		return nil, err
	}

	return &messageTranslation, nil
}

func (s *Service) GetMessageTranslations(messageId string, languageId string) (*entity.MessageTranslation, error) {
	messageTranslation, err := s.FetchTranslation(messageId, languageId)
	if err != nil {
		return nil, err
	}

	return messageTranslation, nil
}
