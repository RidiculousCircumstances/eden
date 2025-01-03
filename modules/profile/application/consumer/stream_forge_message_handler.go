package consumer

import (
	"context"
	consumerIntf "eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/infrastructure/queue/message"
	"eden/shared/broker/interfaces"
	"encoding/base64"
	"encoding/json"
	"errors"
)

var (
	ErrInvalidMessageType   = errors.New("invalid message type: expected []byte")
	ErrBase64DecodingFailed = errors.New("failed to decode Base64 payload")
	ErrJSONUnmarshalFailed  = errors.New("failed to unmarshal JSON payload")
)

type streamForgeMessageHandler struct {
	messageProcessor consumerIntf.StreamForgeMessageProcessor
}

func NewStreamForgeMessageHandler(messageProcessor consumerIntf.StreamForgeMessageProcessor) interfaces.MessageHandler {
	return &streamForgeMessageHandler{messageProcessor: messageProcessor}
}

func (mh *streamForgeMessageHandler) Handle(ctx context.Context, msg []byte) (bool, error) {
	// Декодируем Base64
	decodedPayload, decodeErr := base64.StdEncoding.DecodeString(string(msg))
	if decodeErr != nil {
		return false, errors.Join(ErrBase64DecodingFailed, decodeErr)
	}

	// Десериализуем JSON
	var streamForgeMsg message.StreamForgeMessage
	if unmarshalErr := json.Unmarshal(decodedPayload, &streamForgeMsg); unmarshalErr != nil {
		return false, errors.Join(ErrJSONUnmarshalFailed, unmarshalErr)
	}

	// Передаём сообщение в messageProcessor
	if processErr := mh.messageProcessor.Process(ctx, streamForgeMsg); processErr != nil {
		// Логика обработки ошибки
		return false, processErr
	}

	// Сообщение успешно обработано
	return true, nil
}
