package interfaces

import (
	"context"
)

type Connection interface {
	Publish(ctx context.Context, exchangeName, routingKey string, message []byte) error
	Consume(ctx context.Context, exchangeName, topic string) (<-chan UnitOfWork, error)
	Close() error
}

type MessageBroker interface {
	Publisher
	Subscriber
}

type Publisher interface {
	Publish(ctx context.Context, exchangeName, topic string, data interface{}) error
	Close() error
}

type Subscriber interface {
	Subscribe(ctx context.Context, exchangeName, topic string, handler MessageHandler) error
	Close() error
}

type MessageHandler interface {
	Handle(ctx context.Context, msg interface{}) (bool, error)
}

type ConnFactory interface {
	GetConnection() (Connection, error)
}

// Serializer интерфейс для сериализации и десериализации
type Serializer interface {
	Serialize(data interface{}) ([]byte, error)
	Deserialize(data []byte, v interface{}) error
}

type UnitOfWork interface {
	Nack(requeue bool) error
	Ack() error
	GetPayload() []byte
}
