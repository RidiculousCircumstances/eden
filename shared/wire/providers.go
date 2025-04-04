package wire

import (
	"eden/config/env"
	consumerIntf "eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/application/publisher"
	pubIntf "eden/modules/profile/application/publisher/interfaces"
	"eden/modules/profile/application/service"
	"eden/modules/profile/application/storage"
	storageIntf "eden/modules/profile/application/storage/interfaces"
	"eden/modules/profile/application/usecase"
	servIntf "eden/modules/profile/application/usecase/interfaces"
	profileRepoIntf "eden/modules/profile/domain/interfaces"
	"eden/modules/profile/infrastructure/appstate"
	"eden/modules/profile/infrastructure/eden_gate"
	"eden/modules/profile/infrastructure/queue"
	"eden/modules/profile/infrastructure/reliquarium"
	profileRepo "eden/modules/profile/infrastructure/repository"
	infraStorageIntf "eden/modules/profile/infrastructure/storage"
	"eden/shared/database"
	lifecycleIntf "eden/shared/lifecycle/interfaces"
	"eden/shared/logger"
	loggerIntf "eden/shared/logger/interfaces"
	"github.com/RidiculousCircumstances/netherway/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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

func ProvideEdenGateClient(broker netherway.MessageBroker) pubIntf.EdenGateClient {
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
	takeSnapshot consumerIntf.ManageSnapshotLifecycle,
) []queue.HandlerConfig {
	return queue.RegisterHandlersConfig(
		cfg,
		logger,
		sfMessageProcessor,
		tfMessageProcessor,
		searchMessageHandler,
		takeSnapshot,
	)
}

func ProvideStorageClient(cfg *env.Config) (storageIntf.StorageClient, error) {
	baseMinioClient, err := minio.New(cfg.StorageEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.StorageAccessKeyId, cfg.StorageSecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	minioClient, err := infraStorageIntf.NewMinioClient(baseMinioClient)

	if err != nil {
		return nil, err
	}

	return minioClient, nil
}

func ProvideStorageService(client storageIntf.StorageClient) servIntf.StorageService {
	return storage.NewService(client)
}

func ProvideTakeSnapshot(logger loggerIntf.Logger, storageService servIntf.StorageService, cfg *env.Config) servIntf.TakeSnapshot {
	return usecase.NewTakeSnapshot(logger, storageService, &usecase.TakeSnapshotConfig{
		SnapshotBucket: cfg.SnapshotBucketName,
		DbUser:         cfg.DatabaseUser,
		DbPassword:     cfg.DatabasePassword,
		DbName:         cfg.DatabaseName,
		DbHost:         cfg.DatabaseHost,
	})
}

func ProvideManageSnapshotLifecycle(
	stateManager consumerIntf.AppStateManager,
	confirmationPublisher consumerIntf.ServiceCommandConfirmationPublisher,
	logger loggerIntf.Logger,
	snapshot servIntf.TakeSnapshot,
) consumerIntf.ManageSnapshotLifecycle {
	return usecase.NewManageSnapshotLifecycle(stateManager, confirmationPublisher, snapshot, logger)
}

func ProvideServiceCommandConfirmationPublisher(client pubIntf.ReliquariumClient) consumerIntf.ServiceCommandConfirmationPublisher {
	return publisher.NewServiceCommandConfirmationPublisher(client)
}

func ProvideReliquariumClient(broker netherway.MessageBroker) pubIntf.ReliquariumClient {
	return reliquarium.NewClient(broker)
}

func ProvideLifecycleHooks(handlerCfgs []queue.HandlerConfig, logger loggerIntf.Logger, broker netherway.MessageBroker) []lifecycleIntf.Hook {
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

func ProvideMessageBroker(cfg *env.Config, logger loggerIntf.Logger) netherway.MessageBroker {
	pubFactory := func(conn netherway.Connection) netherway.Publisher {
		return netherway.NewPublisher(conn, logger)
	}

	subFactory := func(conn netherway.Connection) netherway.Subscriber {
		return netherway.NewSubscriber(conn, logger)
	}

	connFactory := netherway.NewConnFactory(netherway.ConnConfig{
		AmqpURI:                  cfg.RabbitMQURL,
		PublisherChannelPoolSize: 10,
		Exchanges: []netherway.Exchange{
			{
				Name: cfg.EdenGateExchangeName,
				Type: "direct",
			},
			{
				Name: cfg.ReliquariumConfirmationExchangeName,
				Type: "direct",
			},
		},
	})

	return netherway.NewMessageBroker(netherway.Config{
		PublisherFactory:  pubFactory,
		SubscriberFactory: subFactory,
		Logger:            logger,
		ConnFactory:       connFactory,
	})
}

func ProvideAppStateManager(broker netherway.MessageBroker, logger loggerIntf.Logger, env *env.Config) consumerIntf.AppStateManager {
	return statemanager.NewAppStateManager(
		broker,
		logger,
		[]string{
			env.EdenExchangeName + ":" + env.EdenProfileQueueName,
			env.EdenExchangeName + ":" + env.EdenSearchQueueName,
			env.EdenExchangeName + ":" + env.EdenIndexedQueueName,
		},
	)
}
