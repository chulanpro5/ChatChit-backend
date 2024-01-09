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

type Handler struct {
	common      *common.Common
	userService *Service
	authService *auth.Service
}

func NewUserHandler(common *common.Common) *Handler {
	return &Handler{
		common:      common,
		userService: NewUserService(common),
		authService: auth.NewAuthService(common),
	}
}

func (h *Handler) GetUser(ctx *fiber.Ctx) error {
	user, err := h.userService.GetUser(fmt.Sprint(ctx.Locals("userId")))
	if err != nil {
		return response.BadRequest(ctx, err, nil)
	}

	return response.SendSuccess(ctx, &Response{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
}
