package publisher

import (
	"context"
	"eden/modules/profile/infrastructure/reliquarium"
	reliquariumMsg "eden/modules/profile/infrastructure/reliquarium/messages"
)

type ServicePausePublisher struct {
	client reliquarium.Client
}

func NewServicePausePublisher(client reliquarium.Client) *ServicePausePublisher {
	return &ServicePausePublisher{client: client}
}

func (p *ServicePausePublisher) Publish(ctx context.Context, event reliquariumMsg.PauseConfirmationEvent) error {
	return p.client.SendPauseConfirmation(ctx, event)
}
