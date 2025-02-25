package messages

type ServiceName string

const (
	Eden ServiceName = "service_name"
)

type CommandType string

const (
	Pause         CommandType = "pause"
	TakeSnapshots CommandType = "takeSnapshots"
	Resume        CommandType = "resume"
)

type CommandConfirmationEvent struct {
	ServiceName ServiceName
	Command     CommandType
}

func NewCommandConfirmationEvent(ServiceName ServiceName, Command CommandType) *CommandConfirmationEvent {
	return &CommandConfirmationEvent{
		ServiceName: ServiceName,
		Command:     Command,
	}
}
