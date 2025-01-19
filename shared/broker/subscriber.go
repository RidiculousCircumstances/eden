package broker

import (
	"context"
	"eden/shared/broker/interfaces"
	loggerIntf "eden/shared/logger/interfaces"
	"go.uber.org/zap"
	"sync"
)

// Логирующие сообщения
var (
	logErrorProcessingMessage  = "Error processing message"
	logMessageNotAcknowledged  = "Message not acknowledged"
	logErrorAcknowledgeMessage = "Error acknowledge message"
)

type subscriber struct {
	connection interfaces.Connection
	logger     loggerIntf.Logger
}

func NewSubscriber(conn interfaces.Connection, logger loggerIntf.Logger) interfaces.Subscriber {
	return &subscriber{
		connection: conn,
		logger:     logger,
	}
}

func (s *subscriber) Subscribe(ctx context.Context, exchangeName, topic string, handler interfaces.MessageHandler) error {
	messages, consumerErr := s.connection.Consume(ctx, exchangeName, topic)
	if consumerErr != nil {
		s.logger.Error("Error creating consumer", zap.Error(consumerErr))
		return consumerErr
	}

	var wg sync.WaitGroup
	defer wg.Wait()

	for msg := range messages {
		wg.Add(1)
		go func(msg interfaces.UnitOfWork) {
			defer wg.Done()

			ack, handlerErr := handler.Handle(ctx, msg.GetPayload())
			if handlerErr != nil {
				s.logger.Error(logErrorProcessingMessage, zap.Error(handlerErr))
				_ = msg.Nack(false)
				return
			}

			if ack {
				if err := msg.Ack(); err != nil {
					s.logger.Error(logErrorAcknowledgeMessage, zap.Error(err))
				}
			} else {
				s.logger.Warn(logMessageNotAcknowledged)
				_ = msg.Nack(true)
			}
		}(msg)
	}

	// Обработка завершения контекста
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		// Если контекст не завершен, продолжаем
	}
	return nil
}

func (s *subscriber) Close() error {
	return s.connection.Close()
}
