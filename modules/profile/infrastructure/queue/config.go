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

func BuildHandlerConfigs(env *env.Config, msgProcessor consumerIntf.StreamForgeMessageProcessor) []HandlerConfig {
	return []HandlerConfig{
		{
			env.EdenQueueName,
			env.EdenExchangeName,
			consumer.NewStreamForgeMessageHandler(msgProcessor),
		},
	}
}
