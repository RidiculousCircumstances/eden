package consumer

import (
	"context"
	"eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/infrastructure/queue/message"
	brokerLib "eden/shared/broker/interfaces"
	loggerIntf "eden/shared/logger/interfaces"
	"encoding/json"
)

type ReliquariumMessageHandler struct {
	stateManager          interfaces.AppStateManager
	confirmationPublisher interfaces.ServiceCommandConfirmationPublisher
	logger                loggerIntf.Logger
	takeSnapshot          interfaces.ManageSnapshotLifecycle
}

func NewReliquariumMessageHandler(
	logger loggerIntf.Logger,
	takeSnapshot interfaces.ManageSnapshotLifecycle,
) brokerLib.MessageHandler {
	return &ReliquariumMessageHandler{
		logger:       logger,
		takeSnapshot: takeSnapshot,
	}
}

func (mh *ReliquariumMessageHandler) Handle(ctx context.Context, msg []byte) (bool, error) {
	var reliquariumCommand message.ServiceControlCommand
	err := json.Unmarshal(msg, &reliquariumCommand)
	if err != nil {
		mh.logger.Error("[ReliquariumMessageHandler] failed unmarshalling message")
		return false, err
	}

	return true, mh.takeSnapshot.Process(ctx, &reliquariumCommand)
}
