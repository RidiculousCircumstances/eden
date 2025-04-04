package consumer

import (
	"context"
	consumerIntf "eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/infrastructure/queue/message"
	loggerIntf "eden/shared/logger/interfaces"
	"encoding/json"
	"errors"
	"github.com/RidiculousCircumstances/netherway/v2"
	"go.uber.org/zap"
)

var (
	ErrMessageProcessing = errors.New("failed to process message")
)

type TraceFaceMessageHandler struct {
	messageProcessor consumerIntf.SaveFaceInfo
	logger           loggerIntf.Logger
}

func NewTraceFaceMessageHandler(messageProcessor consumerIntf.SaveFaceInfo, logger loggerIntf.Logger) netherway.MessageHandler {
	return &TraceFaceMessageHandler{
		messageProcessor: messageProcessor,
		logger:           logger,
	}
}

func (mh *TraceFaceMessageHandler) Handle(ctx context.Context, msg []byte) (bool, error) {
	var parsedMsg message.SaveFacesCommand
	err := json.Unmarshal(msg, &parsedMsg)
	if err != nil {
		mh.logger.Error("[TraceFaceMessageHandler] failed unmarshalling message")
		return false, err
	}

	processErr := mh.messageProcessor.Process(ctx, parsedMsg)
	if processErr != nil {
		mh.logger.Error("[TraceFaceMessageHandler] failed processing message: ", zap.Error(processErr))
		return false, errors.Join(ErrMessageProcessing, processErr)
	}

	return true, nil
}
