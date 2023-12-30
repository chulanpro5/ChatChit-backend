package auth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"test-chat/pkg/common"
	"test-chat/pkg/response"
)

func NewAuthRouter(router fiber.Router) {
	handler := NewAuthHandler(common.GetCommon())

	authRouter := router.Group("/auth")
	authRouter.Post("/register", handler.Register)
	authRouter.Post("/login", handler.Login)
	authRouter.Post("/logout", handler.authService.Middleware, handler.Logout)
}

type Handler struct {
	common      *common.Common
	authService *Service
}

func NewAuthHandler(common *common.Common) *Handler {
	return &Handler{
		common:      common,
		authService: NewAuthService(common),
	}
}

func (h *Handler) Register(ctx *fiber.Ctx) error {
	body := new(RegisterRequest)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	resp, err := h.authService.Register(*body)
	if err != nil {
		return err
	}

	return response.SendSuccess(ctx, resp)
}

func (h *Handler) Login(ctx *fiber.Ctx) error {
	body := new(LoginRequest)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	cookie, err := h.authService.Login(*body)
	if err != nil {
		return response.BadRequest(ctx, err, nil)
	}

	ctx.Cookie(&cookie)

	return response.SendSuccess(ctx, nil)
}

func (h *Handler) Logout(ctx *fiber.Ctx) error {
	cookie, err := h.authService.Logout(fmt.Sprint(ctx.Locals("userId")))
	if err != nil {
		return err
	}

	ctx.Cookie(&cookie)

	return response.SendSuccess(ctx, nil)
}
