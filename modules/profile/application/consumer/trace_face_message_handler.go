package consumer

import (
	"context"
	consumerIntf "eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/infrastructure/queue/message"
	"eden/shared/broker/interfaces"
	"errors"
)

var (
	ErrMessageProcessing = errors.New("failed to process message")
)

type TraceFaceMessageHandler struct {
	messageProcessor consumerIntf.TraceFaceMessageProcessor
}

func NewTraceFaceMessageHandler(messageProcessor consumerIntf.TraceFaceMessageProcessor) interfaces.MessageHandler {
	return &TraceFaceMessageHandler{
		messageProcessor: messageProcessor,
	}
}

func (mh *TraceFaceMessageHandler) Handle(ctx context.Context, msg interface{}) (bool, error) {
	traceFaceMessage, ok := msg.(message.SaveFacesCommand)
	if !ok {
		return false, errors.New("invalid message type, expected: SaveFacesCommand")
	}

	processErr := mh.messageProcessor.Process(ctx, traceFaceMessage)
	if processErr != nil {
		return false, errors.Join(ErrMessageProcessing, processErr)
	}

	return true, nil
}
