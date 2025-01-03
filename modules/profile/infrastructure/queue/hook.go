package queue

import (
	"context"
	"eden/shared/broker/interfaces"
	loggerIntf "eden/shared/logger/interfaces"
	"errors"
	"go.uber.org/zap"
)

var (
	ErrConsumerShutdown = errors.New("consumer failed to shut down")
)

// ConsumerHook - Хук для управления консьюмерами RabbitMQ
type ConsumerHook struct {
	HandlerConfigs []HandlerConfig
	Logger         loggerIntf.Logger
	Broker         interfaces.MessageBroker
}

// Setup выполняет инициализацию хуков (пока без реализации)
func (c *ConsumerHook) Setup(ctx context.Context) error {
	return nil
}

// Start запускает консьюмеров
func (c *ConsumerHook) Start(ctx context.Context) error {
	for _, cfg := range c.HandlerConfigs {
		if err := c.Broker.Subscribe(ctx, cfg.ExchangeName, cfg.QueueName, cfg.Handler); err != nil {
			c.Logger.Error("Failed to subscribe", zap.Error(err))
			return err
		}
	}

	return nil
}

// Shutdown корректно завершает работу потребителей
func (c *ConsumerHook) Shutdown(ctx context.Context) error {
	// Закрываем подписчика
	if err := c.Broker.Close(); err != nil {
		c.Logger.Error("Failed to shutdown subscriber", zap.Error(err))
		return ErrConsumerShutdown
	}

	return nil
}

// NewConsumerHook конструирует ConsumerHook
func NewConsumerHook(handlerCfgs []HandlerConfig, logger loggerIntf.Logger, broker interfaces.MessageBroker) *ConsumerHook {
	return &ConsumerHook{
		HandlerConfigs: handlerCfgs,
		Logger:         logger,
		Broker:         broker,
	}
}
