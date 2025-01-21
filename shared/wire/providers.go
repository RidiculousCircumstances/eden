package wire

import (
	"eden/config/env"
	consumerIntf "eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/application/publisher"
	pubIntf "eden/modules/profile/application/publisher/interfaces"
	"eden/modules/profile/application/service"
	"eden/modules/profile/application/usecase"
	servIntf "eden/modules/profile/application/usecase/interfaces"
	profileRepoIntf "eden/modules/profile/domain/interfaces"
	"eden/modules/profile/infrastructure/eden_gate"
	"eden/modules/profile/infrastructure/queue"
	profileRepo "eden/modules/profile/infrastructure/repository"
	brokerLib "eden/shared/broker"
	brokerLibAmqp "eden/shared/broker/amqp"
	brokerIntf "eden/shared/broker/interfaces"
	"eden/shared/database"
	lifecycleIntf "eden/shared/lifecycle/interfaces"
	"eden/shared/logger"
	loggerIntf "eden/shared/logger/interfaces"
	"gorm.io/gorm"
)

func ProvideProfileService(repo profileRepoIntf.ProfileRepository) servIntf.ProfileService {
	return service.NewProfileService(repo)
}

func ProvidePhotoService(repo profileRepoIntf.PhotoRepository) servIntf.PhotoService {
	return service.NewPhotoService(repo)
}

func ProvideFaceService(repo profileRepoIntf.FaceRepository) servIntf.FaceService {
	return service.NewFaceService(repo)
}

func ProvideStreamForgeMessageProcessor(profileSrv servIntf.ProfileService, photoService servIntf.PhotoService) consumerIntf.SaveProfiles {
	return usecase.NewSaveProfiles(profileSrv, photoService)
}

func ProvideTraceFaceMessageProcessor(faceService servIntf.FaceService, photoService servIntf.PhotoService) consumerIntf.SaveFaceInfo {
	return usecase.NewSaveFaceInfo(faceService, photoService)
}

func ProvideEdenSearchMessageProcessor(
	photoService servIntf.PhotoService,
	publisher consumerIntf.EdenGateSearchResultPublisher,
	logger loggerIntf.Logger,
) consumerIntf.SearchProfiles {
	return usecase.NewSearchProfiles(photoService, publisher, logger)
}

func ProvideEdenGateClient(broker brokerIntf.MessageBroker) pubIntf.EdenGateClient {
	return eden_gate.NewClient(broker)
}

func ProvideEdenGateSearchResultPublisher(client pubIntf.EdenGateClient, logger loggerIntf.Logger) consumerIntf.EdenGateSearchResultPublisher {
	return publisher.NewEdenGateSearchResultPublisher(client, logger)
}

func ProvideHandlerConfigs(
	cfg *env.Config,
	logger loggerIntf.Logger,
	sfMessageProcessor consumerIntf.SaveProfiles,
	tfMessageProcessor consumerIntf.SaveFaceInfo,
	searchMessageHandler consumerIntf.SearchProfiles,
) []queue.HandlerConfig {
	return queue.RegisterHandlersConfig(cfg, logger, sfMessageProcessor, tfMessageProcessor, searchMessageHandler)
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

func ProvideFaceRepository(db *gorm.DB) profileRepoIntf.FaceRepository {
	return profileRepo.NewFaceRepository(db)
}

func ProvideMessageBroker(cfg *env.Config, logger loggerIntf.Logger) brokerIntf.MessageBroker {
	pubFactory := func(conn brokerIntf.Connection) brokerIntf.Publisher {
		return brokerLib.NewPublisher(conn, logger)
	}

	subFactory := func(conn brokerIntf.Connection) brokerIntf.Subscriber {
		return brokerLib.NewSubscriber(conn, logger)
	}

	connFactory := brokerLibAmqp.NewConnFactory(brokerLibAmqp.ConnConfig{
		AmqpURI:                   cfg.RabbitMQURL,
		PublisherChannelPoolSize:  10,
		SubscriberChannelPoolSize: 10,
	})

	return brokerLib.NewMessageBroker(brokerLib.Config{
		PublisherFactory:  pubFactory,
		SubscriberFactory: subFactory,
		Logger:            logger,
		ConnFactory:       connFactory,
	})
}
