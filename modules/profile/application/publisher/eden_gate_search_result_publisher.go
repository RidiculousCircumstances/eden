package publisher

import (
	"context"
	"eden/modules/profile/application/consumer/interfaces"
	pubIntf "eden/modules/profile/application/publisher/interfaces"
	edenGateClientIntf "eden/modules/profile/infrastructure/eden_gate/interfaces"
)

type EdenGateSearchResultPublisher struct {
	client pubIntf.EdenGateClient
}

func NewEdenGateSearchResultPublisher(client pubIntf.EdenGateClient) interfaces.EdenGateSearchResultPublisher {
	return &EdenGateSearchResultPublisher{
		client: client,
	}
}

func (p *EdenGateSearchResultPublisher) Publish(ctx context.Context, msg edenGateClientIntf.ProfileSearchCompletedEvent) error {
	return p.client.SendSearchResult(ctx, msg)
}
