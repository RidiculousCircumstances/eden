package wire

import (
	"eden/config/env"
	consumerIntf "eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/application/service"
	profileRepoIntf "eden/modules/profile/domain/interfaces"
	"eden/modules/profile/infrastructure/queue"
	profileRepo "eden/modules/profile/infrastructure/repository"
	"eden/shared/broker"
	brokerIntf "eden/shared/broker/interfaces"
	"eden/shared/database"
	lifecycleIntf "eden/shared/lifecycle/interfaces"
	"eden/shared/logger"
	loggerIntf "eden/shared/logger/interfaces"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"gorm.io/gorm"
)

func ProvideProfileService(repo profileRepoIntf.ProfileRepository) consumerIntf.ProfileService {
	return service.NewProfileService(repo)
}

func ProvidePhotoService(repo profileRepoIntf.PhotoRepository) consumerIntf.PhotoService {
	return service.NewPhotoService(repo)
}

func ProvideStreamForgeMessageProcessor(profileSrv consumerIntf.ProfileService, photoService consumerIntf.PhotoService) consumerIntf.StreamForgeMessageProcessor {
	return service.NewStreamForgeMessageProcessor(profileSrv, photoService)
}

func ProvideHandlerConfigs(cfg *env.Config, sfMessageProcessor consumerIntf.StreamForgeMessageProcessor) []queue.HandlerConfig {
	return queue.BuildHandlerConfigs(cfg, sfMessageProcessor)
}

func ProvideLifecycleHooks(handlerCfgs []queue.HandlerConfig, logger loggerIntf.Logger, broker brokerIntf.MessageBroker) []lifecycleIntf.Hook {
	return []lifecycleIntf.Hook{
		queue.NewConsumerHook(handlerCfgs, logger, broker),
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
		return broker.NewPublisher(conn, logAdapter)
	}

	subFactory := func(conn *amqp.ConnectionWrapper) (brokerIntf.Subscriber, error) {
		return broker.NewSubscriber(conn, logAdapter)
	}

	connFactory := broker.NewConnFactory(amqp.ConnectionConfig{
		AmqpURI: cfg.RabbitMQURL,
	}, logAdapter)

	return broker.NewMessageBroker(broker.Config{
		PublisherFactory:  pubFactory,
		SubscriberFactory: subFactory,
		Logger:            logAdapter,
		ConnFactory:       connFactory,
	})
}
