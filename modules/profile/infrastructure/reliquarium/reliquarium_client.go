package reliquarium

import (
	"context"
	"eden/modules/profile/infrastructure/reliquarium/messages"
	"github.com/RidiculousCircumstances/netherway/v2"
)

type Client struct {
	broker            netherway.MessageBroker
	confirmationQueue string
	exchange          string
}

func NewClient(broker netherway.MessageBroker) *Client {
	return &Client{
		broker:            broker,
		confirmationQueue: "reliquarium.confirmation_queue",
		exchange:          "reliquarium_confirmation_exchange",
	}
}

func (c *Client) SendCommandConfirmation(ctx context.Context, msg *messages.CommandConfirmationEvent) error {
	return c.broker.Publish(ctx, c.exchange, c.confirmationQueue, msg)
}
