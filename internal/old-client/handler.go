package old_client

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
	"test-chat/pkg/common"
	"test-chat/pkg/entity"
)

type Handler struct {
	common        *common.Common
	clientService *ClientService
}

func NewHandler(common *common.Common) *Handler {
	return &Handler{
		common:        common,
		clientService: NewClientService(common),
	}
}

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) Connect(c *gin.Context) {
	// TODO: auth
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: get info from auth
	clientId := c.Query("id")

	zap.L().Debug(fmt.Sprintf("Client %s connected", clientId))

	cl := &Client{
		Common:  h.common,
		Conn:    conn,
		Message: make(chan *entity.Message, 10),
		Id:      clientId,
	}

	h.clientService.hub.AddClient(cl.Id, cl)
	zap.L().Debug(fmt.Sprintf("Client %s registered", cl.Id))

	go cl.writeMessage()
	go cl.readMessage(h.clientService.hub)
	zap.L().Debug(fmt.Sprintf("Client %s started listening", cl.Id))
}
