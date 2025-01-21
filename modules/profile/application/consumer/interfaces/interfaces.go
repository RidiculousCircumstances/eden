package interfaces

import (
	"context"
	"eden/modules/profile/infrastructure/eden_gate/messages"
	"eden/modules/profile/infrastructure/queue/message"
)

type EdenGateSearchResultPublisher interface {
	Publish(ctx context.Context, msg messages.ProfileSearchCompletedEvent) error
}

type SaveProfiles interface {
	Process(ctx context.Context, msg message.SaveProfileCommand) error
}

type SaveFaceInfo interface {
	Process(ctx context.Context, msg message.SaveFacesCommand) error
}

type SearchProfiles interface {
	Process(ctx context.Context, msg message.SearchProfilesCommand) error
}
