package message

import (
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

func (s *Service) GetMessages(filter *GetMessagesFilter) ([]entity.MessageResponse, uint64, error) {
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
	err := query.Find(&messages).Error
	if err != nil {
		return nil, 0, err
	}

	var messageResponses []entity.MessageResponse
	for _, message := range messages {
		messageResponses = append(messageResponses, entity.MessageResponse{
			ID:       message.ID,
			RoomId:   message.RoomId,
			Content:  message.Content,
			SenderId: message.SenderId,
			Metadata: message.Metadata,
		})
	}

	return messageResponses, uint64(totalItems), nil
}
