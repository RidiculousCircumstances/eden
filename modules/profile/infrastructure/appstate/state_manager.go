package statemanager

import (
	"context"
	"eden/shared/broker/interfaces"
	loggerIntf "eden/shared/logger/interfaces"
	"go.uber.org/zap"
	"time"
)

type AppStateManager struct {
	broker   interfaces.MessageBroker
	logger   loggerIntf.Logger
	isPaused bool
}

func NewAppStateManager(
	broker interfaces.MessageBroker,
	logger loggerIntf.Logger,
) *AppStateManager {
	return &AppStateManager{
		broker: broker,
		logger: logger,
	}
}

func (m *AppStateManager) Pause() {
	if m.isPaused {
		return
	}
	m.logger.Info("Pausing all message consumption")
	m.broker.Pause()
	m.isPaused = true
}

func (m *AppStateManager) Resume() {
	if !m.isPaused {
		return
	}
	m.logger.Info("Resuming all message consumption")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := m.broker.Resume(ctx); err != nil {
		m.logger.Error("Error resuming subscriptions", zap.Error(err))
	}
	m.isPaused = false
}
