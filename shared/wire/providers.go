package wire

import (
	"eden/config/env"
	"eden/modules/profile/application/consumer"
	profileRepoIntf "eden/modules/profile/domain/interfaces"
	"eden/modules/profile/infrastructure/queue"
	profileRepo "eden/modules/profile/infrastructure/repository"
	broker2 "eden/shared/broker"
	brokerIntf "eden/shared/broker/interfaces"
	"eden/shared/database"
	lifecycleIntf "eden/shared/lifecycle/interfaces"
	"eden/shared/logger"
	loggerIntf "eden/shared/logger/interfaces"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"gorm.io/gorm"
)

func ProvideLifecycleHooks(handlers []brokerIntf.MessageHandler, logger loggerIntf.Logger, broker brokerIntf.MessageBroker) []lifecycleIntf.Hook {
	return []lifecycleIntf.Hook{
		queue.NewConsumerHook(handlers, logger, broker),
	}
}

func ProvideLogger(config *env.Config) (loggerIntf.Logger, error) {
	return logger.NewZapLogger(config)
}

func ProvideDatabase(cfg *env.Config) *gorm.DB {
	return database.InitGormDB(cfg.DatabaseDSN)
}

func ProvidePhotoRepository(db *gorm.DB) profileRepoIntf.PhotoRepository {
	return profileRepo.NewPhotoRepository(db)
}

func ProvideProfileRepository(db *gorm.DB) profileRepoIntf.ProfileRepository {
	return profileRepo.NewProfileRepository(db)
}

func ProvideMessageBroker(cfg *env.Config, baseLogger loggerIntf.Logger) brokerIntf.MessageBroker {
	logAdapter := logger.NewZapLoggerAdapter(baseLogger)

	pubFactory := func(conn *amqp.ConnectionWrapper) (brokerIntf.Publisher, error) {
		return broker2.NewPublisher(conn, logAdapter)
	}

	subFactory := func(conn *amqp.ConnectionWrapper) (brokerIntf.Subscriber, error) {
		return broker2.NewSubscriber(conn, logAdapter)
	}

	connFactory := broker2.NewConnFactory(amqp.ConnectionConfig{
		AmqpURI: cfg.RabbitMQURL,
	}, logAdapter)

	return broker2.NewMessageBroker(broker2.Config{
		PublisherFactory:  pubFactory,
		SubscriberFactory: subFactory,
		Logger:            logAdapter,
		ConnFactory:       connFactory,
	})
}

func ProvideMessageHandlers() []brokerIntf.MessageHandler {
	return []brokerIntf.MessageHandler{
		consumer.NewStreamForgeMessageHandler(),
	}
}
