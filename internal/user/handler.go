package user

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"test-chat/internal/auth"
	"test-chat/pkg/common"
	"test-chat/pkg/response"
)

func NewUserRouter(router fiber.Router) {
	handler := NewUserHandler(common.GetCommon())

	userRouter := router.Group("/user")

	userRouter.Get("/", handler.authService.Middleware, handler.GetUser)
}

type UserHandler struct {
	common      *common.Common
	userService *Service
	authService *auth.Service
}

func NewUserHandler(common *common.Common) *UserHandler {
	return &UserHandler{
		common:      common,
		userService: NewUserService(common),
		authService: auth.NewAuthService(common),
	}
}

func (h *UserHandler) GetUser(ctx *fiber.Ctx) error {
	user, err := h.userService.GetUser(fmt.Sprint(ctx.Locals("userId")))
	if err != nil {
		return response.BadRequest(ctx, err, nil)
	}

	return response.SendSuccess(ctx, user)
}
