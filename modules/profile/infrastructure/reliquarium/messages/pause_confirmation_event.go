package messages

type ServiceName string

const (
	Eden ServiceName = "service_name"
)

type PauseConfirmationEvent struct {
	ServiceName ServiceName
}
