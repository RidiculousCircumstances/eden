package message

type CommandType string

const (
	Pause         CommandType = "pause"
	TakeSnapshots CommandType = "takeSnapshots"
	Resume        CommandType = "resume"
)

type ServiceControlCommand struct {
	Command CommandType
}

func NewServiceControlCommand(Command CommandType) *ServiceControlCommand {
	return &ServiceControlCommand{
		Command: Command,
	}
}
