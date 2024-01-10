package client

import (
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"test-chat/internal/auth"
	"test-chat/internal/user"
	"test-chat/pkg/common"
	"test-chat/pkg/entity"
	"test-chat/pkg/response"
)

func NewClientRouter(router fiber.Router) {
	handler := NewClientHandler(common.GetCommon())

	clientRouter := router.Group("/client")
	clientRouter.Get("/ws/:id", handler.authService.Middleware, handler.Upgrade, handler.Connect)
}

type Handler struct {
	common        *common.Common
	authService   *auth.Service
	clientService *Service
	userService   *user.Service
}

func NewClientHandler(common *common.Common) *Handler {
	return &Handler{
		common:        common,
		authService:   auth.NewAuthService(common),
		clientService: NewClientService(common),
		userService:   user.NewUserService(common),
	}
}

func (h *Handler) Upgrade(ctx *fiber.Ctx) error {
	// IsWebSocketUpgrade returns true if the client
	// requested upgrade to the WebSocket protocol.
	if websocket.IsWebSocketUpgrade(ctx) {
		ctx.Locals("allowed", true)
		return ctx.Next()
	}
	return fiber.ErrUpgradeRequired
}

func (h *Handler) Connect(ctx *fiber.Ctx) error {
	if fmt.Sprint(ctx.Params("id")) != fmt.Sprint(ctx.Locals("userId")) {
		return response.Unauthorized(ctx, fiber.ErrUnauthorized)
	}
	return websocket.New(h.HandleWebSocketConnection)(ctx)
}

func (h *Handler) HandleWebSocketConnection(c *websocket.Conn) {
	cl := &Client{
		Common:  h.common,
		Conn:    c,
		Message: make(chan *entity.MessageTranslation, 10),
		Id:      c.Params("id"),
	}

	h.clientService.hub.AddClient(cl.Id, cl)
	zap.L().Debug(fmt.Sprintf("Client %s registered", cl.Id))

	go h.writeMessage(cl)
	h.readMessage(cl, c)
}
