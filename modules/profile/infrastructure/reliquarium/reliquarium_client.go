package reliquarium

import (
	"context"
	"eden/modules/profile/infrastructure/reliquarium/messages"
	brokerIntf "eden/shared/broker/interfaces"
)

type Client struct {
	broker            brokerIntf.MessageBroker
	confirmationQueue string
	exchange          string
}

func NewClient(broker brokerIntf.MessageBroker) *Client {
	return &Client{
		broker:            broker,
		confirmationQueue: "reliquarium.confirmation_queue",
		exchange:          "reliquarium_confirmation_exchange",
	}
}

func (c *Client) SendPauseConfirmation(ctx context.Context, msg messages.PauseConfirmationEvent) error {
	return c.broker.Publish(ctx, c.exchange, c.confirmationQueue, msg)
}
