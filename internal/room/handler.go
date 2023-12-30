package room

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"test-chat/internal/auth"
	"test-chat/internal/user"
	"test-chat/pkg/common"
	"test-chat/pkg/response"
)

func NewRoomRouter(router fiber.Router) {
	handler := NewRoomHandler(common.GetCommon())

	roomRouter := router.Group("/room")
	roomRouter.Post("/create", handler.authService.Middleware, handler.CreateRoom)
	roomRouter.Get("/:id", handler.authService.Middleware, handler.GetRoom)
	roomRouter.Get("/", handler.authService.Middleware, handler.GetRooms)

	roomRouter.Post("/:id/add-member", handler.authService.Middleware, handler.AddMember)
	roomRouter.Post("/:id/remove-member", handler.authService.Middleware, handler.RemoveMember)
	roomRouter.Get("/:id/list-members", handler.authService.Middleware, handler.GetMembers)
}

type Handler struct {
	common      *common.Common
	roomService *Service
	authService *auth.Service
	userService *user.Service
}

func NewRoomHandler(common *common.Common) *Handler {
	return &Handler{
		common:      common,
		roomService: NewRoomService(common),
		authService: auth.NewAuthService(common),
		userService: user.NewUserService(common),
	}
}

func (h *Handler) CreateRoom(ctx *fiber.Ctx) error {
	body := new(CreateRoomRequest)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	room, err := h.roomService.CreateRoom(fmt.Sprint(ctx.Locals("userId")), *body)
	if err != nil {
		return err
	}

	return response.SendSuccess(ctx, room)
}

func (h *Handler) GetRoom(ctx *fiber.Ctx) error {
	room, err := h.roomService.GetRoom(fmt.Sprint(ctx.Locals("userId")), ctx.Params("id"))
	if err != nil {
		return err
	}

	return response.SendSuccess(ctx, room)
}

func (h *Handler) GetRooms(ctx *fiber.Ctx) error {
	rooms, err := h.roomService.GetRooms(fmt.Sprint(ctx.Locals("userId")))
	if err != nil {
		return err
	}

	return response.SendSuccess(ctx, rooms)
}

func (h *Handler) GetMembers(ctx *fiber.Ctx) error {
	members, err := h.roomService.GetMembers(ctx.Params("id"))
	if err != nil {
		return err
	}

	return response.SendSuccess(ctx, members)
}

func (h *Handler) AddMember(ctx *fiber.Ctx) error {
	body := new(AddMemberRequest)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	// Check if user exists
	_, err := h.userService.GetUser(body.MemberId)
	if err != nil {
		return err
	}

	// Check if room exists
	_, err = h.roomService.GetRoom(fmt.Sprint(ctx.Locals("userId")), ctx.Params("id"))
	if err != nil {
		return err
	}

	err = h.roomService.AddMember(body.MemberId, ctx.Params("id"))
	if err != nil {
		return err
	}

	return response.SendSuccess(ctx, nil)
}

func (h *Handler) RemoveMember(ctx *fiber.Ctx) error {
	body := new(RemoveMemberRequest)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	// Check if room exists
	_, err := h.roomService.GetRoom(fmt.Sprint(ctx.Locals("userId")), ctx.Params("id"))
	if err != nil {
		return err
	}

	err = h.roomService.RemoveMember(body.MemberId, ctx.Params("id"))
	if err != nil {
		return err
	}

	return response.SendSuccess(ctx, nil)
}
