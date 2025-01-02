package lifecycle

import (
	"context"
	"eden/shared/lifecycle/interfaces"
	loggerIntf "eden/shared/logger/interfaces"
	"errors"
	"go.uber.org/zap"
)

// Ошибки для управления жизненным циклом
var (
	ErrAppAlreadyStarted = errors.New("application already started")
	ErrAppNotRunning     = errors.New("application not running")
	ErrStarting          = errors.New("error during starting")
	ErrSetup             = errors.New("error during setup")
)

// AppLifecycle управляет жизненным циклом приложения.
type AppLifecycle struct {
	Logger     loggerIntf.Logger
	Hooks      []interfaces.Hook
	Started    bool
	Terminated bool
}

// NewAppLifecycle создает и инициализирует объект жизненного цикла приложения.
func NewAppLifecycle(log loggerIntf.Logger, hooks []interfaces.Hook) *AppLifecycle {
	return &AppLifecycle{
		Logger: log,
		Hooks:  hooks, // Получаем хуки через DI
	}
}

// RegisterHook добавляет новый хук в жизненный цикл.
func (l *AppLifecycle) RegisterHook(hook interfaces.Hook) {
	l.Hooks = append(l.Hooks, hook)
}

// Start запускает приложение, вызывая все хуки в нужном порядке.
func (l *AppLifecycle) Start(ctx context.Context) error {
	if l.Started {
		return ErrAppAlreadyStarted
	}

	l.Logger.Info("Starting application")

	for _, hook := range l.Hooks {
		if err := hook.Setup(ctx); err != nil {
			return errors.Join(err, ErrSetup)
		}

		if err := hook.Start(ctx); err != nil {
			return errors.Join(err, ErrStarting)
		}
	}

	l.Started = true
	l.Logger.Info("Application started successfully")
	return nil
}

// Shutdown завершает работу приложения, вызывая хуки в нужном порядке.
func (l *AppLifecycle) Shutdown(ctx context.Context) error {
	if !l.Started || l.Terminated {
		return ErrAppNotRunning
	}

	l.Logger.Info("Shutting down application")

	var shutdownErrors []error
	for i := len(l.Hooks) - 1; i >= 0; i-- {
		hook := l.Hooks[i]
		if err := hook.Shutdown(ctx); err != nil {
			l.Logger.Error("Error during Stop", zap.Error(err))
			shutdownErrors = append(shutdownErrors, err)
		}
	}

	if len(shutdownErrors) > 0 {
		return errors.Join(shutdownErrors...)
	}

	l.Terminated = true
	l.Logger.Info("Application shut down successfully")
	return nil
}
