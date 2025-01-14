package publisher

import (
	"context"
	"eden/modules/profile/application/consumer/interfaces"
	edenClientIntf "eden/modules/profile/infrastructure/eden_gate/interfaces"
)

type EdenGateSearchResultPublisher struct {
	client edenClientIntf.Client
}

func NewEdenGateSearchResultPublisher(client edenClientIntf.Client) interfaces.EdenGateSearchResultPublisher {
	return &EdenGateSearchResultPublisher{
		client: client,
	}
}

func (p *EdenGateSearchResultPublisher) Publish(ctx context.Context, msg edenClientIntf.ProfileSearchCompletedEvent) error {
	return p.client.SendSearchResult(ctx, msg)
}
