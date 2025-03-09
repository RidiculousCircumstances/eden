package interfaces

import (
	"context"
	edenGateMsg "eden/modules/profile/infrastructure/eden_gate/messages"
	"eden/modules/profile/infrastructure/queue/message"
	reliquariumMsg "eden/modules/profile/infrastructure/reliquarium/messages"
)

type EdenGateSearchResultPublisher interface {
	Publish(ctx context.Context, msg edenGateMsg.ProfileSearchCompletedEvent) error
}

type ServiceCommandConfirmationPublisher interface {
	Publish(ctx context.Context, msg *reliquariumMsg.CommandConfirmationEvent) error
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

type ManageSnapshotLifecycle interface {
	Process(ctx context.Context, msg *message.ServiceControlCommand) error
}

type AppStateManager interface {
	Pause()
	Resume()
}
