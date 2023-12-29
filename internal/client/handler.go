package client

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"log"
	"test-chat/internal/auth"
	"test-chat/pkg/common"
)

func NewClientRouter(router fiber.Router) {
	handler := NewClientHandler(common.GetCommon())

	clientRouter := router.Group("/client")
	clientRouter.Get("/ws/connect", handler.authService.Middleware, upgrade, handler.Connect)
}

type ClientHandler struct {
	common      *common.Common
	authService *auth.Service
}

func NewClientHandler(common *common.Common) *ClientHandler {
	return &ClientHandler{
		common:      common,
		authService: auth.NewAuthService(common),
	}
}

func upgrade(ctx *fiber.Ctx) error {
	// IsWebSocketUpgrade returns true if the client
	// requested upgrade to the WebSocket protocol.
	if websocket.IsWebSocketUpgrade(ctx) {
		ctx.Locals("allowed", true)
		return ctx.Next()
	}
	return fiber.ErrUpgradeRequired
}

func (h *ClientHandler) Connect(ctx *fiber.Ctx) error {
	return nil
}

func handleWebSocketConnection(c *websocket.Conn) {
	log.Println(c.Locals("allowed"))  // true
	log.Println(c.Params("id"))       // 123
	log.Println(c.Query("v"))         // 1.0
	log.Println(c.Cookies("session")) // ""

	var (
		mt  int
		msg []byte
		err error
	)

	for {
		if mt, msg, err = c.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", msg)

		if err = c.WriteMessage(mt, msg); err != nil {
			log.Println("write:", err)
			break
		}
	}
}
