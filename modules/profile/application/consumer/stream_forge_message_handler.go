package consumer

import (
	"context"
	"eden/shared/broker/interfaces"
)

type streamForgeMessageHandler struct {
}

func NewStreamForgeMessageHandler() interfaces.MessageHandler {
	return &streamForgeMessageHandler{}
}

func (mh *streamForgeMessageHandler) Handle(ctx context.Context, msg interface{}) (bool, error) {
	return false, nil
}
