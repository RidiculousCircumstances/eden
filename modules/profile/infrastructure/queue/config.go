package queue

import (
	"eden/config/env"
	"eden/modules/profile/application/consumer"
	consumerIntf "eden/modules/profile/application/consumer/interfaces"
	"eden/shared/broker/interfaces"
)

type HandlerConfig struct {
	QueueName    string
	ExchangeName string
	Handler      interfaces.MessageHandler
}

func RegisterHandlersConfig(
	env *env.Config,
	sfMsgProcessor consumerIntf.StreamForgeMessageProcessor,
	tfMsgProcessor consumerIntf.TraceFaceMessageProcessor,
	searchMessageProcessor consumerIntf.EdenSearchMessageProcessor,
) []HandlerConfig {
	return []HandlerConfig{
		{
			env.EdenProfileQueueName,
			env.EdenExchangeName,
			consumer.NewStreamForgeMessageHandler(sfMsgProcessor),
		},
		{
			env.EdenIndexedQueueName,
			env.EdenExchangeName,
			consumer.NewTraceFaceMessageHandler(tfMsgProcessor),
		},
		{
			env.EdenSearchQueueName,
			env.EdenExchangeName,
			consumer.NewEdenSearchMessageHandler(searchMessageProcessor),
		},
	}
}
