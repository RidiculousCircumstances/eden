package queue

import (
	"eden/config/env"
	"eden/modules/profile/application/consumer"
	consumerIntf "eden/modules/profile/application/consumer/interfaces"
	"eden/shared/broker/interfaces"
	loggerIntf "eden/shared/logger/interfaces"
)

type HandlerConfig struct {
	QueueName    string
	ExchangeName string
	Handler      interfaces.MessageHandler
}

func RegisterHandlersConfig(
	env *env.Config,
	logger loggerIntf.Logger,
	sfMsgProcessor consumerIntf.SaveProfiles,
	tfMsgProcessor consumerIntf.SaveFaceInfo,
	searchMessageProcessor consumerIntf.SearchProfiles,
	stateManager consumerIntf.AppStateManager,
	confirmationPublisher consumerIntf.ServiceCommandConfirmationPublisher,
) []HandlerConfig {
	return []HandlerConfig{
		{
			env.EdenProfileQueueName,
			env.EdenExchangeName,
			consumer.NewStreamForgeMessageHandler(sfMsgProcessor, logger),
		},
		{
			env.EdenIndexedQueueName,
			env.EdenExchangeName,
			consumer.NewTraceFaceMessageHandler(tfMsgProcessor, logger),
		},
		{
			env.EdenSearchQueueName,
			env.EdenExchangeName,
			consumer.NewEdenSearchMessageHandler(searchMessageProcessor, logger),
		},
		{
			env.EdenSnapshotControlQueueName,
			env.ReliquariumCommandExchangeName,
			consumer.NewReliquariumMessageHandler(stateManager, confirmationPublisher, logger),
		},
	}
}
