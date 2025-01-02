package broker

import (
	"context"
	"eden/shared/broker/interfaces"
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"sync"
)

type PubFactory func(conn *amqp.ConnectionWrapper) (interfaces.Publisher, error)
type SubFactory func(conn *amqp.ConnectionWrapper) (interfaces.Subscriber, error)

type messageBroker struct {
	publisherFactory  PubFactory
	subscriberFactory SubFactory
	conn              *amqp.ConnectionWrapper

	publisher     interfaces.Publisher
	subscriber    interfaces.Subscriber
	publisherMux  sync.Mutex
	subscriberMux sync.Mutex

	logger watermill.LoggerAdapter
}

// Config содержит фабричные функции и логгер для инициализации брокера
type Config struct {
	PublisherFactory  PubFactory
	SubscriberFactory SubFactory
	ConnFactory       interfaces.ConnFactory
	Logger            watermill.LoggerAdapter
}

func NewMessageBroker(cfg Config) interfaces.MessageBroker {
	conn, err := cfg.ConnFactory.GetConnection()
	if err != nil {
		panic(err)
	}

	return &messageBroker{
		publisherFactory:  cfg.PublisherFactory,
		subscriberFactory: cfg.SubscriberFactory,
		logger:            cfg.Logger,
		conn:              conn,
	}
}

// lazyInitPublisher инициализирует publisher при первом вызове
func (mb *messageBroker) lazyInitPublisher() error {
	mb.publisherMux.Lock()
	defer mb.publisherMux.Unlock()

	if mb.publisher != nil {
		return nil // Уже инициализирован
	}

	pub, err := mb.publisherFactory(mb.conn)
	if err != nil {
		return err
	}

	mb.publisher = pub
	return nil
}

// lazyInitSubscriber инициализирует subscriber при первом вызове
func (mb *messageBroker) lazyInitSubscriber() error {
	mb.subscriberMux.Lock()
	defer mb.subscriberMux.Unlock()

	if mb.subscriber != nil {
		return nil // Уже инициализирован
	}

	sub, err := mb.subscriberFactory(mb.conn)
	if err != nil {
		return err
	}

	mb.subscriber = sub
	return nil
}

// Publish отправляет сообщение в указанный топик
func (mb *messageBroker) Publish(ctx context.Context, exchangeName, topic string, data interface{}) error {
	if err := mb.lazyInitPublisher(); err != nil {
		return err
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	msg := message.NewMessage(watermill.NewUUID(), payload)
	return mb.publisher.Publish(ctx, exchangeName, topic, msg)
}

// Subscribe подписывается на топик и обрабатывает сообщения с помощью переданного обработчика
func (mb *messageBroker) Subscribe(ctx context.Context, exchangeName, topic string, handler interfaces.MessageHandler) error {
	if err := mb.lazyInitSubscriber(); err != nil {
		return err
	}

	return mb.subscriber.Subscribe(ctx, exchangeName, topic, handler)
}

// Close закрывает publisher и subscriber, если они были инициализированы
func (mb *messageBroker) Close() error {
	var pubErr, subErr error

	if mb.publisher != nil {
		pubErr = mb.publisher.Close()
	}

	if mb.subscriber != nil {
		subErr = mb.subscriber.Close()
	}

	if pubErr != nil {
		return pubErr
	}
	return subErr
}
