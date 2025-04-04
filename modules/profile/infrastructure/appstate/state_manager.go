package statemanager

import (
	"context"
	loggerIntf "eden/shared/logger/interfaces"
	"github.com/RidiculousCircumstances/netherway/v2"
	"go.uber.org/zap"
)

type AppStateManager struct {
	broker   netherway.MessageBroker
	logger   loggerIntf.Logger
	isPaused bool
	stopList []string
	appCtx   context.Context
}

func NewAppStateManager(
	broker netherway.MessageBroker,
	logger loggerIntf.Logger,
	stopList []string,
) *AppStateManager {
	// Здесь создаём глобальный контекст, который живёт, пока всё приложение работает
	appCtx := context.Background()

	return &AppStateManager{
		broker:   broker,
		logger:   logger,
		stopList: stopList,
		appCtx:   appCtx,
	}
}

func (m *AppStateManager) Pause() {
	if m.isPaused {
		return
	}
	m.logger.Info("Pausing all message consumption")
	m.broker.Pause(m.stopList)
	m.isPaused = true
}

func (m *AppStateManager) Resume() {
	if !m.isPaused {
		return
	}
	m.logger.Info("Resuming all message consumption")

	// Используем appCtx, который не отменяется сразу
	if err := m.broker.Resume(m.appCtx, []string{}); err != nil {
		m.logger.Error("Error resuming subscriptions", zap.Error(err))
	}
	m.isPaused = false
}
