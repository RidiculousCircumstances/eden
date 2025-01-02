package logger

import (
	"eden/config/env"
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

// Ошибки для логирования
var (
	ErrLogDirCreation       = errors.New("failed to create log directory")
	ErrLoggerInitialization = errors.New("failed to initialize logger")
)

// логирование уровней
var logLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
}

// NewZapLogger создает новый zap.Logger
func NewZapLogger(cfg *env.Config) (*zap.Logger, error) {
	logDir := cfg.LogPath

	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return nil, errors.Join(ErrLogDirCreation, err)
	}

	logFile := filepath.Join(logDir, "application.log")

	config := zap.NewProductionConfig()
	config.OutputPaths = []string{
		"stdout",
		logFile,
	}
	config.ErrorOutputPaths = []string{
		"stderr",
		logFile,
	}

	// Конфигурация уровня логирования
	logLevel, ok := logLevelMap[cfg.LogLevel]
	if !ok {
		logLevel = zapcore.InfoLevel
	}

	config.Level = zap.NewAtomicLevelAt(logLevel)

	// Инициализация логгера
	logger, err := config.Build()
	if err != nil {
		return nil, errors.Join(ErrLoggerInitialization, err)
	}

	return logger, nil
}
