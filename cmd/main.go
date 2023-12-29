package main

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
	"test-chat/internal/auth"
	"test-chat/internal/room"
	"test-chat/internal/user"
	"test-chat/pkg/common"
)

func main() {
	c := common.NewCommon()
	defer c.CleanUp()

	logger, _ := zap.NewProduction()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}),
		fiberzap.New(fiberzap.Config{
			Logger: logger,
		}),
	)

	api := app.Group("/api/v1")
	auth.NewAuthRouter(api)
	user.NewUserRouter(api)
	room.NewRoomRouter(api)

	err := app.Listen(c.Config.App.Address)
	if err != nil {
		panic(err)
	}
}
