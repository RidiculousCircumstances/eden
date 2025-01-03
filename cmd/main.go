package main

import (
	"eden/shared/utils"
	"eden/shared/wire"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Инициализируем все компоненты приложения
	app, err := wire.InitializeApp()
	if err != nil {
		log.Fatalf("Application initialization error: %v", err)
	}

	// Создаем канал для остановки приложения
	stopCh := make(chan interface{})
	ctx := utils.CreateContextWithStopChannel(stopCh)

	// Получаем экземпляр жизненного цикла
	appLifecycle := app.Lifecycle

	if err := appLifecycle.Start(ctx); err != nil {
		app.Logger.Fatal("Failed to start application", zap.Error(err))
	}

	// Ожидаем сигнала завершения работы (например, SIGINT, SIGTERM)
	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGTERM)

	// Ожидаем завершения работы приложения
	select {
	case sig := <-stopSignal:
		app.Logger.Info("Received shutdown signal", zap.String("signal", sig.String()))
	}

	// Завершаем жизненный цикл
	if err := appLifecycle.Shutdown(ctx); err != nil {
		app.Logger.Fatal("Failed to gracefully shutdown application", zap.Error(err))
	}

	// Завершаем работу всех сервисов
	app.Logger.Info("Shutting down application")
	close(stopCh)

	app.Logger.Info("Application stopped gracefully")
}
