package response

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func SendSuccess(ctx *fiber.Ctx, data interface{}) error {
	if data == nil {
		data = new(interface{})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		Messsage{
			Success:    true,
			Message:    "successfully",
			StatusCode: fiber.StatusOK,
			Data:       data,
		},
	)
}

func SendWithCode(ctx *fiber.Ctx, code int, message string, data interface{}, success bool) error {
	if data == nil {
		data = new(interface{})
	}
	return ctx.Status(code).JSON(
		Messsage{
			Success:    success,
			Message:    message,
			StatusCode: code,
			Data:       data,
		},
	)
}

func Unauthorized(ctx *fiber.Ctx, err error) error {
	zap.L().Info(err.Error())
	return SendWithCode(ctx, fiber.StatusUnauthorized, "unauthorized", nil, false)
}

func BadRequest(ctx *fiber.Ctx, err error, data interface{}) error {
	zap.L().Info(err.Error())
	return SendWithCode(ctx, fiber.StatusBadRequest, "malformed data", data, false)
}
