package broker

import (
	"context"
	"eden/shared/broker/interfaces"
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
)

type watermillPublisher struct {
	connection *amqp.ConnectionWrapper // Общий объект подключения
	logger     watermill.LoggerAdapter // Логгер
}

func NewPublisher(conn *amqp.ConnectionWrapper, logger watermill.LoggerAdapter) (interfaces.Publisher, error) {
	return &watermillPublisher{
		connection: conn,
		logger:     logger,
	}, nil
}

func (p *watermillPublisher) Publish(ctx context.Context, exchangeName, topic string, data interface{}) error {
	pubConfig := amqp.NewDurableQueueConfig("")
	pubConfig.Exchange.GenerateName = func(topic string) string {
		return exchangeName
	}
	pubConfig.Exchange.Type = "direct"
	pubConfig.Exchange.Durable = true
	pubConfig.QueueBind = amqp.QueueBindConfig{
		GenerateRoutingKey: func(topic string) string {
			return topic
		},
	}

	pub, err := amqp.NewPublisherWithConnection(pubConfig, p.logger, p.connection)
	if err != nil {
		return err
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	msg := message.NewMessage(watermill.NewUUID(), payload)
	return pub.Publish(topic, msg)
}

func (p *watermillPublisher) Close() error {
	return p.connection.Close() // Закрываем соединение
}
