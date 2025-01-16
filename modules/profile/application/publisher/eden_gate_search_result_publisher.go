package publisher

import (
	"context"
	"eden/modules/profile/application/consumer/interfaces"
	pubIntf "eden/modules/profile/application/publisher/interfaces"
	"eden/modules/profile/infrastructure/eden_gate/messages"
)

type EdenGateSearchResultPublisher struct {
	client pubIntf.EdenGateClient
}

func NewEdenGateSearchResultPublisher(client pubIntf.EdenGateClient) interfaces.EdenGateSearchResultPublisher {
	return &EdenGateSearchResultPublisher{
		client: client,
	}
}

func (p *EdenGateSearchResultPublisher) Publish(ctx context.Context, msg messages.ProfileSearchCompletedEvent) error {
	return p.client.SendSearchResult(ctx, msg)
}
