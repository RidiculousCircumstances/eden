package consumer

import (
	"context"
	consumerIntf "eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/infrastructure/queue/message"
	"eden/shared/broker/interfaces"
	loggerIntf "eden/shared/logger/interfaces"
	"encoding/json"
	"go.uber.org/zap"
)

type streamForgeMessageHandler struct {
	messageProcessor consumerIntf.SaveProfiles
	logger           loggerIntf.Logger
}

func NewStreamForgeMessageHandler(messageProcessor consumerIntf.SaveProfiles, logger loggerIntf.Logger) interfaces.MessageHandler {
	return &streamForgeMessageHandler{
		messageProcessor: messageProcessor,
		logger:           logger,
	}
}

func (mh *streamForgeMessageHandler) Handle(ctx context.Context, msg []byte) (bool, error) {
	var parsedMsg message.SaveProfileCommand
	err := json.Unmarshal(msg, &parsedMsg)
	if err != nil {
		mh.logger.Error("[StreamForgeMessageHandler] failed unmarshalling message")
		return false, err
	}

	if processErr := mh.messageProcessor.Process(ctx, parsedMsg); processErr != nil {
		mh.logger.Error("[StreamForgeMessageHandler] failed processing message: ", zap.Error(processErr))
		return false, processErr
	}

	// Сообщение успешно обработано
	return true, nil
}
