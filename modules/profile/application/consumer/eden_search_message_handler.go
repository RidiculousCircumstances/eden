package consumer

import (
	"context"
	"eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/infrastructure/queue/message"
	brokerLib "eden/shared/broker/interfaces"
	"errors"
)

type edenSearchMessageHandler struct {
	messageProcessor interfaces.EdenSearchMessageProcessor
}

func NewEdenSearchMessageHandler(messageProcessor interfaces.EdenSearchMessageProcessor) brokerLib.MessageHandler {
	return &edenSearchMessageHandler{
		messageProcessor: messageProcessor,
	}
}

func (mh *edenSearchMessageHandler) Handle(ctx context.Context, msg interface{}) (bool, error) {
	searchMessage, ok := msg.(message.SearchProfileCommand)
	if !ok {
		return false, errors.New("invalid message type, expected SearchProfileCommand")
	}

	processErr := mh.messageProcessor.Process(ctx, searchMessage)
	if processErr != nil {
		return false, processErr
	}

	return true, nil
}
