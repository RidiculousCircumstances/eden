package queue

import (
	"eden/config/env"
	"eden/modules/profile/application/consumer"
	consumerIntf "eden/modules/profile/application/consumer/interfaces"
	loggerIntf "eden/shared/logger/interfaces"
	"github.com/RidiculousCircumstances/netherway/v2"
)

type HandlerConfig struct {
	QueueName    string
	ExchangeName string
	Handler      netherway.MessageHandler
}

func RegisterHandlersConfig(
	env *env.Config,
	logger loggerIntf.Logger,
	sfMsgProcessor consumerIntf.SaveProfiles,
	tfMsgProcessor consumerIntf.SaveFaceInfo,
	searchMessageProcessor consumerIntf.SearchProfiles,
	takeSnapshot consumerIntf.ManageSnapshotLifecycle,
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
			consumer.NewReliquariumMessageHandler(logger, takeSnapshot),
		},
	}
}
