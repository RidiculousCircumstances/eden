package message

type CommandType string

const (
	Pause  CommandType = "pause"
	Resume CommandType = "resume"
)

type ServiceControlCommand struct {
	Command CommandType
}
