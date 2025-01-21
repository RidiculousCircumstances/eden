package consumer

import (
	"context"
	"eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/infrastructure/queue/message"
	brokerLib "eden/shared/broker/interfaces"
	loggerIntf "eden/shared/logger/interfaces"
	"encoding/json"
	"go.uber.org/zap"
)

type edenSearchMessageHandler struct {
	messageProcessor interfaces.SearchProfiles
	logger           loggerIntf.Logger
}

func NewEdenSearchMessageHandler(messageProcessor interfaces.SearchProfiles, logger loggerIntf.Logger) brokerLib.MessageHandler {
	return &edenSearchMessageHandler{
		messageProcessor: messageProcessor,
		logger:           logger,
	}
}

func (mh *edenSearchMessageHandler) Handle(ctx context.Context, msg []byte) (bool, error) {
	var parsedMsg message.SearchProfilesCommand
	err := json.Unmarshal(msg, &parsedMsg)
	if err != nil {
		mh.logger.Error("[EdenSearchMessageHandler] failed unmarshalling message")
		return false, err
	}

	processErr := mh.messageProcessor.Process(ctx, parsedMsg)
	if processErr != nil {
		mh.logger.Error("[EdenSearchMessageHandler] failed processing message: ", zap.Error(processErr))
		return false, processErr
	}

	return true, nil
}
