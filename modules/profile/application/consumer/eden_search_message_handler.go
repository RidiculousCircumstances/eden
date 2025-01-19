package consumer

import (
	"context"
	"eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/infrastructure/queue/message"
	brokerLib "eden/shared/broker/interfaces"
	"encoding/json"
)

type edenSearchMessageHandler struct {
	messageProcessor interfaces.EdenSearchMessageProcessor
}

func NewEdenSearchMessageHandler(messageProcessor interfaces.EdenSearchMessageProcessor) brokerLib.MessageHandler {
	return &edenSearchMessageHandler{
		messageProcessor: messageProcessor,
	}
}

func (mh *edenSearchMessageHandler) Handle(ctx context.Context, msg []byte) (bool, error) {
	var parsedMsg message.SearchProfilesCommand
	err := json.Unmarshal(msg, &parsedMsg)
	if err != nil {
		return false, err
	}

	processErr := mh.messageProcessor.Process(ctx, parsedMsg)
	if processErr != nil {
		return false, processErr
	}

	return true, nil
}
