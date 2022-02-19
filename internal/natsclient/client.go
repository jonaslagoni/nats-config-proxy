package natsclient

import (
	"github.com/nats-io/nats.go"

	"github.com/nats-io/nats-config-proxy/internal/models"
)

type Client struct {
	// nc is the encoded connection
	nc *nats.EncodedConn
}

// NewServer returns a configured server.
func NewNatsClient() *Client {
	return &Client{}
}

func (s *Client) ConnectToX() {
	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()
	s.nc = c
}

func (s *Client) PublishToAccountsV1AuthIdentsDelete(m *models.User) error {
	return s.nc.Publish("accounts.v1.auth.idents.delete", m)
}

func (s *Client) SubscribeToAccountsV1AuthIdentsDelete(callback func(*models.User)) (*nats.Subscription, error) {
	return s.nc.Subscribe("accounts.v1.auth.idents.delete", callback)
}
