package consumer

import (
	"context"
	consumerIntf "eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/infrastructure/queue/message"
	"eden/shared/broker/interfaces"
	"encoding/json"
)

type streamForgeMessageHandler struct {
	messageProcessor consumerIntf.StreamForgeMessageProcessor
}

func NewStreamForgeMessageHandler(messageProcessor consumerIntf.StreamForgeMessageProcessor) interfaces.MessageHandler {
	return &streamForgeMessageHandler{messageProcessor: messageProcessor}
}

func (mh *streamForgeMessageHandler) Handle(ctx context.Context, msg []byte) (bool, error) {
	var parsedMsg message.SaveProfileCommand
	err := json.Unmarshal(msg, &parsedMsg)
	if err != nil {
		return false, err
	}

	if processErr := mh.messageProcessor.Process(ctx, parsedMsg); processErr != nil {
		// Логика обработки ошибки
		return false, processErr
	}

	// Сообщение успешно обработано
	return true, nil
}
