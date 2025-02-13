package interfaces

import (
	"context"
	edenGateMsg "eden/modules/profile/infrastructure/eden_gate/messages"
	reliquariumMsg "eden/modules/profile/infrastructure/reliquarium/messages"
)

type EdenGateClient interface {
	SendSearchResult(ctx context.Context, msg edenGateMsg.ProfileSearchCompletedEvent) error
}

type ReliquariumClient interface {
	SendPauseConfirmation(ctx context.Context, msg reliquariumMsg.PauseConfirmationEvent) error
}
