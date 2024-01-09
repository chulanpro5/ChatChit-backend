package main

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
	"test-chat/internal/auth"
	"test-chat/internal/broadcaster"
	"test-chat/internal/client"
	"test-chat/internal/friend"
	"test-chat/internal/language"
	"test-chat/internal/message"
	"test-chat/internal/room"
	"test-chat/internal/user"
	"test-chat/pkg/common"
	"test-chat/pkg/response"
)

func main() {
	c := common.NewCommon()
	defer c.CleanUp()

	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	app := fiber.New(fiber.Config{
		ErrorHandler: response.ErrorHandler,
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}),
		fiberzap.New(fiberzap.Config{
			Logger: logger,
		}),
	)

	broadcasterService := broadcaster.NewBroadcasterService(common.GetCommon())
	go broadcasterService.Run()

	api := app.Group("/api/v1")
	auth.NewAuthRouter(api)
	user.NewUserRouter(api)
	room.NewRoomRouter(api)
	client.NewClientRouter(api)
	message.NewMessageRouter(api)
	friend.NewFriendRouter(api)
	language.NewLanguageRouter(api)

	err := app.Listen(c.Config.App.Address)
	if err != nil {
		panic(err)
	}
}
