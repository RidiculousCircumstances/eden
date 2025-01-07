package consumer

import (
	"context"
	consumerIntf "eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/infrastructure/queue/message"
	"eden/shared/broker/interfaces"
	"errors"
)

var (
	ErrBase64DecodingFailed = errors.New("failed to decode Base64 payload")
	ErrJSONUnmarshalFailed  = errors.New("failed to unmarshal JSON payload")
)

type streamForgeMessageHandler struct {
	messageProcessor consumerIntf.StreamForgeMessageProcessor
}

func NewStreamForgeMessageHandler(messageProcessor consumerIntf.StreamForgeMessageProcessor) interfaces.MessageHandler {
	return &streamForgeMessageHandler{messageProcessor: messageProcessor}
}

func (mh *streamForgeMessageHandler) Handle(ctx context.Context, msg interface{}) (bool, error) {
	streamForgeMessage, ok := msg.(message.StreamForgeMessage)
	if !ok {
		return false, errors.New("invalid message type, expected: StreamForgeMessage")
	}

	if processErr := mh.messageProcessor.Process(ctx, streamForgeMessage); processErr != nil {
		// Логика обработки ошибки
		return false, processErr
	}

	// Сообщение успешно обработано
	return true, nil
}
