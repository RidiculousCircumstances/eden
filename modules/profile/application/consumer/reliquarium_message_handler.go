package consumer

import (
	"context"
	"eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/infrastructure/queue/message"
	loggerIntf "eden/shared/logger/interfaces"
	"encoding/json"
	"errors"
)

var ErrUnknownCommand = errors.New("unknown command")

type ReliquariumMessageHandler struct {
	stateManager interfaces.AppStateManager
	logger       loggerIntf.Logger
}

func (mh *ReliquariumMessageHandler) Handle(ctx context.Context, msg []byte) (bool, error) {
	var reliquariumCommand message.ServiceControlCommand
	err := json.Unmarshal(msg, &reliquariumCommand)
	if err != nil {
		mh.logger.Error("[ReliquariumMessageHandler] failed unmarshalling message")
		return false, err
	}

	switch reliquariumCommand.Command {
	case message.Pause:
		mh.stateManager.Pause()
	case message.Resume:
		mh.stateManager.Resume()
	default:
		return true, ErrUnknownCommand
	}

	return true, nil
}
