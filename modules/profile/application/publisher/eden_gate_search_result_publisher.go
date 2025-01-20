package publisher

import (
	"context"
	"eden/modules/profile/application/consumer/interfaces"
	pubIntf "eden/modules/profile/application/publisher/interfaces"
	"eden/modules/profile/infrastructure/eden_gate/messages"
	loggerIntf "eden/shared/logger/interfaces"
)

type EdenGateSearchResultPublisher struct {
	client pubIntf.EdenGateClient
	logger loggerIntf.Logger
}

func NewEdenGateSearchResultPublisher(client pubIntf.EdenGateClient, logger loggerIntf.Logger) interfaces.EdenGateSearchResultPublisher {
	return &EdenGateSearchResultPublisher{
		client: client,
		logger: logger,
	}
}

func (p *EdenGateSearchResultPublisher) Publish(ctx context.Context, msg messages.ProfileSearchCompletedEvent) error {
	p.logger.Info("[EdenGateSearchResultPublisher] Publishing result to Eden Gate")
	return p.client.SendSearchResult(ctx, msg)
}
