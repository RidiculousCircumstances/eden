package broker

import (
	"context"
	"eden/shared/broker/interfaces"
	loggerIntf "eden/shared/logger/interfaces"
	"sync"
)

type PubFactory func(conn interfaces.Connection) interfaces.Publisher
type SubFactory func(conn interfaces.Connection) interfaces.Subscriber

type messageBroker struct {
	publisherFactory  PubFactory
	subscriberFactory SubFactory
	conn              interfaces.Connection
	publisher         interfaces.Publisher
	subscriber        interfaces.Subscriber
	publisherMux      sync.Mutex
	subscriberMux     sync.Mutex
	logger            loggerIntf.Logger
}

type Config struct {
	PublisherFactory  PubFactory
	SubscriberFactory SubFactory
	ConnFactory       interfaces.ConnFactory
	Logger            loggerIntf.Logger
	Serializer        interfaces.Serializer // Добавляем сериализатор в конфиг
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

	mb.publisher = mb.publisherFactory(mb.conn)
	return nil
}

// lazyInitSubscriber инициализирует subscriber при первом вызове
func (mb *messageBroker) lazyInitSubscriber() error {
	mb.subscriberMux.Lock()
	defer mb.subscriberMux.Unlock()

	if mb.subscriber != nil {
		return nil // Уже инициализирован
	}

	mb.subscriber = mb.subscriberFactory(mb.conn)
	return nil
}

// Publish отправляет сообщение в указанный топик
func (mb *messageBroker) Publish(ctx context.Context, exchangeName, topic string, data interface{}) error {
	if err := mb.lazyInitPublisher(); err != nil {
		return err
	}

	return mb.publisher.Publish(ctx, exchangeName, topic, data)
}

// Subscribe подписывается на топик и обрабатывает сообщения с помощью переданного обработчика
func (mb *messageBroker) Subscribe(ctx context.Context, exchangeName, topic string, handler interfaces.MessageHandler) error {
	// Инициализируем подписчика (ленивый метод)
	if err := mb.lazyInitSubscriber(); err != nil {
		return err
	}

	// Просто вызываем метод Subscribe из уже инициализированного подписчика
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
