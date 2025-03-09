package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Config структура для хранения конфигурационных параметров
type Config struct {
	DatabaseDSN                         string
	LogPath                             string
	RabbitMQURL                         string
	LogLevel                            string
	EdenExchangeName                    string
	EdenProfileQueueName                string
	EdenIndexedQueueName                string
	EdenSearchQueueName                 string
	EdenGateExchangeName                string
	EdenGateSearchResultQueueName       string
	ReliquariumCommandExchangeName      string
	ReliquariumConfirmationExchangeName string
	ReliquariumConfirmationQueueName    string
	EdenSnapshotControlQueueName        string
	StorageEndpoint                     string
	StorageAccessKeyId                  string
	StorageSecretAccessKey              string
	SnapshotBucketName                  string
	DatabaseUser                        string
	DatabasePassword                    string
	DatabaseName                        string
	DatabaseHost                        string
}

// LoadConfig загружает конфигурацию из .env файла
func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Ошибка загрузки .env файла: %v", err)
	}

	return &Config{
		DatabaseDSN:                         os.Getenv("DATABASE_DSN"),
		LogPath:                             os.Getenv("LOG_PATH"),
		RabbitMQURL:                         os.Getenv("RABBITMQ_URL"),
		LogLevel:                            os.Getenv("LOG_LEVEL"),
		EdenExchangeName:                    os.Getenv("EDEN_EXCHANGE_NAME"),
		EdenProfileQueueName:                os.Getenv("EDEN_PROFILE_QUEUE_NAME"),
		EdenIndexedQueueName:                os.Getenv("EDEN_INDEXED_QUEUE_NAME"),
		EdenSearchQueueName:                 os.Getenv("EDEN_SEARCH_QUEUE_NAME"),
		EdenGateExchangeName:                os.Getenv("EDEN_GATE_EXCHANGE_NAME"),
		EdenGateSearchResultQueueName:       os.Getenv("EDEN_GATE_SEARCH_RESULT_QUEUE_NAME"),
		ReliquariumCommandExchangeName:      os.Getenv("RELIQUARIUM_COMMAND_EXCHANGE_NAME"),
		EdenSnapshotControlQueueName:        os.Getenv("EDEN_SNAPSHOT_CONTROL_QUEUE_NAME"),
		ReliquariumConfirmationExchangeName: os.Getenv("RELIQUARIUM_CONFIRMATION_EXCHANGE_NAME"),
		ReliquariumConfirmationQueueName:    os.Getenv("RELIQUARIUM_CONFIRMATION_QUEUE_NAME"),
		StorageEndpoint:                     os.Getenv("STORAGE_ENDPOINT"),
		StorageAccessKeyId:                  os.Getenv("STORAGE_ACCESS_KEY_ID"),
		StorageSecretAccessKey:              os.Getenv("STORAGE_SECRET_ACCESS_KEY"),
		SnapshotBucketName:                  os.Getenv("SNAPSHOT_BUCKET_NAME"),
		DatabaseUser:                        os.Getenv("DATABASE_USER"),
		DatabasePassword:                    os.Getenv("DATABASE_PASSWORD"),
		DatabaseName:                        os.Getenv("DATABASE_NAME"),
		DatabaseHost:                        os.Getenv("DATABASE_HOST"),
	}, nil
}
