package interfaces

import (
	"context"
	"eden/modules/profile/infrastructure/eden_gate/messages"
)

type EdenGateClient interface {
	SendSearchResult(ctx context.Context, msg messages.ProfileSearchCompletedEvent) error
}
