package interfaces

import (
	"context"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
)

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
	GetConnection() (*amqp.ConnectionWrapper, error)
}
