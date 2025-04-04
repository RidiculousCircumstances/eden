package eden_gate

import (
	"context"
	pubIntf "eden/modules/profile/application/publisher/interfaces"
	"eden/modules/profile/infrastructure/eden_gate/messages"
	"github.com/RidiculousCircumstances/netherway/v2"
)

type Client struct {
	broker            netherway.MessageBroker
	searchResultQueue string
	exchange          string
}

func NewClient(broker netherway.MessageBroker) pubIntf.EdenGateClient {
	return &Client{
		broker:            broker,
		searchResultQueue: "eden_gate_profiles_search_result_queue",
		exchange:          "eden_gate_exchange",
	}
}

func (c *Client) SendSearchResult(ctx context.Context, msg messages.ProfileSearchCompletedEvent) error {
	return c.broker.Publish(ctx, c.exchange, c.searchResultQueue, msg)
}
