package broker

import (
	"context"
	"eden/shared/broker/interfaces"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
)

type watermillSubscriber struct {
	connection *amqp.ConnectionWrapper
	logger     watermill.LoggerAdapter
}

func NewSubscriber(conn *amqp.ConnectionWrapper, logger watermill.LoggerAdapter) (interfaces.Subscriber, error) {
	return &watermillSubscriber{
		connection: conn,
		logger:     logger,
	}, nil
}

func (s *watermillSubscriber) Subscribe(ctx context.Context, exchangeName, topic string, handler interfaces.MessageHandler) error {
	// Создаем конфигурацию подписчика
	subConfig := amqp.NewDurableQueueConfig("")

	// Настраиваем динамическое имя обменника
	subConfig.Exchange.GenerateName = func(topic string) string {
		return exchangeName
	}

	subConfig.Exchange.Type = "direct"
	subConfig.Exchange.Durable = true

	subConfig.QueueBind = amqp.QueueBindConfig{
		GenerateRoutingKey: func(topic string) string {
			return topic
		},
	}

	// Создаем подписчика
	sub, err := amqp.NewSubscriberWithConnection(subConfig, s.logger, s.connection)
	if err != nil {
		return err
	}

	// Подписываемся на топик
	messages, err := sub.Subscribe(ctx, topic)
	if err != nil {
		return err
	}

	// Обрабатываем сообщения
	go func() {
		for {
			select {
			case <-ctx.Done():
				_ = sub.Close() // Закрываем подписчика при завершении контекста
				return
			case msg, ok := <-messages:
				if !ok {
					s.logger.Error("Message channel closed", nil, watermill.LogFields{"topic": topic})
					return
				}

				// Обрабатываем сообщение
				go func(m *message.Message) {
					ack, err := handler.Handle(ctx, m.Payload)
					if err != nil {
						s.logger.Error("Error processing message", err, watermill.LogFields{
							"message_id": m.UUID,
						})
					}

					if ack {
						m.Ack()
					} else {
						m.Nack()
					}
				}(msg)
			}
		}
	}()

	return nil
}

func (s *watermillSubscriber) Close() error {
	return s.connection.Close() // Закрываем общее подключение
}
