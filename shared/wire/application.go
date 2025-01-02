package wire

import (
	"eden/shared/lifecycle"
	logIntf "eden/shared/logger/interfaces"
)

// Application включает все основные компоненты
type Application struct {
	Logger    logIntf.Logger
	Lifecycle *lifecycle.AppLifecycle
}

// ProvideApplication инициализирует приложение с компонентами
func ProvideApplication(logger logIntf.Logger, lifecycle *lifecycle.AppLifecycle) Application {
	return Application{
		Logger:    logger,
		Lifecycle: lifecycle,
	}
}
