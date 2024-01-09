package user

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"test-chat/internal/auth"
	"test-chat/internal/language"
	"test-chat/pkg/common"
	"test-chat/pkg/response"
)

func NewUserRouter(router fiber.Router) {
	handler := NewUserHandler(common.GetCommon())

	userRouter := router.Group("/user")

	userRouter.Get("/", handler.authService.Middleware, handler.GetUser)
	userRouter.Put("/language", handler.authService.Middleware, handler.UpdateLanguage)
	userRouter.Get("/language", handler.authService.Middleware, handler.GetLanguage)
}

type Handler struct {
	common          *common.Common
	userService     *Service
	authService     *auth.Service
	languageService *language.Service
}

func NewUserHandler(common *common.Common) *Handler {
	return &Handler{
		common:          common,
		userService:     NewUserService(common),
		authService:     auth.NewAuthService(common),
		languageService: language.NewLanguageService(common),
	}
}

func (h *Handler) GetUser(ctx *fiber.Ctx) error {
	user, err := h.userService.GetUser(fmt.Sprint(ctx.Locals("userId")))
	if err != nil {
		return response.BadRequest(ctx, err, nil)
	}

	return response.SendSuccess(ctx, user)
}

func (h *Handler) UpdateLanguage(ctx *fiber.Ctx) error {
	body := new(UpdateLanguageRequest)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	// Check if language exists
	languageFound, err := h.languageService.GetLanguage(fmt.Sprint(body.LanguageId))
	if err != nil {
		return err
	}
	if languageFound == nil {
		return response.BadRequest(ctx, fmt.Errorf("language not found"), nil)
	}

	if err := h.userService.UpdateLanguage(fmt.Sprint(ctx.Locals("userId")), fmt.Sprint(body.LanguageId)); err != nil {
		return response.BadRequest(ctx, err, nil)
	}

	return response.SendSuccess(ctx, nil)
}

func (h *Handler) GetLanguage(ctx *fiber.Ctx) error {
	languageFound, err := h.userService.GetLanguage(fmt.Sprint(ctx.Locals("userId")))
	if err != nil {
		return err
	}

	return response.SendSuccess(ctx, languageFound)
}
