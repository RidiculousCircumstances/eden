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

var DatabaseSet = wire.NewSet(
	ProvideDatabase,
	ProvidePhotoRepository,
	ProvideProfileRepository,
	ProvideFaceRepository,
)

var ApplicationSet = wire.NewSet(
	env.LoadConfig,
	ProvideLogger,
	ProvideApplication,
	ProvideStreamForgeMessageProcessor,
	ProvideTraceFaceMessageProcessor,
	ProvideMessageBroker,
	ProvideHandlerConfigs,
	ProvidePhotoService,
	ProvideProfileService,
	ProvideFaceService,
	ProvideEdenSearchMessageProcessor,
	ProvideEdenGateClient,
	ProvideEdenGateSearchResultPublisher,
)

func InitializeApp() (Application, error) {
	wire.Build(
		DatabaseSet,
		ApplicationSet,
		LifecycleSet,
	)

	return Application{}, nil
}
