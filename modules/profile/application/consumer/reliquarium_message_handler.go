package consumer

import (
	"context"
	"eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/infrastructure/queue/message"
	"eden/modules/profile/infrastructure/reliquarium/messages"
	brokerLib "eden/shared/broker/interfaces"
	loggerIntf "eden/shared/logger/interfaces"
	"encoding/json"
	"errors"
)

var ErrUnknownCommand = errors.New("unknown command")

type ReliquariumMessageHandler struct {
	stateManager          interfaces.AppStateManager
	confirmationPublisher interfaces.ServiceCommandConfirmationPublisher
	logger                loggerIntf.Logger
}

func NewReliquariumMessageHandler(
	stateManager interfaces.AppStateManager,
	confirmationPublisher interfaces.ServiceCommandConfirmationPublisher,
	logger loggerIntf.Logger,
) brokerLib.MessageHandler {
	return &ReliquariumMessageHandler{
		stateManager:          stateManager,
		confirmationPublisher: confirmationPublisher,
		logger:                logger,
	}
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
		err := mh.confirmationPublisher.Publish(
			ctx,
			messages.NewCommandConfirmationEvent(messages.Eden, messages.Pause, ""),
		)
		if err != nil {
			return true, err
		}
	case message.TakeSnapshots:
		//TODO: реализовать юзкейс снапшота
		return false, nil
	case message.Resume:
		mh.stateManager.Resume()
	default:
		return true, ErrUnknownCommand
	}

	return true, nil
}
