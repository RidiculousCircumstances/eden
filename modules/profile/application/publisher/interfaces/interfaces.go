package interfaces

import (
	"context"
	"eden/modules/profile/infrastructure/eden_gate/interfaces"
)

type EdenGateClient interface {
	SendSearchResult(ctx context.Context, msg interfaces.ProfileSearchCompletedEvent) error
}
