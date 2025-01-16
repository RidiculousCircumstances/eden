package interfaces

import (
	"context"
	"eden/modules/profile/infrastructure/eden_gate/messages"
	"eden/modules/profile/infrastructure/queue/message"
)

type EdenGateSearchResultPublisher interface {
	Publish(ctx context.Context, msg messages.ProfileSearchCompletedEvent) error
}

type StreamForgeMessageProcessor interface {
	Process(ctx context.Context, msg message.SaveProfileCommand) error
}

type TraceFaceMessageProcessor interface {
	Process(ctx context.Context, msg message.SaveFacesCommand) error
}

type EdenSearchMessageProcessor interface {
	Process(ctx context.Context, msg message.SearchProfilesCommand) error
}
