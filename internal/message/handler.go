package message

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"math"
	"test-chat/internal/auth"
	"test-chat/internal/room"
	"test-chat/pkg/common"
	"test-chat/pkg/entity"
	"test-chat/pkg/response"
)

func NewMessageRouter(router fiber.Router) {
	handler := NewMessageHandler(common.GetCommon())

	clientRouter := router.Group("/message")
	clientRouter.Get("/", handler.authService.Middleware, handler.GetMessages)
}

type Handler struct {
	common         *common.Common
	messageService *Service
	authService    *auth.Service
	roomService    *room.Service
}

func NewMessageHandler(common *common.Common) *Handler {
	return &Handler{
		common:         common,
		messageService: NewMessageService(common),
		authService:    auth.NewAuthService(common),
		roomService:    room.NewRoomService(common),
	}
}

func (h *Handler) GetMessages(ctx *fiber.Ctx) error {
	userId := fmt.Sprint(ctx.Locals("userId"))

	filter := new(GetMessagesFilter)
	if err := ctx.QueryParser(filter); err != nil {
		return err
	}

	// Check if user is member of room
	_, err := h.roomService.GetRoom(userId, filter.RoomId)
	if err != nil {
		return err
	}

	messages, total, err := h.messageService.GetMessages(filter)
	if err != nil {
		return err
	}

	return response.SendSuccess(ctx, entity.Pagination{
		Rows:      messages,
		Page:      filter.Page,
		Limit:     filter.Limit,
		TotalRows: total,
		TotalPage: uint(math.Ceil(float64(total) / float64(filter.Limit))),
	})

}
