package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Config структура для хранения конфигурационных параметров
type Config struct {
	DatabaseDSN             string
	LogPath                 string
	RabbitMQURL             string
	LogLevel                string
	StreamForgeQueueName    string
	StreamForgeExchangeName string
}

// LoadConfig загружает конфигурацию из .env файла
func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Ошибка загрузки .env файла: %v", err)
	}

	return &Config{
		DatabaseDSN:             os.Getenv("DATABASE_DSN"),
		LogPath:                 os.Getenv("LOG_PATH"),
		RabbitMQURL:             os.Getenv("RABBITMQ_URL"),
		LogLevel:                os.Getenv("LOG_LEVEL"),
		StreamForgeQueueName:    os.Getenv("STREAM_FORGE_QUEUE_NAME"),
		StreamForgeExchangeName: os.Getenv("STREAM_FORGE_EXCHANGE_NAME"),
	}, nil
}
