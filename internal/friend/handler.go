package friend

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"test-chat/internal/auth"
	"test-chat/internal/room"
	"test-chat/pkg/common"
	"test-chat/pkg/response"
)

func NewFriendRouter(router fiber.Router) {
	handler := NewFriendHandler(common.GetCommon())

	clientRouter := router.Group("/friend")
	clientRouter.Get("/", handler.authService.Middleware, handler.GetFriends)
	clientRouter.Post("/", handler.authService.Middleware, handler.AddFriend)
	clientRouter.Delete("/:id", handler.authService.Middleware, handler.DeleteFriend)
	clientRouter.Post("/find", handler.authService.Middleware, handler.FindUserByEmail)
}

type Handler struct {
	common        *common.Common
	friendService *Service
	authService   *auth.Service
	roomService   *room.Service
}

func NewFriendHandler(common *common.Common) *Handler {
	return &Handler{
		common:        common,
		friendService: NewFriendService(common),
		authService:   auth.NewAuthService(common),
		roomService:   room.NewRoomService(common),
	}
}

func (h *Handler) GetFriends(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")

	friends, err := h.friendService.GetFriends(fmt.Sprint(userId))
	if err != nil {
		return err
	}

	return response.SendSuccess(ctx, friends)
}

func (h *Handler) AddFriend(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")

	friendId := new(AddFriendRequest)
	if err := ctx.BodyParser(friendId); err != nil {
		return err
	}

	if fmt.Sprint(userId) == fmt.Sprint(friendId.FriendId) {
		return response.BadRequest(ctx, nil, "Cannot add yourself as friend")
	}

	// Check if friend exists
	got, err := h.friendService.GetFriend(fmt.Sprint(userId), fmt.Sprint(friendId.FriendId))
	if err != nil {
		return err
	}
	if got == true {
		return response.BadRequest(ctx, nil, "Friend already exists")
	}

	err = h.friendService.AddFriend(fmt.Sprint(userId), fmt.Sprint(friendId.FriendId))
	if err != nil {
		return err
	}

	// Create room
	createdRoom, err := h.roomService.CreateRoom(fmt.Sprint(userId), room.CreateRoomRequest{
		Name: "Private chat",
	}, "private")
	if err != nil {
		return err
	}

	// Add friend to room
	err = h.roomService.AddMember(fmt.Sprint(friendId.FriendId), fmt.Sprint(createdRoom.ID))
	if err != nil {
		return err
	}

	return response.SendSuccess(ctx, nil)
}

func (h *Handler) DeleteFriend(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")

	friendId := ctx.Params("id")

	got, err := h.friendService.GetFriend(fmt.Sprint(userId), friendId)
	if err != nil {
		return err
	}
	if got != true {
		return response.BadRequest(ctx, nil, "Friend does not exists")
	}

	err = h.friendService.RemoveFriend(fmt.Sprint(userId), friendId)
	if err != nil {
		return err
	}

	return response.SendSuccess(ctx, nil)
}

func (h *Handler) FindUserByEmail(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")

	body := new(FindUserByEmailRequest)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	friend, isAdded, err := h.friendService.FindFriendByEmail(fmt.Sprint(userId), body)
	if err != nil {
		return response.BadRequest(ctx, err, nil)
	}

	fmt.Println(friend)

	return response.SendSuccess(ctx, FindUserByEmailResponse{
		IsFriendAdded: isAdded,
		User:          friend,
	})
}
