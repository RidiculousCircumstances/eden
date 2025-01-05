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

func BuildHandlerConfigs(
	env *env.Config,
	sfMsgProcessor consumerIntf.StreamForgeMessageProcessor,
	tfMsgProcessor consumerIntf.TraceFaceMessageProcessor,
) []HandlerConfig {
	return []HandlerConfig{
		{
			env.EdenQueueName,
			env.EdenExchangeName,
			consumer.NewStreamForgeMessageHandler(sfMsgProcessor),
		},
		{
			env.TraceFaceQueueName,
			env.EdenExchangeName,
			consumer.NewTraceFaceMessageHandler(tfMsgProcessor),
		},
	}
}
