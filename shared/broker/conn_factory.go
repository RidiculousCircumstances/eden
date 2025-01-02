package broker

import (
	"eden/shared/broker/interfaces"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
)

type connFactory struct {
	connCfg amqp.ConnectionConfig
	logger  watermill.LoggerAdapter
}

func NewConnFactory(connCfg amqp.ConnectionConfig, logger watermill.LoggerAdapter) interfaces.ConnFactory {
	return &connFactory{
		connCfg: connCfg,
		logger:  logger,
	}
}

func (cf *connFactory) GetConnection() (*amqp.ConnectionWrapper, error) {
	conn, err := amqp.NewConnection(cf.connCfg, cf.logger)
	if err != nil {
		return &amqp.ConnectionWrapper{}, err
	}

	return conn, nil
}
