package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"test-chat/pkg/entity"
	"test-chat/pkg/response"
)

func (s *Service) Middleware(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return response.Unauthorized(ctx, err)
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	var user entity.User
	s.common.Database.Where("id = ?", claims.Issuer).First(&user)

	ctx.Locals("userId", user.ID)

	return ctx.Next()
}
