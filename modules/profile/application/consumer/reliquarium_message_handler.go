package consumer

import "context"

type reliquariumMessageHandler struct{}

func (mh *reliquariumMessageHandler) Handle(ctx context.Context, msg []byte) (bool, error) {
	return true, nil
}
