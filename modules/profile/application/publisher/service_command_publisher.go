package publisher

import (
	"context"
	"eden/modules/profile/application/publisher/interfaces"
	reliquariumMsg "eden/modules/profile/infrastructure/reliquarium/messages"
)

type ServiceCommandConfirmationPublisher struct {
	client interfaces.ReliquariumClient
}

func NewServiceCommandConfirmationPublisher(client interfaces.ReliquariumClient) *ServiceCommandConfirmationPublisher {
	return &ServiceCommandConfirmationPublisher{client: client}
}

func (p *ServiceCommandConfirmationPublisher) Publish(ctx context.Context, event *reliquariumMsg.CommandConfirmationEvent) error {
	return p.client.SendCommandConfirmation(ctx, event)
}
