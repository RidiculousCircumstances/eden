package interfaces

import (
	"context"
	"eden/modules/profile/infrastructure/eden_gate/interfaces"
	"eden/modules/profile/infrastructure/queue/message"
)

type EdenGateSearchResultPublisher interface {
	Publish(ctx context.Context, msg interfaces.EdenGateSearchResultMessage) error
}

type StreamForgeMessageProcessor interface {
	Process(ctx context.Context, msg message.StreamForgeMessage) error
}

type TraceFaceMessageProcessor interface {
	Process(ctx context.Context, msg message.TraceFaceMessage) error
}

type EdenSearchMessageProcessor interface {
	Process(ctx context.Context, msg message.SearchProfileMessage) error
}
