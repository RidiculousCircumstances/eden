package consumer

import (
	"context"
	"eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/infrastructure/queue/message"
	loggerIntf "eden/shared/logger/interfaces"
	"encoding/json"
	"github.com/RidiculousCircumstances/netherway/v2"
)

type ReliquariumMessageHandler struct {
	stateManager          interfaces.AppStateManager
	confirmationPublisher interfaces.ServiceCommandConfirmationPublisher
	logger                loggerIntf.Logger
	manageSnapshot        interfaces.ManageSnapshotLifecycle
}

func NewReliquariumMessageHandler(
	logger loggerIntf.Logger,
	manageSnapshot interfaces.ManageSnapshotLifecycle,
) netherway.MessageHandler {
	return &ReliquariumMessageHandler{
		logger:         logger,
		manageSnapshot: manageSnapshot,
	}
}

func (mh *ReliquariumMessageHandler) Handle(ctx context.Context, msg []byte) (bool, error) {
	var reliquariumCommand message.ServiceControlCommand
	err := json.Unmarshal(msg, &reliquariumCommand)
	if err != nil {
		mh.logger.Error("[ReliquariumMessageHandler] failed unmarshalling message")
		return false, err
	}

	return true, mh.manageSnapshot.Process(ctx, &reliquariumCommand)
}
