package room

import (
	"errors"
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
	roomRouter.Get("/", handler.authService.Middleware, handler.GetRooms)

	roomRouter.Post("/:id/add-member", handler.authService.Middleware, handler.AddMember)
	roomRouter.Delete("/:id/remove-member", handler.authService.Middleware, handler.RemoveMember)
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

	room, err := h.roomService.CreateRoom(fmt.Sprint(ctx.Locals("userId")), *body, "group")
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
	// Check if room exists
	roomFound, err := h.roomService.GetRoom(fmt.Sprint(ctx.Locals("userId")), ctx.Params("id"))
	if err != nil {
		return err
	}
	if roomFound == nil {
		return response.BadRequest(ctx, errors.New("room not found"), nil)
	}

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
	roomFound, err := h.userService.GetUser(body.MemberId)
	if err != nil {
		return err
	}
	if roomFound == nil {
		return response.BadRequest(ctx, errors.New("user not found"), nil)
	}

	// Check if room exists
	userFound, err := h.roomService.GetRoom(fmt.Sprint(ctx.Locals("userId")), ctx.Params("id"))
	if err != nil {
		return err
	}
	if userFound == nil {
		return response.BadRequest(ctx, errors.New("room not found"), nil)
	}

	// Check if user is already a member
	member, err := h.roomService.GetMember(ctx.Params("id"), body.MemberId)
	if err != nil {
		return err
	}
	if member != nil {
		return response.BadRequest(ctx, errors.New("user is already a member"), nil)
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

	// Check if member and room exists
	member, err := h.roomService.GetMember(ctx.Params("id"), body.MemberId)
	if err != nil {
		return err
	}

	if member == nil {
		return response.BadRequest(ctx, errors.New("member or room not found"), nil)
	}

	err = h.roomService.RemoveMember(body.MemberId, ctx.Params("id"))
	if err != nil {
		return err
	}

	return response.SendSuccess(ctx, nil)
}
