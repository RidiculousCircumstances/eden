package queue

import (
	"context"
	"eden/shared/broker/interfaces"
	loggerIntf "eden/shared/logger/interfaces"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

var (
	ErrConsumerShutdown = errors.New("consumer failed to shut down")
)

// ConsumerHook - Хук для управления консьюмерами RabbitMQ
type ConsumerHook struct {
	MessageHandlers []interfaces.MessageHandler
	Logger          loggerIntf.Logger
	Broker          interfaces.MessageBroker
}

// Setup выполняет инициализацию хуков (пока без реализации)
func (c *ConsumerHook) Setup(ctx context.Context) error {
	return nil
}

// Start запускает консьюмеров
func (c *ConsumerHook) Start(ctx context.Context) error {
	// Для каждого обработчика создаем подписку на топик
	for _, handler := range c.MessageHandlers {
		topic := fmt.Sprintf("topic_for_handler_%T", handler) // Пример генерации топика
		if err := c.Broker.Subscribe(ctx, "exchangeName", topic, handler); err != nil {
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
func NewConsumerHook(handlers []interfaces.MessageHandler, logger loggerIntf.Logger, broker interfaces.MessageBroker) *ConsumerHook {
	return &ConsumerHook{
		MessageHandlers: handlers,
		Logger:          logger,
		Broker:          broker,
	}
}
