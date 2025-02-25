//go:build wireinject

package wire

import (
	"eden/config/env"
	"eden/shared/lifecycle"
	"github.com/google/wire"
)

var LifecycleSet = wire.NewSet(
	lifecycle.NewAppLifecycle,
	ProvideLifecycleHooks,
)

var InfraSet = wire.NewSet(
	ProvideDatabase,
	ProvidePhotoRepository,
	ProvideProfileRepository,
	ProvideFaceRepository,
	ProvideAppStateManager,
	ProvideLogger,
	ProvideApplication,
	ProvideMessageBroker,
	ProvideEdenGateClient,
	ProvideHandlerConfigs,
	ProvideReliquariumClient,
	env.LoadConfig,
)

var ApplicationSet = wire.NewSet(
	ProvideStreamForgeMessageProcessor,
	ProvideTraceFaceMessageProcessor,
	ProvidePhotoService,
	ProvideProfileService,
	ProvideFaceService,
	ProvideEdenSearchMessageProcessor,
	ProvideEdenGateSearchResultPublisher,
	ProvideServiceCommandConfirmationPublisher,
)

func InitializeApp() (Application, error) {
	wire.Build(
		InfraSet,
		ApplicationSet,
		LifecycleSet,
	)

	return Application{}, nil
}
