package language

import (
	"github.com/gofiber/fiber/v2"
	"test-chat/pkg/common"
	"test-chat/pkg/response"
)

func NewLanguageRouter(router fiber.Router) {
	handler := NewLanguageHandler(common.GetCommon())

	userRouter := router.Group("/language")

	userRouter.Get("/", handler.GetLanguages)
}

type Handler struct {
	common          *common.Common
	languageService *Service
}

func NewLanguageHandler(common *common.Common) *Handler {
	return &Handler{
		common:          common,
		languageService: NewLanguageService(common),
	}
}

func (h *Handler) GetLanguages(ctx *fiber.Ctx) error {
	languages, err := h.languageService.GetLanguages()
	if err != nil {
		return err
	}

	return response.SendSuccess(ctx, languages)
}
