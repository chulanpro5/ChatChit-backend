package old_client

import (
	"errors"
	"go.uber.org/zap"
	"sync"
)

type Hub struct {
	Clients map[string]*Client
	mu      sync.Mutex
}

var (
	hubInstance *Hub
	once        sync.Once
)

// GetHubInstance returns the singleton instance of Hub
func GetHubInstance() *Hub {
	once.Do(func() {
		hubInstance = &Hub{
			Clients: make(map[string]*Client),
		}
	})
	return hubInstance
}

func (h *Hub) AddClient(id string, client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.Clients[id] = client

	zap.L().Debug("Client added", zap.String("id", id))
}

func (h *Hub) RemoveClient(id string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.Clients, id)

	zap.L().Debug("Client removed", zap.String("id", id))
}

func (h *Hub) GetClient(id string) (*Client, error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if _, ok := h.Clients[id]; ok {

		zap.L().Debug("Client found", zap.String("id", id))
		return h.Clients[id], nil
	}
	zap.L().Debug("Client not found", zap.String("id", id))
	return nil, errors.New("old-client not found")
}
