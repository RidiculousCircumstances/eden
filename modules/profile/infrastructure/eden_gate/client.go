package eden_gate

import (
	"context"
	pubIntf "eden/modules/profile/application/publisher/interfaces"
	"eden/modules/profile/infrastructure/eden_gate/interfaces"
	brokerIntf "eden/shared/broker/interfaces"
)

type Client struct {
	broker    brokerIntf.MessageBroker
	queueName string
	exchange  string
}

func NewClient(broker brokerIntf.MessageBroker) pubIntf.EdenGateClient {
	return &Client{
		broker: broker,
	}
}

func (c *Client) SendSearchResult(ctx context.Context, msg interfaces.ProfileSearchCompletedEvent) error {
	return c.broker.Publish(ctx, c.exchange, c.queueName, msg)
}
