package response

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Messsage struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

func SendValidationError(ctx *fiber.Ctx, errs map[string]string) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(
		Messsage{
			Success:    false,
			Message:    "Validation Error",
			StatusCode: fiber.StatusBadRequest,
			Data:       errs,
		})
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	zap.L().Error(err.Error())
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	if code == fiber.StatusInternalServerError {
		return ctx.Status(code).JSON(
			Messsage{
				Success:    false,
				Message:    err.Error(),
				StatusCode: code,
			})
	} else if code == fiber.StatusUnauthorized {
		return ctx.Status(code).JSON(
			Messsage{
				Success:    false,
				Message:    "Unauthorized",
				StatusCode: code,
			})
	}

	return ctx.Status(code).JSON(
		Messsage{
			Success:    false,
			Message:    err.Error(),
			StatusCode: code,
		})
}
