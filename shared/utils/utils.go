package utils

import (
	"context"
)

func CreateContextWithStopChannel(stopCh <-chan interface{}) context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		select {
		case <-stopCh:
			cancel()
		case <-ctx.Done():
		}
	}()

	return ctx
}

func CreateContextWithTimeoutAndStopChannel(stopCh <-chan interface{}) context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		select {
		case <-stopCh:
			cancel()
		case <-ctx.Done():
		}
	}()

	return ctx
}
