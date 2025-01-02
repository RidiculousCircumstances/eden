package interfaces

import "context"

// Hook представляет собой интерфейс хука для жизненного цикла.
type Hook interface {
	Setup(ctx context.Context) error
	Start(ctx context.Context) error
	Shutdown(ctx context.Context) error
}
