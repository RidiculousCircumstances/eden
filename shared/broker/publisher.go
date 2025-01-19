package broker

import (
	"context"
	"eden/shared/broker/interfaces"
	loggerIntf "eden/shared/logger/interfaces"
	"encoding/json"
	"go.uber.org/zap"
)

type publisher struct {
	connection interfaces.Connection // Используем интерфейс соединения
	logger     loggerIntf.Logger     // Логгер
}

// NewPublisher создает нового паблишера, используя интерфейс соединения.
func NewPublisher(conn interfaces.Connection, logger loggerIntf.Logger) interfaces.Publisher {
	return &publisher{
		connection: conn,
		logger:     logger,
	}
}

func (p *publisher) Publish(ctx context.Context, exchangeName, topic string, data interface{}) error {
	// Сериализация данных
	payload, err := json.Marshal(data)
	if err != nil {
		p.logger.Error("Failed to serialize data", zap.Error(err))
		return err
	}

	// Публикация сообщения через интерфейс соединения
	err = p.connection.Publish(ctx, exchangeName, topic, payload)
	if err != nil {
		p.logger.Error("Failed to publish message", zap.Error(err))
	}
	return err
}

func (p *publisher) Close() error {
	// Закрываем соединение через интерфейс
	return p.connection.Close()
}
