
package server

import (
	"github.com/nats-io/nats-rest-config-proxy/internal/models"
	"github.com/nats-io/nats.go"
)

type NatsClient struct {
	// nc is the encoded connection
	nc *nats.EncodedConn
}

// NewServer returns a configured server.
func NewNatsClient(opts *Options) *NatsClient {
	if opts == nil {
		opts = &Options{}
	}
	return &NatsClient{opts: opts}
}

func (s *NatsClient) connectToX() {
	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()
	s.nc = c
}


func (s *NatsClient) publishToAccountsV1AuthIdentsUpdated(m AnonymousSchema_1) error {
    return s.nc.Publish("accounts.v1.auth.idents.updated", m)
}



func (s *NatsClient) subscribeToAccountsV1AuthIdents(callback func(AnonymousSchema_1)) (*nats.Subscription, error) {
	return s.nc.Subscribe("accounts.v1.auth.idents", callback)
}

  