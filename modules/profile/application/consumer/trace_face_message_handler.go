package consumer

import (
	"context"
	consumerIntf "eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/infrastructure/queue/message"
	"eden/shared/broker/interfaces"
	"encoding/json"
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

func (mh *TraceFaceMessageHandler) Handle(ctx context.Context, msg []byte) (bool, error) {
	// Десериализуем JSON
	var traceFaceMsg message.TraceFaceMessage
	if unmarshalErr := json.Unmarshal(msg, &traceFaceMsg); unmarshalErr != nil {
		return false, errors.Join(ErrJSONUnmarshalFailed, unmarshalErr)
	}

	// Обрабатываем сообщение
	processErr := mh.messageProcessor.Process(ctx, traceFaceMsg)
	if processErr != nil {
		return false, errors.Join(ErrMessageProcessing, processErr)
	}

	return true, nil
}
